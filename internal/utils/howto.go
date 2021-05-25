package utils

import (
	"errors"

	"github.com/ozoncp/ocp-howto-api/internal/howto"
)

func SplitToBulks(questions []howto.Howto, batchSize int) [][]howto.Howto {
	var batches [][]howto.Howto
	if len(questions) == 0 || batchSize <= 0 {
		return batches
	}

	index := 0
	for {
		remaining := len(questions) - index
		if remaining <= batchSize {
			batches = append(batches, questions[index:index+remaining])
			break
		}
		batches = append(batches, questions[index:index+batchSize])
		index += batchSize
	}
	return batches
}

func ConvertToMap(howtos []howto.Howto) (map[uint64]howto.Howto, error) {
	mapped := make(map[uint64]howto.Howto, len(howtos))
	for _, value := range howtos {
		if _, found := mapped[value.Id]; found {
			return nil, errors.New("duplicate value detected. Make sure collection has no duplicates")
		}
		mapped[value.Id] = value
	}
	return mapped, nil
}
