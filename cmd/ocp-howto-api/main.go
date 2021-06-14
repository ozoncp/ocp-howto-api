package main

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	api "github.com/ozoncp/ocp-howto-api/internal/api"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 82
	dbDriver = "postgres"
	dbDsn    = "user=tiger password=scott dbname=postgres sslmode=disable"
)

func runGrpc(port int) error {
	address := ":" + fmt.Sprint(port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start listening: %v", err)
	}

	db, err := sqlx.Connect(dbDriver, dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	server := grpc.NewServer()
	reflection.Register(server)
	api := api.NewOcpHowtoApi(repo.NewRepo(*db))
	desc.RegisterOcpHowtoApiServer(server, api)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve server: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("Howto API. Author: Ivan Levin")

	if err := runGrpc(grpcPort); err != nil {
		log.Fatal().Msgf("failed to start GRPC server: %v", err)
	}
}
