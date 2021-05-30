package flusher

import (
	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
)

type Flusher interface {
	Flush(howtos []howto.Howto) []howto.Howto
}

type flusher struct {
	repo repo.Repo
}

func (f *flusher) Flush(howtos []howto.Howto) []howto.Howto {
	for i, howto := range howtos {
		if _, err := f.repo.AddHowto(howto); err != nil {
			return howtos[i:]
		}
	}
	return howtos
}

func New(repo repo.Repo) Flusher {
	return &flusher{
		repo: repo,
	}
}
