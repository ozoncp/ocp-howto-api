package alarmer_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-howto-api/internal/alarmer"
)

var _ = Describe("Alarmer", func() {

	Context("PeriodAlarmer", func() {

		It("alarms successfully", func() {
			var period uint = 100
			timeout := 1500

			alarmer := alarmer.NewPeriodAlarmer(period)
			alarmer.Init()
			defer alarmer.Close()

			timeoutTimer := time.NewTimer(time.Duration(timeout) * time.Millisecond)
			select {
			case <-alarmer.Alarm():
				break
			case <-timeoutTimer.C:
				Fail("Timeout")
			}
		})

		It("alarms periodically", func() {
			var period uint = 10
			numOfAlarms := 50

			alarmer := alarmer.NewPeriodAlarmer(period)
			alarmer.Init()
			defer alarmer.Close()

			startTime := time.Now()
			for i := 0; i < numOfAlarms; i++ {
				<-alarmer.Alarm()
			}
			realDuration := time.Since(startTime).Milliseconds()
			expectedDuration := period * uint(numOfAlarms)

			Expect(realDuration).Should(BeNumerically(">=", expectedDuration))
		})

		It("closes successfully", func() {
			alarmer := alarmer.NewPeriodAlarmer(100)
			alarmer.Init()
			<-alarmer.Alarm()
			alarmer.Close()
			Eventually(alarmer.Alarm()).Should(BeClosed())
		})
	})
})
