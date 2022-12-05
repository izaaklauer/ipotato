package server

import (
	"context"
	"log"

        "github.com/izaaklauer/ipotato/config"
	ipotatov1 "github.com/izaaklauer/ipotato/gen/proto/go/ipotato/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IpotatoServer struct {
	ipotatov1.UnimplementedIpotatoServiceServer

	config config.Ipotato
}

// NewIpotatoServer initializes a new server from config
func NewIpotatoServer(config config.Ipotato) (*IpotatoServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := IpotatoServer{
		config: config,
	}

	return &server, nil
}

func (s * IpotatoServer) HelloWorld(
	ctx context.Context,
	req *ipotatov1.HelloWorldRequest,
) (*ipotatov1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &ipotatov1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
