package saver_test

import (
	"context"
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
		ctx         context.Context
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		alarm       alarmer.Alarmer
		hourAlarm   alarmer.Alarmer
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
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		alarm = alarmer.NewPeriodAlarmer(period)
		hourAlarm = alarmer.NewPeriodAlarmer(time.Hour)
		alarm.Init()
		hourAlarm.Init()
		saved = make([]howto.Howto, 0)
	})

	AfterEach(func() {
		hourAlarm.Close()
		alarm.Close()
		ctrl.Finish()
	})

	Context("", func() {

		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(ctx, gomock.Any()).MinTimes(1).DoAndReturn(
				func(ctx context.Context, howtos []howto.Howto) []howto.Howto {
					saved = append(saved, howtos...)
					return nil
				})
		})

		It("on overflow clears only one", func() {
			capacity := 3
			numSaves := 1000
			saver := saver.NewSaver(ctx, uint(capacity), saver.OnOverflowClearOldest, mockFlusher, alarm)
			defer saver.Close()
			saver.Init()
			for i := 0; i < numSaves; i++ {
				saver.Save(toSave)
			}
			waitAlarms(2)

			shouldSave := min(numSaves, capacity)
			Expect(len(saved)).Should(BeEquivalentTo(shouldSave))
		})

		It("on overflow clears all", func() {
			capacity := 3
			numSaves := 1000
			saver := saver.NewSaver(ctx, uint(capacity), saver.OnOverflowClearAll, mockFlusher, alarm)
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

		It("saves on close", func() {
			saver := saver.NewSaver(ctx, 10, saver.OnOverflowClearAll, mockFlusher, hourAlarm)
			saver.Init()
			saver.Save(toSave)
			saver.Close()
			Expect(len(saved)).Should(BeNumerically(">", 0))
		})

		It("panics if closed", func() {
			mockFlusher.Flush(ctx, nil)
			saver := saver.NewSaver(ctx, 10, saver.OnOverflowClearAll, mockFlusher, alarm)
			saver.Init()
			saver.Close()
			save := func() {
				saver.Save(toSave)
			}
			Expect(save).Should(Panic())
		})
	})

	Context("Clear() blocks until flush finished", func() {
		It("", func() {
			finished := false
			mockFlusher.EXPECT().Flush(ctx, gomock.Any()).MinTimes(1).DoAndReturn(
				func(ctx context.Context, howtos []howto.Howto) []howto.Howto {
					time.Sleep(time.Second * 3)
					finished = true
					return nil
				})

			saver := saver.NewSaver(ctx, 1, saver.OnOverflowClearAll, mockFlusher, hourAlarm)
			saver.Init()
			saver.Save(toSave)
			saver.Close()
			Expect(finished).Should(BeEquivalentTo(true))
		})
	})
})
