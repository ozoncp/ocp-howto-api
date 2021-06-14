package flusher

import (
	"context"

	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
)

type Flusher interface {
	Flush(context.Context, []howto.Howto) []howto.Howto
}

type flusher struct {
	repo repo.Repo
}

func (f *flusher) Flush(ctx context.Context, howtos []howto.Howto) []howto.Howto {
	for i, howto := range howtos {
		if _, err := f.repo.AddHowto(ctx, howto); err != nil {
			return howtos[i:]
		}
	}
	return nil
}

func New(repo repo.Repo) Flusher {
	return &flusher{
		repo: repo,
	}
}
