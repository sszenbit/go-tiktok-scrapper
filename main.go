package main

import (
	"net"

	srpc "github.com/sszenbit/go-tiktok-scrapper/scrapperpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	srpc.RegisterTikTokScrapperServiceServer(srv, &srpc.Server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
