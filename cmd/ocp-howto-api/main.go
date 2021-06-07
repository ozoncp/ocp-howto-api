package main

import (
	"fmt"
	"net"

	api "github.com/ozoncp/ocp-howto-api/internal/api"
	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 82
)

func startGrpc(port int) error {
	address := ":" + fmt.Sprint(port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start listening: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	api := api.NewOcpHowtoApi()
	desc.RegisterOcpHowtoApiServer(server, api)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve server: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("Howto API. Author: Ivan Levin")

	if err := startGrpc(grpcPort); err != nil {
		log.Fatal().Msgf("failed to start GRPC server: %v", err)
	}
}
