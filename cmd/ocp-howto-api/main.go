package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	api "github.com/ozoncp/ocp-howto-api/internal/api"
	"github.com/ozoncp/ocp-howto-api/internal/metrics"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcAddress    = ":82"
	metricsAddress = ":9100"
	dbDriver       = "postgres"
	dbDsn          = "user=postgres password=postgres dbname=postgres sslmode=disable"
)

func runMetrics() {
	metrics.Register()
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(metricsAddress, nil)
		if err != nil {
			panic(err)
		}
	}()
}

func runGrpc() error {

	listener, err := net.Listen("tcp", grpcAddress)
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
	runMetrics()

	if err := runGrpc(); err != nil {
		log.Fatal().Msgf("failed to start GRPC server: %v", err)
	}
}
