package main

import (
	"context"
	"fmt"
	pinger "grpcserver/proto"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

type Pinger struct {
	pinger.UnimplementedPingerServer
	port string
}

func (p *Pinger) Ping(ctx context.Context, in *pinger.PingRequest) (*pinger.PingResponse, error) {
	return &pinger.PingResponse{Port: p.port}, nil
}

func servegrpc(ctx context.Context, port string) {
	// port listener
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer()
	pinger.RegisterPingerServer(grpcServer, &Pinger{})
	go func() {
		grpcServer.Serve(lis)
	}()

	<-ctx.Done()
	grpcServer.GracefulStop()

}

func serveHTTP(ctx context.Context, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		return
	}
	http.Serve(lis, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
}

func main() {

}
