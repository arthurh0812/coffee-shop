package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	return grpc.NewServer(grpc.ConnectionTimeout(30 *time.Second), grpc.MaxHeaderListSize(2048))
}

func WaitAndGracefulShutdown(srv *grpc.Server, l *log.Logger) {
	sigchan := make(chan os.Signal, 1)

	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, syscall.SIGTERM)

	sg := <-sigchan
	l.SetPrefix("[Server]")
	l.Printf("Received terminate signal: %s\n", sg)
	l.Println("Gracefully shutting server down...")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	srv.Stop()
	cancel()
}
