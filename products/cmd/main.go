package main

import (
	"log"
	"net"
	"os"

	pb "github.com/arthurh0812/coffee-shop/protos/products"

	"github.com/arthurh0812/products/pkg/server"
)

func main() {
	ls, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)

	srv := server.NewServer()

	go func() {
		err := srv.Serve(ls)
		if err != nil {
			log.Fatal(err)
		}
	}()

	server.WaitAndGracefulShutdown(srv, logger)
}
