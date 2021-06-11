package repo

import (
	"context"

	"github.com/ozoncp/ocp-howto-api/internal/howto"

	sqr "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

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
				id:        "id",
				course_id: "course_id",
				question:  "question",
				answer:    "answer",
			},
		},
		placeholder: sqr.Dollar,
	}
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
		Columns(cols.course_id, cols.question, cols.answer).
		Values(howto.CourseId, howto.Question, howto.Answer).
		Suffix("RETURNING \"$\"", cols.id).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
	if err := query.QueryRowContext(ctx).Scan(&howto.Id); err != nil {
		return 0, err
	}

	return howto.Id, nil
}

func (repo *repo) AddHowtos(howtos []howto.Howto) error {

	cols := repo.table.columns
	query := sqr.Insert(repo.table.name).
		Columns(cols.course_id, cols.question, cols.answer).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	for _, howto := range howtos {
		query = query.Values(howto.CourseId, howto.Question, howto.Answer)
	}

	ctx := context.TODO()
	_, err := query.ExecContext(ctx)
	return err
}

func (repo *repo) RemoveHowto(id uint64) error {

	query := sqr.Delete(repo.table.name).
		Where(sqr.Eq{repo.table.columns.id: id}).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
	_, err := query.ExecContext(ctx)
	return err
}

func (repo *repo) DescribeHowto(id uint64) (howto.Howto, error) {

	query := sqr.Select(repo.table.columns.ordered()...).
		From(repo.table.name).
		Where(sqr.Eq{repo.table.columns.id: id}).
		RunWith(repo.db).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
	var result howto.Howto
	err := query.QueryRowContext(ctx).Scan(&result)
	return result, err
}

func (repo *repo) ListHowtos(startWith uint64, count uint64) ([]howto.Howto, error) {

	query := sqr.Select(repo.table.columns.ordered()...).
		From(repo.table.name).
		RunWith(repo.db).
		Limit(count).
		Offset(startWith).
		PlaceholderFormat(repo.placeholder)

	ctx := context.TODO()
	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var howtos []howto.Howto
	for rows.Next() {
		var howto howto.Howto
		if err := sqlx.StructScan(rows, &howto); err != nil {
			continue
		}

		howtos = append(howtos, howto)
	}
	return howtos, nil
}
