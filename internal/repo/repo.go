package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/utils"
	"github.com/rs/zerolog/log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var dummyHowto = howto.Howto{}

type Repo interface {
	AddHowto(context.Context, howto.Howto) (uint64, error)
	AddHowtos(context.Context, []howto.Howto) ([]uint64, error)
	UpdateHowto(context.Context, howto.Howto) error
	RemoveHowto(ctx context.Context, id uint64) error
	DescribeHowto(ctx context.Context, id uint64) (howto.Howto, error)
	ListHowtos(ctx context.Context, offset uint64, count uint64) ([]howto.Howto, error)
}

func NewRepo(db sqlx.DB, batchSize int) Repo {
	return &repo{
		db: db,
		table: howtoTable{
			name: "howtos",
			columns: howtoColumns{
				id:       "id",
				courseId: "course_id",
				question: "question",
				answer:   "answer",
			},
		},
		placeholder: sqr.Dollar,
		batchSize:   batchSize,
	}
}

func howtoRows(h *howto.Howto) []interface{} {
	return []interface{}{&h.Id, &h.CourseId, &h.Question, &h.Answer}
}

type repo struct {
	Repo
	db          sqlx.DB
	table       howtoTable
	placeholder sqr.PlaceholderFormat
	batchSize   int
}

func (repo *repo) AddHowto(ctx context.Context, h howto.Howto) (uint64, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Add howto")
	defer span.Finish()

	added, err := repo.insertBatch(ctx, []howto.Howto{h})
	if err != nil {
		return 0, err
	}

	if len(added) == 0 {
		return 0, errors.New("unexpected fail to insert howto")
	}

	return added[0], nil
}

func (repo *repo) AddHowtos(ctx context.Context, howtos []howto.Howto) ([]uint64, error) {

	opName := fmt.Sprintf("Add %v howtos", len(howtos))
	span, ctx := opentracing.StartSpanFromContext(ctx, opName)
	defer span.Finish()

	if len(howtos) == 0 {
		return nil, nil
	}

	added := make([]uint64, 0, len(howtos))
	batches := utils.SplitToBulks(howtos, repo.batchSize)
	for _, batch := range batches {
		ids, err := repo.insertBatch(ctx, batch)
		if err != nil {
			log.Error().Err(err).Msg("failed to insert batch of howtos")
			return added, err
		}
		added = append(added, ids...)
	}

	return added, nil
}

func (repo *repo) insertBatch(ctx context.Context, howtos []howto.Howto) ([]uint64, error) {

	opName := fmt.Sprintf("Insert batch with %v howtos", len(howtos))
	span, ctx := opentracing.StartSpanFromContext(ctx, opName)
	defer span.Finish()

	cols := repo.table.columns
	query := sqr.Insert(repo.table.name).
		Columns(cols.courseId, cols.question, cols.answer).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder).Columns().
		Suffix(fmt.Sprintf("RETURNING %v", cols.id))

	for _, howto := range howtos {
		query = query.Values(howto.CourseId, howto.Question, howto.Answer)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	added := make([]uint64, 0, len(howtos))
	for rows.Next() {
		var id uint64
		if err := rows.Scan(&id); err != nil {
			continue
		}

		added = append(added, id)
	}

	return added, nil
}

func (repo *repo) UpdateHowto(ctx context.Context, howto howto.Howto) error {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Update howto")
	defer span.Finish()

	cols := repo.table.columns
	query := sqr.Update(repo.table.name).
		Where(sqr.Eq{cols.id: howto.Id}).
		Set(cols.courseId, howto.CourseId).
		Set(cols.question, howto.Question).
		Set(cols.answer, howto.Answer).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); err == nil && affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repo *repo) RemoveHowto(ctx context.Context, id uint64) error {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Remove howto")
	defer span.Finish()

	query := sqr.Delete(repo.table.name).
		Where(sqr.Eq{repo.table.columns.id: id}).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); err == nil && affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repo *repo) DescribeHowto(ctx context.Context, id uint64) (howto.Howto, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Describe howto")
	defer span.Finish()

	query := sqr.Select(repo.table.columns.ordered()...).
		From(repo.table.name).
		Where(sqr.Eq{repo.table.columns.id: id}).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	var result howto.Howto
	if err := query.QueryRowContext(ctx).Scan(howtoRows(&result)...); err != nil {
		return dummyHowto, err
	}

	return result, nil
}

func (repo *repo) ListHowtos(ctx context.Context, offset uint64, count uint64) ([]howto.Howto, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "List howtos")
	defer span.Finish()

	if count == 0 {
		return nil, nil
	}

	query := sqr.Select(repo.table.columns.ordered()...).
		From(repo.table.name).
		RunWith(repo.db).
		Limit(count).
		Offset(offset).
		PlaceholderFormat(repo.placeholder)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []howto.Howto
	for rows.Next() {
		var howto howto.Howto
		if err := rows.Scan(howtoRows(&howto)...); err != nil {
			continue
		}

		result = append(result, howto)
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, nil
}
