package api

import (
	"context"

	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	log "github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpHowtoApiServer
	repo repo.Repo
}

func NewOcpHowtoApi(repo repo.Repo) desc.OcpHowtoApiServer {
	return &api{
		repo: repo,
	}
}

func toMessage(howto howto.Howto) *desc.Howto {
	return &desc.Howto{
		Id:       howto.Id,
		CourseId: howto.CourseId,
		Question: howto.Question,
		Answer:   howto.Answer,
	}
}

func (a *api) CreateHowtoV1(
	ctx context.Context,
	req *desc.CreateHowtoV1Request,
) (*desc.CreateHowtoV1Response, error) {

	log.Info().
		Str("Q", req.Question).
		Str("A", req.Answer).
		Msg("Requested to create howto")

	id, err := a.repo.AddHowto(
		ctx,
		howto.Howto{
			CourseId: req.CourseId,
			Question: req.Question,
			Answer:   req.Answer,
		})

	if err != nil {
		log.Error().Err(err).Msg("Failed to create howto")
		return &desc.CreateHowtoV1Response{}, err
	}

	log.Info().Uint64("Id", id).Msg("Howto created successfully")
	return &desc.CreateHowtoV1Response{Id: id}, nil
}

func (a *api) DescribeHowtoV1(
	ctx context.Context,
	req *desc.DescribeHowtoV1Request,
) (*desc.DescribeHowtoV1Response, error) {

	log.Info().Uint64("Id", req.Id).Msg("Requested to describe howto")

	howto, err := a.repo.DescribeHowto(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe howto")
		return &desc.DescribeHowtoV1Response{}, err
	}

	log.Info().Msg("Howto described successfully")
	return &desc.DescribeHowtoV1Response{Howto: toMessage(howto)}, nil
}

func (a *api) ListHowtosV1(
	ctx context.Context,
	req *desc.ListHowtosV1Request,
) (*desc.ListHowtosV1Response, error) {

	log.Info().Msgf("Requested to list %v howtos", req.Count)

	howtos, err := a.repo.ListHowtos(ctx, req.Offset, req.Count)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list howtos")
		return &desc.ListHowtosV1Response{}, err
	}

	log.Info().Msg("Howtos listed successfully")

	if len(howtos) != int(req.Count) {
		log.Warn().
			Uint64("expected", req.Count).
			Int("resulted", len(howtos)).
			Msg("Result count does not match requested count")
	}

	var result []*desc.Howto
	for _, h := range howtos {
		result = append(result, toMessage(h))
	}

	return &desc.ListHowtosV1Response{Howtos: result}, nil
}

func (a *api) RemoveHowtoV1(
	ctx context.Context,
	req *desc.RemoveHowtoV1Request,
) (*desc.RemoveHowtoV1Response, error) {

	log.Info().Uint64("Id", req.Id).Msg("Requested to remove howto")

	if err := a.repo.RemoveHowto(ctx, req.Id); err != nil {
		log.Error().Err(err).Msg("Failed to remove howto")
		return &desc.RemoveHowtoV1Response{}, err
	}

	log.Info().Msg("Howto removed successfully")
	return &desc.RemoveHowtoV1Response{}, nil
}
