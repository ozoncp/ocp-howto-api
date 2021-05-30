package repo

import "github.com/ozoncp/ocp-howto-api/internal/howto"

type Repo interface {
	AddHowto(howto.Howto) (uint64, error)
}
