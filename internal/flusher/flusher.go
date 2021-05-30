package flusher

import (
	"github.com/ozoncp/ocp-howto-api/internal/howto"
)

type Flusher interface {
	Flush(howtos []howto.Howto) []howto.Howto
}
