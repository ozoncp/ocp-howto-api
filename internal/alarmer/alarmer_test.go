package alarmer_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-howto-api/internal/alarmer"
)

func milliseconds(milliseconds time.Duration) time.Duration {
	return milliseconds * time.Millisecond
}

var _ = Describe("Alarmer", func() {

	Context("PeriodAlarmer", func() {

		It("alarms successfully", func() {
			period := milliseconds(100)
			timeout := milliseconds(150)

			alarmer := alarmer.NewPeriodAlarmer(period)
			alarmer.Init()
			defer alarmer.Close()

			timeoutExceeded := time.After(timeout)
			select {
			case <-alarmer.Alarm():
			case <-timeoutExceeded:
				Fail("Timeout exceeded. Alarmer does not alarm")
			}
		})

		It("alarms periodically", func() {
			period := milliseconds(10)
			numOfAlarms := 50

			alarmer := alarmer.NewPeriodAlarmer(period)
			alarmer.Init()
			defer alarmer.Close()

			startTime := time.Now()
			for i := 0; i < numOfAlarms; i++ {
				<-alarmer.Alarm()
			}
			realDuration := time.Since(startTime)
			expectedDuration := int(period) * numOfAlarms

			Expect(realDuration).Should(BeNumerically(">=", expectedDuration))
		})

		It("closes successfully", func() {
			alarmer := alarmer.NewPeriodAlarmer(milliseconds(100))
			alarmer.Init()
			<-alarmer.Alarm()
			alarmer.Close()
			Eventually(alarmer.Alarm()).Should(BeClosed())
		})
	})
})
