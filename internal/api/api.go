package api

import (
	"context"

	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	log "github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpHowtoApiServer
}

func NewOcpHowtoApi() desc.OcpHowtoApiServer {
	return &api{}
}

func (a *api) CreateHowtoV1(
	ctx context.Context,
	req *desc.CreateHowtoV1Request,
) (*desc.CreateHowtoV1Response, error) {
	log.Info().
		Str("Q", req.Question).
		Str("A", req.Answer).
		Msg("Requested to create howto")

	return &desc.CreateHowtoV1Response{}, nil
}

func (a *api) DescribeHowtoV1(
	ctx context.Context,
	req *desc.DescribeHowtoV1Request,
) (*desc.DescribeHowtoV1Response, error) {
	log.Info().Uint64("Id", req.Id).Msg("Requested to describe howto")
	return &desc.DescribeHowtoV1Response{}, nil
}

func (a *api) ListHowtosV1(
	ctx context.Context,
	req *desc.ListHowtosV1Request,
) (*desc.ListHowtosV1Response, error) {
	log.Info().Msgf("Requested to list %v howtos", req.Length)
	return &desc.ListHowtosV1Response{}, nil
}

func (a *api) RemoveHowtoV1(
	ctx context.Context,
	req *desc.RemoveHowtoV1Request,
) (*desc.RemoveHowtoV1Response, error) {
	log.Info().Uint64("Id", req.Id).Msg("Requested to remove howto")
	return &desc.RemoveHowtoV1Response{}, nil
}
