package flusher_test

import (
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

var _ = Describe("Flusher", func() {

	var (
		ctrl          *gomock.Controller
		mockRepo      *mocks.MockRepo
		f             flusher.Flusher
		toFlush       []howto.Howto
		failedToFlush []howto.Howto
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		f = flusher.New(mockRepo)
		toFlush = []howto.Howto{
			*howto.New(1, "question0", "answer0"),
			*howto.New(1, "question1", "answer1"),
			*howto.New(1, "question2", "answer2"),
			*howto.New(1, "question3", "answer3"),
			*howto.New(1, "question4", "answer4"),
		}
	})

	JustBeforeEach(func() {
		failedToFlush = f.Flush(toFlush)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Flushed successfully", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddHowto(gomock.Any()).Return(dummyId, nil).MinTimes(1)
		})
		It("", func() {
			Expect(f.Flush(toFlush)).Should(BeEmpty())
		})
	})

	Context("Flushed partially", func() {
		succeeded := 2
		BeforeEach(func() {
			mockRepo.EXPECT().AddHowto(gomock.Any()).Return(dummyId, nil).Times(succeeded)
			mockRepo.EXPECT().AddHowto(gomock.Any()).Return(dummyId, errDummy)
		})
		It("", func() {
			Expect(len(failedToFlush)).Should(BeEquivalentTo(len(toFlush) - succeeded))
		})
	})

	Context("Flush failed", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddHowto(gomock.Any()).Return(dummyId, errDummy)
		})
		It("", func() {
			Expect(failedToFlush).Should(BeEquivalentTo(toFlush))
		})
	})
})
