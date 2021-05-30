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

var errSome = errors.New("some error")

var _ = Describe("Flusher", func() {

	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		f        flusher.Flusher
		howtos   []howto.Howto
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		f = flusher.New(mockRepo)
		howtos = []howto.Howto{
			*howto.New(1, "question0", "answer0"),
			*howto.New(1, "question1", "answer1"),
			*howto.New(1, "question2", "answer2"),
			*howto.New(1, "question3", "answer3"),
			*howto.New(1, "question4", "answer4"),
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Flushed successfully", func() {
		It("", func() {
			mockRepo.EXPECT().AddHowto(gomock.Any()).Return(uint64(0), nil).MinTimes(1)
			Expect(f.Flush(howtos)).Should(BeEmpty())
		})
	})

	Context("Flushed partially", func() {
		It("", func() {
			index := 0
			failOn := 3
			mockRepo.EXPECT().AddHowto(gomock.Any()).DoAndReturn(
				func(h howto.Howto) (uint64, error) {
					if index == failOn {
						return 0, errSome
					}
					index++
					return h.Id, nil
				},
			).MinTimes(1)

			failed := f.Flush(howtos)
			Expect(failed).ShouldNot(BeEmpty())
			Expect(len(failed)).Should(BeEquivalentTo(len(howtos) - failOn))
		})
	})

	Context("Flush failed", func() {
		It("", func() {
			mockRepo.EXPECT().AddHowto(gomock.Any()).Return(uint64(0), errSome)
			failed := f.Flush(howtos)
			Expect(failed).Should(BeEquivalentTo(howtos))
		})
	})
})
