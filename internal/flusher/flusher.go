package flusher

import (
	"context"

	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
)

// Flusher - интерфейс для добавления в Repo нескольких сущностей Howto
type Flusher interface {
	// Flush добавляет в Repo несколько сущностей Howto
	// В случае ошибки возвращает слайс из Howto, которые не удалось добавить
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

// New создает экземпляр Flusher
func New(repo repo.Repo) Flusher {
	return &flusher{
		repo: repo,
	}
}
