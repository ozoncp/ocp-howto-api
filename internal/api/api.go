package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/metrics"
	"github.com/ozoncp/ocp-howto-api/internal/producer"
	"github.com/ozoncp/ocp-howto-api/internal/repo"
	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func invalidArgErr(err error) error {
	return status.Error(codes.InvalidArgument, err.Error())
}

func toStatus(err error) error {
	if err == nil {
		return status.Error(codes.OK, "")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return status.Error(codes.NotFound, err.Error())
	}
	return status.Error(codes.Internal, err.Error())
}

func (a *api) CreateHowtoV1(
	ctx context.Context,
	req *desc.CreateHowtoV1Request,
) (*desc.CreateHowtoV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Requested to create howto with invalid arguments")
		return nil, invalidArgErr(err)
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "Create howto")
	span.SetTag("Question", req.Params.Question)
	defer span.Finish()

	metrics.IncrementCreateRequests(1)
	log.Info().
		Uint64("CourseId", req.Params.CourseId).
		Str("Q", req.Params.Question).
		Str("A", req.Params.Answer).
		Msg("Requested to create howto")

	id, err := a.repo.AddHowto(ctx, fromParams(req.Params))

	if err != nil {
		log.Error().Err(err).Msg("Failed to create howto")
		return nil, toStatus(err)
	}

	log.Info().Uint64("Id", id).Msg("Howto created successfully")
	a.recordCreates([]uint64{id})

	return &desc.CreateHowtoV1Response{Id: id}, toStatus(nil)
}

func (a *api) MultiCreateHowtoV1(
	ctx context.Context,
	req *desc.MultiCreateHowtoV1Request,
) (*desc.MultiCreateHowtoV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Requested to create howtos with invalid arguments")
		return nil, invalidArgErr(err)
	}

	opName := fmt.Sprintf("Create %v howtos", len(req.Params))
	span, ctx := opentracing.StartSpanFromContext(ctx, opName)
	span.SetTag("Count", len(req.Params))
	defer span.Finish()

	metrics.IncrementCreateRequests(len(req.Params))
	log.Info().Msgf("Requested to create %v howtos", len(req.Params))

	var toAdd []howto.Howto
	for _, p := range req.Params {
		toAdd = append(toAdd, fromParams(p))
	}

	added, err := a.repo.AddHowtos(ctx, toAdd)
	if err != nil {
		log.Error().Int("successfully added", len(added)).Err(err).
			Msg("Error occurred when creating howtos.")
		return nil, toStatus(err)
	}

	addedCount := len(added)
	expectedCount := len(req.Params)
	log.Info().Msgf("%v howtos created successfully", addedCount)
	if addedCount != expectedCount {
		log.Warn().
			Int("expected", expectedCount).
			Int("added", addedCount).
			Msg("Number of added howtos does not match requested number")
	}

	a.recordCreates(added)
	return &desc.MultiCreateHowtoV1Response{Ids: added}, toStatus(nil)
}

func (a *api) UpdateHowtoV1(
	ctx context.Context,
	req *desc.UpdateHowtoV1Request,
) (*desc.UpdateHowtoV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Requested to update howto with invalid arguments")
		return nil, invalidArgErr(err)
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "Update howto")
	span.SetTag("Id", req.Howto.Id)
	defer span.Finish()

	metrics.IncrementUpdateRequests(1)
	log.Info().Uint64("Id", req.Howto.Id).Msg("Requested to update howto")

	if err := a.repo.UpdateHowto(ctx, fromMessage(req.Howto)); err != nil {
		log.Error().Err(err).Msg("Failed to update howto")
		return nil, toStatus(err)
	}

	log.Info().Msg("Howto updated successfully")
	a.recordUpdates([]uint64{req.Howto.Id})

	return &desc.UpdateHowtoV1Response{}, toStatus(nil)
}

func (a *api) DescribeHowtoV1(
	ctx context.Context,
	req *desc.DescribeHowtoV1Request,
) (*desc.DescribeHowtoV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Requested to describe howto with invalid arguments")
		return nil, invalidArgErr(err)
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "Describe howto")
	span.SetTag("Id", req.Id)
	defer span.Finish()

	log.Info().Uint64("Id", req.Id).Msg("Requested to describe howto")

	howto, err := a.repo.DescribeHowto(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe howto")
		return nil, toStatus(err)
	}

	log.Info().Msg("Howto described successfully")
	return &desc.DescribeHowtoV1Response{Howto: toMessage(howto)}, toStatus(nil)
}

func (a *api) ListHowtosV1(
	ctx context.Context,
	req *desc.ListHowtosV1Request,
) (*desc.ListHowtosV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Requested to list howtos with invalid arguments")
		return nil, invalidArgErr(err)
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "List howtos")
	span.SetTag("Count", req.Count)
	defer span.Finish()

	log.Info().Msgf("Requested to list %v howtos starting from %v", req.Count, req.Offset)

	howtos, err := a.repo.ListHowtos(ctx, req.Offset, req.Count)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list howtos")
		return nil, toStatus(err)
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

	return &desc.ListHowtosV1Response{Howtos: result}, toStatus(nil)
}

func (a *api) RemoveHowtoV1(
	ctx context.Context,
	req *desc.RemoveHowtoV1Request,
) (*desc.RemoveHowtoV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Requested to remove howto with invalid arguments")
		return nil, invalidArgErr(err)
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "Remove howto")
	span.SetTag("Id", req.Id)
	defer span.Finish()

	metrics.IncrementRemoveRequests(1)
	log.Info().Uint64("Id", req.Id).Msg("Requested to remove howto")

	if err := a.repo.RemoveHowto(ctx, req.Id); err != nil {
		log.Error().Err(err).Msg("Failed to remove howto")
		return nil, toStatus(err)
	}

	log.Info().Msg("Howto removed successfully")
	a.recordRemoves([]uint64{req.Id})

	return &desc.RemoveHowtoV1Response{}, toStatus(nil)
}

func (a *api) recordCreates(ids []uint64) {
	metrics.IncrementCreates(len(ids))
	a.prod.SendEvent(newCudEvent(producer.EventTypeCreated, ids))
}

func (a *api) recordUpdates(ids []uint64) {
	metrics.IncrementUpdates(len(ids))
	a.prod.SendEvent(newCudEvent(producer.EventTypeUpdated, ids))
}

func (a *api) recordRemoves(ids []uint64) {
	metrics.IncrementRemoves(len(ids))
	a.prod.SendEvent(newCudEvent(producer.EventTypeRemoved, ids))
}

func newCudEvent(type_ producer.EventType, ids []uint64) producer.Event {
	return producer.Event{
		Type:      type_,
		Timestamp: time.Now(),
		Body: map[string]interface{}{
			"Ids": ids,
		},
	}
}
