package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-howto-api/internal/flusher"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/mocks"
)

var errDummy = errors.New("some error")
var dummyId = uint64(0)

func newHowto(course uint64, q string, a string) howto.Howto {
	return howto.Howto{
		Id:       dummyId,
		CourseId: course,
		Question: q,
		Answer:   q,
	}
}

var _ = Describe("Flusher", func() {

	var (
		ctx           context.Context
		ctrl          *gomock.Controller
		mockRepo      *mocks.MockRepo
		f             flusher.Flusher
		toFlush       []howto.Howto
		failedToFlush []howto.Howto
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		f = flusher.New(mockRepo)
		toFlush = []howto.Howto{
			newHowto(1, "question0", "answer0"),
			newHowto(1, "question1", "answer1"),
			newHowto(1, "question2", "answer2"),
			newHowto(1, "question3", "answer3"),
			newHowto(1, "question4", "answer4"),
		}
	})

	JustBeforeEach(func() {
		failedToFlush = f.Flush(ctx, toFlush)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Flushed successfully", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddHowto(ctx, gomock.Any()).Return(dummyId, nil).MinTimes(1)
		})
		It("", func() {
			Expect(failedToFlush).Should(BeEmpty())
		})
	})

	Context("Flushed partially", func() {
		succeeded := 2
		BeforeEach(func() {
			mockRepo.EXPECT().AddHowto(ctx, gomock.Any()).Return(dummyId, nil).Times(succeeded)
			mockRepo.EXPECT().AddHowto(ctx, gomock.Any()).Return(dummyId, errDummy)
		})
		It("", func() {
			Expect(len(failedToFlush)).Should(BeEquivalentTo(len(toFlush) - succeeded))
		})
	})

	Context("Flush failed", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddHowto(ctx, gomock.Any()).Return(dummyId, errDummy)
		})
		It("", func() {
			Expect(failedToFlush).Should(BeEquivalentTo(toFlush))
		})
	})
})
