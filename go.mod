module github.com/ozoncp/ocp-howto-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/golang/mock v1.5.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.10.2
	github.com/onsi/ginkgo v1.16.3
	github.com/onsi/gomega v1.13.0
	github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0
	google.golang.org/grpc v1.38.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api => ./pkg/ocp-howto-api
