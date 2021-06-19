package api

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/metrics"
	"github.com/ozoncp/ocp-howto-api/internal/producer"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	log "github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpHowtoApiServer
	repo repo.Repo
	prod producer.Producer
}

func NewOcpHowtoApi(
	repo repo.Repo,
	prod producer.Producer,
) desc.OcpHowtoApiServer {
	return &api{
		repo: repo,
		prod: prod,
	}
}

func toMessage(h howto.Howto) *desc.Howto {
	return &desc.Howto{
		Id: h.Id,
		Params: &desc.HowtoParams{
			CourseId: h.CourseId,
			Question: h.Question,
			Answer:   h.Answer,
		}}
}

func fromMessage(h *desc.Howto) howto.Howto {
	return howto.Howto{
		Id:       h.Id,
		CourseId: h.Params.CourseId,
		Question: h.Params.Question,
		Answer:   h.Params.Answer,
	}
}

func fromParams(params *desc.HowtoParams) howto.Howto {
	return howto.Howto{
		CourseId: params.CourseId,
		Question: params.Question,
		Answer:   params.Answer,
	}
}

func (a *api) CreateHowtoV1(
	ctx context.Context,
	req *desc.CreateHowtoV1Request,
) (*desc.CreateHowtoV1Response, error) {

	log.Info().
		Uint64("CourseId", req.Params.CourseId).
		Str("Q", req.Params.Question).
		Str("A", req.Params.Answer).
		Msg("Requested to create howto")

	id, err := a.repo.AddHowto(ctx, fromParams(req.Params))

	if err != nil {
		log.Error().Err(err).Msg("Failed to create howto")
		return &desc.CreateHowtoV1Response{}, err
	}

	log.Info().Uint64("Id", id).Msg("Howto created successfully")
	metrics.IncrementCreates(1)

	event := producer.Event{
		Type: producer.EventTypeCreated,
		Body: map[string]interface{}{
			"Id":        id,
			"Timestamp": time.Now(),
		},
	}
	a.prod.SendEvent(event)

	return &desc.CreateHowtoV1Response{Id: id}, nil
}

func (a *api) MultiCreateHowtoV1(
	ctx context.Context,
	req *desc.MultiCreateHowtoV1Request,
) (*desc.MultiCreateHowtoV1Response, error) {

	opName := fmt.Sprintf("Create %v howtos", len(req.Params))
	span, ctx := opentracing.StartSpanFromContext(ctx, opName)
	defer span.Finish()

	log.Info().Msgf("Requested to create %v howtos", len(req.Params))

	var toAdd []howto.Howto
	for _, p := range req.Params {
		toAdd = append(toAdd, fromParams(p))
	}

	added, err := a.repo.AddHowtos(ctx, toAdd)
	addedCount := len(added)
	metrics.IncrementCreates(uint64(addedCount))

	event := producer.Event{
		Type: producer.EventTypeCreated,
		Body: map[string]interface{}{
			"Id":        added,
			"Timestamp": time.Now(),
		},
	}
	a.prod.SendEvent(event)

	if err != nil {
		log.Error().
			Err(err).
			Msgf("Error occurred when creating howtos. %v actually was added", addedCount)
		return &desc.MultiCreateHowtoV1Response{Ids: added}, err
	}

	log.Info().Msgf("%v howtos created successfully", addedCount)

	expectedCount := len(req.Params)
	if addedCount != expectedCount {
		log.Warn().
			Int("expected", expectedCount).
			Int("added", addedCount).
			Msg("Number of added howtos does not match requested number")
	}

	return &desc.MultiCreateHowtoV1Response{Ids: added}, nil
}

func (a *api) UpdateHowtoV1(
	ctx context.Context,
	req *desc.UpdateHowtoV1Request,
) (*desc.UpdateHowtoV1Response, error) {

	log.Info().Uint64("Id", req.Howto.Id).Msg("Requested to update howto")

	if err := a.repo.UpdateHowto(ctx, fromMessage(req.Howto)); err != nil {
		log.Error().Err(err).Msg("Failed to update howto")
		return &desc.UpdateHowtoV1Response{}, err
	}

	log.Info().Msg("Howto updated successfully")
	metrics.IncrementUpdates(1)

	event := producer.Event{
		Type: producer.EventTypeUpdated,
		Body: map[string]interface{}{
			"Id":        req.Howto.Id,
			"Timestamp": time.Now(),
		},
	}
	a.prod.SendEvent(event)

	return &desc.UpdateHowtoV1Response{}, nil
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

	log.Info().Msgf("Requested to list %v howtos starting from %v", req.Count, req.Offset)

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
	metrics.IncrementRemoves(1)

	event := producer.Event{
		Type: producer.EventTypeRemoved,
		Body: map[string]interface{}{
			"Id":        req.Id,
			"Timestamp": time.Now(),
		},
	}
	a.prod.SendEvent(event)

	return &desc.RemoveHowtoV1Response{}, nil
}
