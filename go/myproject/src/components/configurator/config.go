package configurator

import (
	"log"
	"os"
	"sync"

	"github.com/gustavohenrique/coolconf"
)

const (
	CONTEXT_CLIENT_ID = "clientID"
)

type Config struct {
	Debug bool `env:"DEBUG" `
	Auth  struct {
		Disabled    bool `env:"AUTH_DISABLED" yaml:"disabled"`
		SkipRouters []string
		Jwt         struct {
			Header     string `env:"AUTH_HEADER" yaml:"header" default:"X-Auth-Token"`
			Secret     string `env:"AUTH_TOKEN_SECRET" yaml:"secret"`
			Expiration string `env:"AUTH_TOKEN_EXPIRATION" yaml:"expiration" default:"86400"`
			Audience   string `env:"AUTH_AUDIENCE" yaml:"audience" default:"web"`
			Test       struct {
				ClientID         string `yaml:"client_id"`
				ClientSecret     string `yaml:"client_secret"`
				ClientSecretHash string `yaml:"client_secret_hash"`
			}
		}
	}
	Log struct {
		Level  string `env:"LOG_LEVEL" yaml:"level" default:"info"`
		Output string `env:"LOG_OUTPUT" yaml:"output" default:"stdout"`
		Format string `env:"LOG_FORMAT" yaml:"format" default:"text"`
	}
{{ if or .HasGrpcServer .HasGrpcWebServer }}
	Grpc struct {
		Address           string   `env:"GRPC_ADDRESS" yaml:"address" default:"0.0.0.0"`
		Port              int      `env:"GRPC_PORT" yaml:"port" default:"8002"`
		SkipRouters       []string `default:"/SignIn,/SignUp"`
		MaxReceiveMsgSize int      `env:"GRPC_MAX_RECEIVE_MSG_SIZE" yaml:"max_receive_msg_size" default:"5368709120"` // 5MB
		MaxSendMsgSize    int      `env:"GRPC_MAX_SEND_MSG_SIZE" yaml:"max_send_msg_size" default:"5368709120"`       // 5MB
		TLS               struct {
			Enabled bool   `env:"GRPC_TLS_ENABLED" yaml:"enabled"`
			Key     string `env:"GRPC_TLS_KEY" yaml:"key"`
			Cert    string `env:"GRPC_TLS_CERT" yaml:"cert"`
		}
	}
{{ end }}
{{ if .HasHttpServer }}
	Http struct {
		Address string   `env:"HTTP_ADDRESS" yaml:"address" default:"0.0.0.0"`
		Port    int      `env:"HTTP_PORT" yaml:"port" default:"8001"`
		Origins []string `env:"HTTP_ALLOW_ORIGINS" yaml:"origins" default:"*"`
		TLS     struct {
			Enabled bool   `env:"HTTP_TLS_ENABLED" yaml:"enabled"`
			Key     string `env:"HTTP_TLS_KEY" yaml:"key"`
			Cert    string `env:"HTTP_TLS_CERT" yaml:"cert"`
		}
	}
{{ end }}
{{ if .HasNatsServer }}
	Nats struct {
		Address    string `env:"NATS_ADDRESS" yaml:"address" default:"0.0.0.0"`
		Port       int    `env:"NATS_PORT" yaml:"port" default:"8001"`
		ServerName string `env:"NATS_SERVER_NAME" yaml:"server_name" default:"backend"`
		ClientName string `env:"NATS_CLIENT_NAME" yaml:"client_name" default:"agent"`
		NoSigs     bool   `env:"NATS_NO_SIGS" yaml:"nosigs"`
		Debug      bool   `env:"NATS_DEBUG" yaml:"debug"`
		Trace      bool   `env:"NATS_TRACE" yaml:"trace"`
		TLS        struct {
			Enabled bool   `env:"NATS_TLS_ENABLED" yaml:"enabled"`
			Key     string `env:"NATS_TLS_KEY" yaml:"key"`
			Cert    string `env:"NATS_TLS_CERT" yaml:"cert"`
		}
	}
{{ end }}
	Websocket struct {
		RouterPrefix    string `env:"WS_ROUTER_PREFIX" yaml:"router_prefix" default:"/ws"`
		ReadBufferSize  int    `env:"WS_READ_BUFFER_SIZE" yaml:"read_buffer_size"`
		WriteBufferSize int    `env:"WS_WRITE_BUFFER_SIZE" yaml:"write_buffer_size"`
		// Time allowed to write a message to the peer.
		WriteWait int `env:"WS_WRITE_WAIT" yaml:"write_wait"`
		// Time allowed to read the next pong message from the peer.
		PongWait int `env:"WS_PONG_WAIT" yaml:"pong_wait"`
		// Send pings to peer with this period. Must be less than pongWait.
		PingPeriod int `env:"WS_PING_PERIOD" yaml:"ping_period"`
		// Maximum message size allowed from peer.
		MaxMessageSize int64 `env:"WS_MAX_MESSAGE_SIZE" yaml:"max_message_size"`
	}
	Store struct {
		{{ if .HasPostgres }}
		Postgres struct {
			URL             string `env:"STORE_POSTGRES_URL" yaml:"url" default:"postgres://admin:123456@127.0.0.1/maindb?sslmode=disable"`
			MaxOpenConns    int    `env:"STORE_POSTGRES_MAX_OPEN_CONN" yaml:"max_open_conns"`
			MaxIdleConns    int    `env:"STORE_POSTGRES_MAX_IDLE_CONN" yaml:"max_idle_conns"`
			MaxConnLifetime int    `env:"STORE_POSTGRES_MAX_CONN_LIFETIME" yaml:"max_conn_lifetime" default:"480"`
			Schema          string `env:"STORE_POSTGRES_SCHEMA" yaml:"schema"`
		}
		{{ end }}
		{{ if .HasSqlite }}
		Sqlite struct {
			Address string `env:"STORE_SQLITE_ADDRESS" yaml:"address" default:":memory:"`
			Schema  string `env:"STORE_SQLITE_SCHEMA" yaml:"schema"`
		}
		{{ end }}
		{{ if .HasDgraph }}
		Dgraph struct {
			Address string `env:"STORE_DGRAPH_ADDRESS" yaml:"address" default:"localhost:9080"`
		}
		{{ end }}
	}
}

var globalConfig *Config

func Parse(configFile string) *Config {
	coolconf.New(coolconf.Settings{
		Source:                  configFile,
		ShouldCreateDefaultYaml: true,
	})
	if err := coolconf.Load(&globalConfig); err != nil {
		log.Fatalln("Parser:", err)
	}
	globalConfig.Auth.SkipRouters = []string{"/SignIn", "/FindOneByEmail", "/CreatePasswordFor"}
	singleton(globalConfig)
	return globalConfig
}

func Get() *Config {
	if globalConfig != nil {
		return globalConfig
	}
	key := "CONFIG_FILE"
	var configFile string = os.Getenv(key)
	if configFile == "" {
		log.Fatalln("[INFO]", key, "is empty. Cannot read the configuration file")
	}
	return Parse(configFile)
}

func singleton(c *Config) {
	var once sync.Once
	once.Do(func() {
		globalConfig = c
	})
}
