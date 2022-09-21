package websocketserver

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/logger"
)

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

var hub = Hub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[*connection]bool),
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		hub.unregister <- s
		c.ws.Close()
	}()

	maxMessageSize := conf.Get().Websocket.MaxMessageSize
	c.ws.SetReadLimit(maxMessageSize)

	pongWait := time.Duration(conf.Get().Websocket.PongWait) * time.Second
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logger.Error(err)
			}
			break
		}
		m := message{msg, s.room}
		hub.broadcast <- m
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	writeWait := time.Duration(conf.Get().Websocket.WriteWait) * time.Second
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	pongWait := time.Duration(conf.Get().Websocket.PongWait) * time.Second
	pingPeriod := time.Duration(conf.Get().Websocket.PingPeriod) * time.Second
	if pingPeriod >= pongWait {
		pingPeriod = (pongWait * 9) / 10
	}
	ticker := time.NewTicker(time.Duration(pingPeriod))
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func (h *WebsocketServer) serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  h.config.Websocket.ReadBufferSize,
		WriteBufferSize: h.config.Websocket.WriteBufferSize,
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, roomId}
	hub.register <- s
	go s.writePump()
	go s.readPump()
}
