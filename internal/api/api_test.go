package api_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-howto-api/internal/api"
	"github.com/ozoncp/ocp-howto-api/internal/repo"

	desc "github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api"

	"github.com/ozoncp/ocp-howto-api/internal/mocks"
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
		ctrl   *gomock.Controller
		prod   *mocks.MockProducer
		params *desc.HowtoParams
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		prod = mocks.NewMockProducer(ctrl)
		db, mock, _ = sqlmock.New()
		dbx = sqlx.NewDb(db, "sqlmock")
		server = api.NewOcpHowtoApi(repo.NewRepo(*dbx, 10), prod)
		params = &desc.HowtoParams{
			CourseId: 42,
			Question: "Question",
			Answer:   "Answer",
		}
	})

	AfterEach(func() {
		db.Close()
		dbx.Close()
		ctrl.Finish()
	})

	Context("Invalid arguments", func() {
		It("Create", func() {
			response, err := server.CreateHowtoV1(ctx, &desc.CreateHowtoV1Request{
				Params: &desc.HowtoParams{CourseId: 123, Question: "Question", Answer: ""},
			})
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
		It("MultiCreate", func() {
			response, err := server.MultiCreateHowtoV1(ctx, &desc.MultiCreateHowtoV1Request{})
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
		It("Update", func() {
			response, err := server.UpdateHowtoV1(ctx, &desc.UpdateHowtoV1Request{})
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
		It("List", func() {
			response, err := server.ListHowtosV1(ctx, &desc.ListHowtosV1Request{Offset: 5, Count: 0})
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
		It("Remove", func() {
			response, err := server.RemoveHowtoV1(ctx, &desc.RemoveHowtoV1Request{Id: 0})
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
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
			request = &desc.CreateHowtoV1Request{Params: params}
			rows = sqlmock.NewRows([]string{"id"})
			query = mock.ExpectQuery(fmt.Sprintf("INSERT INTO %v", tableName)).
				WithArgs(params.CourseId, params.Question, params.Answer)
		})

		It("successfully", func() {
			prod.EXPECT().SendEvent(gomock.Any()).Times(1)
			query.WillReturnRows(rows.AddRow(id))
			response, err = server.CreateHowtoV1(ctx, request)
			Expect(response.Id).Should(BeEquivalentTo(id))
			Expect(err).Should(BeNil())
		})

		It("failed", func() {
			query.WillReturnError(errors.New(""))
			response, err = server.CreateHowtoV1(ctx, request)
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("MultiCreate", func() {

		var (
			request  *desc.MultiCreateHowtoV1Request
			response *desc.MultiCreateHowtoV1Response
			query    *sqlmock.ExpectedQuery
			rows     *sqlmock.Rows
			err      error
		)

		BeforeEach(func() {
			request = &desc.MultiCreateHowtoV1Request{
				Params: []*desc.HowtoParams{params, params, params},
			}
			args := []driver.Value{}
			for i := 0; i < len(request.Params); i++ {
				h := request.Params[i]
				args = append(args, h.CourseId, h.Question, h.Answer)
			}
			rows = sqlmock.NewRows([]string{"id"})
			query = mock.ExpectQuery(fmt.Sprintf("INSERT INTO %v", tableName)).WithArgs(args...)
		})

		It("successfully", func() {
			for i := 0; i < len(request.Params); i++ {
				rows.AddRow(dummyId)
			}
			query.WillReturnRows(rows)
			prod.EXPECT().SendEvent(gomock.Any()).Times(1)
			response, err = server.MultiCreateHowtoV1(ctx, request)
			Expect(len(response.Ids)).Should(BeEquivalentTo(len(request.Params)))
			Expect(err).Should(BeNil())
		})

		It("partially", func() {
			added := 1
			for i := 0; i < added; i++ {
				rows.AddRow(dummyId)
			}
			query.WillReturnRows(rows)
			prod.EXPECT().SendEvent(gomock.Any()).Times(1)
			response, err = server.MultiCreateHowtoV1(ctx, request)
			Expect(len(response.Ids)).Should(BeEquivalentTo(added))
			Expect(err).Should(BeNil())
		})

		It("failed", func() {
			query.WillReturnError(errors.New(""))
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
				Id:     12,
				Params: params,
			}
			request = &desc.UpdateHowtoV1Request{Howto: &h}
			exec = mock.ExpectExec(fmt.Sprintf("UPDATE %v", tableName)).
				WithArgs(params.CourseId, params.Question, params.Answer, h.Id)
		})

		It("successfully", func() {
			prod.EXPECT().SendEvent(gomock.Any()).Times(1)
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
			request = &desc.RemoveHowtoV1Request{Id: 1}
			exec = mock.ExpectExec(fmt.Sprintf("DELETE FROM %v", tableName)).WithArgs(request.Id)
		})

		It("successfully", func() {
			prod.EXPECT().SendEvent(gomock.Any()).Times(1)
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
				Id:     5,
				Params: params,
			}
			request = &desc.DescribeHowtoV1Request{
				Id: row.Id,
			}
			rows = sqlmock.NewRows([]string{"id", "course_id", "question", "answer"})
			query = mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %v", tableName)).
				WithArgs(request.Id)
		})

		It("successfully", func() {
			rows.AddRow(row.Id, params.CourseId, params.Question, params.Answer)
			query.WillReturnRows(rows)
			response, err = server.DescribeHowtoV1(ctx, request)
			rParams := response.Howto.Params
			Expect(response.Howto.Id).Should(BeEquivalentTo(row.Id))
			Expect(rParams.CourseId).Should(BeEquivalentTo(params.CourseId))
			Expect(rParams.Question).Should(BeEquivalentTo(params.Question))
			Expect(rParams.Answer).Should(BeEquivalentTo(params.Answer))
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
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})

})
