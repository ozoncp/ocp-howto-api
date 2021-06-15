package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/utils"

	sqr "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var dummyHowto = howto.Howto{}

type Repo interface {
	AddHowto(context.Context, howto.Howto) (uint64, error)
	AddHowtos(context.Context, []howto.Howto) (uint64, error)
	UpdateHowto(context.Context, howto.Howto) error
	RemoveHowto(ctx context.Context, id uint64) error
	DescribeHowto(ctx context.Context, id uint64) (howto.Howto, error)
	ListHowtos(ctx context.Context, offset uint64, count uint64) ([]howto.Howto, error)
}

func NewRepo(db sqlx.DB) Repo {
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
		batchSize:   2,
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

func (repo *repo) AddHowto(ctx context.Context, howto howto.Howto) (uint64, error) {

	cols := repo.table.columns
	query := sqr.Insert(repo.table.name).
		Columns(cols.courseId, cols.question, cols.answer).
		Values(howto.CourseId, howto.Question, howto.Answer).
		Suffix(fmt.Sprintf("RETURNING %v", cols.id)).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	if err := query.QueryRowContext(ctx).Scan(&howto.Id); err != nil {
		return dummyHowto.Id, err
	}

	return howto.Id, nil
}

func (repo *repo) AddHowtos(ctx context.Context, howtos []howto.Howto) (uint64, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Add howtos")
	defer span.Finish()

	if len(howtos) == 0 {
		return 0, nil
	}

	added := uint64(0)
	batches := utils.SplitToBulks(howtos, repo.batchSize)
	for _, batch := range batches {
		inserted, err := repo.insertBatch(ctx, batch)
		added += uint64(inserted)
		if err != nil {
			return added, err
		}
	}

	return added, nil
}

func (repo *repo) insertBatch(ctx context.Context, howtos []howto.Howto) (int64, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Insert howtos batch")
	defer span.Finish()

	cols := repo.table.columns
	query := sqr.Insert(repo.table.name).
		Columns(cols.courseId, cols.question, cols.answer).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	for _, howto := range howtos {
		query = query.Values(howto.CourseId, howto.Question, howto.Answer)
	}

	result, err := query.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	if added, err := result.RowsAffected(); err == nil {
		return added, nil
	}

	return int64(len(howtos)), nil
}

func (repo *repo) UpdateHowto(ctx context.Context, howto howto.Howto) error {

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
		return errors.New("row not found")
	}
	return nil
}

func (repo *repo) RemoveHowto(ctx context.Context, id uint64) error {

	query := sqr.Delete(repo.table.name).
		Where(sqr.Eq{repo.table.columns.id: id}).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); err == nil && affected == 0 {
		return errors.New("row not found")
	}
	return nil
}

func (repo *repo) DescribeHowto(ctx context.Context, id uint64) (howto.Howto, error) {

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

	var result []howto.Howto

	query := sqr.Select(repo.table.columns.ordered()...).
		From(repo.table.name).
		RunWith(repo.db).
		Limit(count).
		Offset(offset).
		PlaceholderFormat(repo.placeholder)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var howto howto.Howto
		if err := rows.Scan(howtoRows(&howto)...); err != nil {
			continue
		}

		result = append(result, howto)
	}
	return result, nil
}
