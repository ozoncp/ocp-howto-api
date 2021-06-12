package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-howto-api/internal/howto"

	sqr "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var dummyHowto = howto.Howto{}

type Repo interface {
	AddHowto(howto.Howto) (uint64, error)
	AddHowtos([]howto.Howto) error
	RemoveHowto(id uint64) error
	DescribeHowto(id uint64) (howto.Howto, error)
	ListHowtos(startWith uint64, count uint64) ([]howto.Howto, error)
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
}

func (repo *repo) AddHowto(howto howto.Howto) (uint64, error) {

	cols := repo.table.columns
	query := sqr.Insert(repo.table.name).
		Columns(cols.courseId, cols.question, cols.answer).
		Values(howto.CourseId, howto.Question, howto.Answer).
		Suffix(fmt.Sprintf("RETURNING %v", cols.id)).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
	if err := query.QueryRowContext(ctx).Scan(&howto.Id); err != nil {
		return dummyHowto.Id, err
	}

	return howto.Id, nil
}

func (repo *repo) AddHowtos(howtos []howto.Howto) error {
	if len(howtos) == 0 {
		return nil
	}

	cols := repo.table.columns
	query := sqr.Insert(repo.table.name).
		Columns(cols.courseId, cols.question, cols.answer).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	for _, howto := range howtos {
		query = query.Values(howto.CourseId, howto.Question, howto.Answer)
	}

	ctx := context.TODO()
	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	expected := int64(len(howtos))
	if added, err := result.RowsAffected(); err == nil && added != expected {
		return fmt.Errorf("only %v rows out of %v have been added", added, expected)
	}

	return nil
}

func (repo *repo) RemoveHowto(id uint64) error {

	query := sqr.Delete(repo.table.name).
		Where(sqr.Eq{repo.table.columns.id: id}).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); err == nil && affected == 0 {
		return errors.New("row not found")
	}
	return nil
}

func (repo *repo) DescribeHowto(id uint64) (howto.Howto, error) {

	query := sqr.Select(repo.table.columns.ordered()...).
		From(repo.table.name).
		Where(sqr.Eq{repo.table.columns.id: id}).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
	var result howto.Howto
	if err := query.QueryRowContext(ctx).Scan(howtoRows(&result)...); err != nil {
		return dummyHowto, err
	}

	return result, nil
}

func (repo *repo) ListHowtos(startWith uint64, count uint64) ([]howto.Howto, error) {

	var result []howto.Howto

	query := sqr.Select(repo.table.columns.ordered()...).
		From(repo.table.name).
		RunWith(repo.db).
		Limit(count).
		Offset(startWith).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
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
