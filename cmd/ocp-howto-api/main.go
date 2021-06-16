package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	api "github.com/ozoncp/ocp-howto-api/internal/api"
	"github.com/ozoncp/ocp-howto-api/internal/config"
	"github.com/ozoncp/ocp-howto-api/internal/metrics"
	"github.com/ozoncp/ocp-howto-api/internal/producer"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var cfg *config.Config

func runMetrics() {
	metrics.Register()
	http.Handle(cfg.Metrics.Pattern, promhttp.Handler())
	go func() {
		err := http.ListenAndServe(cfg.Metrics.Address, nil)
		if err != nil {
			panic(err)
		}
	}()
}

func runGrpc() error {

	listener, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		return fmt.Errorf("failed to start listening: %v", err)
	}

	dbCfg := &cfg.Database
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=%v",
		dbCfg.User, dbCfg.Password, dbCfg.Database, dbCfg.Port, dbCfg.SslMode)

	db, err := sqlx.Connect(dbCfg.Driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	prod, err := producer.New(cfg.Kafka.Brokers[0], cfg.Kafka.Topic, 100)
	if err != nil {
		return err
	}
	prod.Init()
	defer prod.Close()

	server := grpc.NewServer()
	reflection.Register(server)
	api := api.NewOcpHowtoApi(repo.NewRepo(*db, 10), prod)
	desc.RegisterOcpHowtoApiServer(server, api)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve server: %v", err)
	}

	return nil
}

func main() {
	var err error
	cfg, err = config.Read("config.yml")
	if err != nil {
		log.Fatal().Msgf("failed to open configuration file: %v", err)
		return
	}

	fmt.Printf("%v. Author: %v", cfg.Project.Name, cfg.Project.Author)

	runMetrics()
	if err := runGrpc(); err != nil {
		log.Fatal().Msgf("failed to start GRPC server: %v", err)
	}
}
