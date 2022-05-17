package main

import (
	"context"
	"log"

	"myproject/sdk/gogrpc"
	pb "myproject/sdk/gogrpc/proto"
)

func main() {
	config := &gogrpc.Config{
		Host:         "localhost",
		Port:         9001,
		PingInterval: 1,  // seconds
		Timeout:      10, // seconds
		Token:        "xpto",
	}
	// TLS will be ignored if key and cert are empty
	config.TLS.Key = ""
	config.TLS.Cert = ""

	conn, err := gogrpc.Connect(config)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}
	client := gogrpc.NewClient(conn, config)
	req := &pb.SearchRequest{PerPage: 10}
	ctx := client.Auth(context.Background())
	res, err := client.GetTodoClient().Search(ctx, req)
	if err != nil {
		log.Fatalln("[ERROR] Cannot fetch all TODO items:", err)
	}
	items := res.GetTodoItems()
	log.Printf("[INFO] Found %d TODO item(s)\n", len(items))
	log.Println("[INFO] Title:", items[0].Title)
}
