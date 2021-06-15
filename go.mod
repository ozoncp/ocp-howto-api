module github.com/ozoncp/ocp-howto-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.0
	github.com/golang/mock v1.5.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.2
	github.com/onsi/ginkgo v1.16.3
	github.com/onsi/gomega v1.13.0
	github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.22.0
	golang.org/x/net v0.0.0-20210610132358-84b48f89b13b // indirect
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	google.golang.org/genproto v0.0.0-20210614143202-012ab6975634 // indirect
	google.golang.org/grpc v1.38.0
)

replace github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api => ./pkg/ocp-howto-api
