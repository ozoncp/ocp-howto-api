package api_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-howto-api/internal/api"
	"github.com/ozoncp/ocp-howto-api/internal/repo"

	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"
)

var _ = Describe("Api", func() {

	const (
		tableName       = "howtos"
		dummyId   int64 = 0
	)

	var (
		ctx    context.Context
		db     *sql.DB
		dbx    *sqlx.DB
		mock   sqlmock.Sqlmock
		server desc.OcpHowtoApiServer
	)

	BeforeEach(func() {
		ctx = context.Background()
		db, mock, _ = sqlmock.New()
		dbx = sqlx.NewDb(db, "sqlmock")
		server = api.NewOcpHowtoApi(repo.NewRepo(*dbx, 10))
	})

	AfterEach(func() {
		db.Close()
		dbx.Close()
	})

	Context("Create", func() {

		var (
			id       uint64
			request  *desc.CreateHowtoV1Request
			response *desc.CreateHowtoV1Response
			rows     *sqlmock.Rows
			query    *sqlmock.ExpectedQuery
			err      error
		)

		BeforeEach(func() {
			id = 12
			request = &desc.CreateHowtoV1Request{
				CourseId: 42,
				Question: "Question",
				Answer:   "Answer",
			}
			rows = sqlmock.NewRows([]string{"id"})
			query = mock.ExpectQuery(fmt.Sprintf("INSERT INTO %v", tableName)).
				WithArgs(request.CourseId, request.Question, request.Answer)
		})

		It("successfully", func() {
			query.WillReturnRows(rows.AddRow(id))
			response, err = server.CreateHowtoV1(ctx, request)
			Expect(response.Id).Should(BeEquivalentTo(id))
			Expect(err).Should(BeNil())
		})

		It("failed", func() {
			query.WillReturnError(errors.New(""))
			response, err = server.CreateHowtoV1(ctx, request)
			Expect(response.Id).Should(BeEquivalentTo(dummyId))
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("MultiCreate", func() {

		var (
			request  *desc.MultiCreateHowtoV1Request
			response *desc.MultiCreateHowtoV1Response
			exec     *sqlmock.ExpectedExec
			err      error
		)

		BeforeEach(func() {
			request = &desc.MultiCreateHowtoV1Request{
				Howtos: []*desc.Howto{{}, {}, {}},
			}
			args := []driver.Value{}
			for i := 0; i < len(request.Howtos); i++ {
				h := request.Howtos[i]
				args = append(args, h.CourseId, h.Question, h.Answer)
			}
			exec = mock.ExpectExec(fmt.Sprintf("INSERT INTO %v", tableName)).WithArgs(args...)
		})

		It("successfully", func() {
			exec.WillReturnResult(sqlmock.NewResult(dummyId, int64(len(request.Howtos))))
			response, err = server.MultiCreateHowtoV1(ctx, request)
			Expect(response.Added).Should(BeEquivalentTo(len(request.Howtos)))
			Expect(err).Should(BeNil())
		})

		It("partially", func() {
			added := int64(1)
			exec.WillReturnResult(sqlmock.NewResult(dummyId, added))
			response, err = server.MultiCreateHowtoV1(ctx, request)
			Expect(response.Added).Should(BeEquivalentTo(added))
			Expect(err).Should(BeNil())
		})

		It("failed", func() {
			exec.WillReturnError(errors.New(""))
			response, err = server.MultiCreateHowtoV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Update", func() {

		var (
			request  *desc.UpdateHowtoV1Request
			exec     *sqlmock.ExpectedExec
			affected int64
			err      error
		)

		BeforeEach(func() {
			h := desc.Howto{
				Id:       12,
				CourseId: 42,
				Question: "Question",
				Answer:   "Answer",
			}
			request = &desc.UpdateHowtoV1Request{Howto: &h}
			exec = mock.ExpectExec(fmt.Sprintf("UPDATE %v", tableName)).
				WithArgs(h.CourseId, h.Question, h.Answer, h.Id)
		})

		It("successfully", func() {
			affected = 1
			exec.WillReturnResult(sqlmock.NewResult(dummyId, affected))
			_, err = server.UpdateHowtoV1(ctx, request)
			Expect(err).Should(BeNil())
		})

		It("not found", func() {
			affected = 0
			exec.WillReturnResult(sqlmock.NewResult(dummyId, affected))
			_, err = server.UpdateHowtoV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("failed", func() {
			exec.WillReturnError(errors.New(""))
			_, err = server.UpdateHowtoV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Remove", func() {
		var (
			request  *desc.RemoveHowtoV1Request
			exec     *sqlmock.ExpectedExec
			affected int64
			err      error
		)

		BeforeEach(func() {
			request = &desc.RemoveHowtoV1Request{
				Id: 1,
			}
			exec = mock.ExpectExec(fmt.Sprintf("DELETE FROM %v", tableName)).WithArgs(request.Id)
		})

		It("successfully", func() {
			affected = 1
			exec.WillReturnResult(sqlmock.NewResult(dummyId, affected))
			_, err = server.RemoveHowtoV1(ctx, request)
			Expect(err).Should(BeNil())
		})

		It("not found", func() {
			affected = 0
			exec.WillReturnResult(sqlmock.NewResult(dummyId, affected))
			_, err = server.RemoveHowtoV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("failed", func() {
			exec.WillReturnError(errors.New(""))
			_, err = server.RemoveHowtoV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("Describe", func() {

		var (
			row      desc.Howto
			request  *desc.DescribeHowtoV1Request
			response *desc.DescribeHowtoV1Response
			rows     *sqlmock.Rows
			query    *sqlmock.ExpectedQuery
			err      error
		)

		BeforeEach(func() {
			row = desc.Howto{
				Id:       5,
				CourseId: 10,
				Question: "Question",
				Answer:   "Answer",
			}
			request = &desc.DescribeHowtoV1Request{
				Id: row.Id,
			}
			rows = sqlmock.NewRows([]string{"id", "course_id", "question", "answer"})
			query = mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %v", tableName)).
				WithArgs(request.Id)
		})

		It("successfully", func() {
			rows.AddRow(row.Id, row.CourseId, row.Question, row.Answer)
			query.WillReturnRows(rows)
			response, err = server.DescribeHowtoV1(ctx, request)
			Expect(response.Howto.Id).Should(BeEquivalentTo(row.Id))
			Expect(response.Howto.CourseId).Should(BeEquivalentTo(row.CourseId))
			Expect(response.Howto.Question).Should(BeEquivalentTo(row.Question))
			Expect(response.Howto.Answer).Should(BeEquivalentTo(row.Answer))
			Expect(err).Should(BeNil())
		})

		It("failed", func() {
			query.WillReturnError(errors.New(""))
			response, err = server.DescribeHowtoV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("List", func() {

		var (
			request  *desc.ListHowtosV1Request
			response *desc.ListHowtosV1Response
			rows     *sqlmock.Rows
			query    *sqlmock.ExpectedQuery
			err      error
		)

		BeforeEach(func() {
			request = &desc.ListHowtosV1Request{
				Offset: 0,
				Count:  5,
			}
			rows = sqlmock.NewRows([]string{"id", "course_id", "question", "answer"})
			query = mock.ExpectQuery(
				fmt.Sprintf("SELECT (.+) FROM %v LIMIT %d OFFSET %d",
					tableName, request.Count, request.Offset,
				))
		})

		It("successfully", func() {
			for i := 0; i < int(request.Count); i++ {
				rows.AddRow(dummyId, dummyId, "", "")
			}
			query.WillReturnRows(rows)
			response, err = server.ListHowtosV1(ctx, request)
			Expect(len(response.Howtos)).Should(BeEquivalentTo(request.Count))
			Expect(err).Should(BeNil())
		})

		It("failed", func() {
			query.WillReturnError(errors.New(""))
			response, err = server.ListHowtosV1(ctx, request)
			Expect(len(response.Howtos)).Should(BeEquivalentTo(0))
			Expect(err).ShouldNot(BeNil())
		})
	})

})
