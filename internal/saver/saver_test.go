package saver_test

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-howto-api/internal/alarmer"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/mocks"
	"github.com/ozoncp/ocp-howto-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		alarm       alarmer.Alarmer
		saved       []howto.Howto
		period      time.Duration = 200 * time.Millisecond
		toSave      howto.Howto   = howto.Howto{}
	)

	waitAlarms := func(times int) {
		for i := 0; i < times; i++ {
			<-alarm.Alarm()
		}
	}

	min := func(lhs int, rhs int) int {
		if lhs < rhs {
			return lhs
		}
		return rhs
	}

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		alarm = alarmer.NewPeriodAlarmer(period)
		alarm.Init()
		saved = make([]howto.Howto, 0)

		mockFlusher.EXPECT().Flush(gomock.Any()).MinTimes(1).DoAndReturn(
			func(howtos []howto.Howto) []howto.Howto {
				saved = append(saved, howtos...)
				return nil
			})
	})

	AfterEach(func() {
		alarm.Close()
		ctrl.Finish()
	})

	Context("", func() {
		It("On overflow clears only one", func() {
			capacity := 3
			numSaves := 1000
			saver := saver.NewSaver(uint(capacity), saver.OnOverflowClearOldest, mockFlusher, alarm)
			defer saver.Close()
			saver.Init()
			for i := 0; i < numSaves; i++ {
				saver.Save(toSave)
			}
			waitAlarms(2)

			shouldSave := min(numSaves, capacity)
			Expect(len(saved)).Should(BeEquivalentTo(shouldSave))
		})

		It("On overflow clears all", func() {
			capacity := 3
			numSaves := 1000
			saver := saver.NewSaver(uint(capacity), saver.OnOverflowClearAll, mockFlusher, alarm)
			defer saver.Close()
			saver.Init()
			for i := 0; i < numSaves; i++ {
				saver.Save(toSave)
			}
			waitAlarms(2)

			shouldSave := min(numSaves%capacity, numSaves)
			if shouldSave == 0 {
				shouldSave = capacity
			}

			Expect(len(saved)).Should(BeEquivalentTo(shouldSave))
		})

		It("Saves on close", func() {
			hourAlarm := alarmer.NewPeriodAlarmer(time.Hour)
			hourAlarm.Init()
			defer hourAlarm.Close()
			saver := saver.NewSaver(10, saver.OnOverflowClearAll, mockFlusher, hourAlarm)
			saver.Init()
			saver.Save(toSave)
			saver.Close()
			time.Sleep(time.Millisecond)
			Expect(len(saved)).Should(BeNumerically(">", 0))
		})

		It("Panics if closed", func() {
			mockFlusher.Flush(nil)
			saver := saver.NewSaver(10, saver.OnOverflowClearAll, mockFlusher, alarm)
			saver.Init()
			saver.Close()
			save := func() {
				saver.Save(toSave)
			}
			Expect(save).Should(Panic())
		})
	})
})
