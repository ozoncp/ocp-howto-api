package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-howto-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-howto-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozoncp/ocp-howto-api/internal/producer Producer
