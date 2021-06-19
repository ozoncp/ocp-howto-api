package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	counters *prometheus.CounterVec
	label    string = "action"
	create   string = "create"
	update   string = "update"
	remove   string = "remove"
)

func Register() {
	counters = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "howto_actions",
			Help: "Number of howto CUD actions successfully performed.",
		},
		[]string{label},
	)

	prometheus.MustRegister(counters)
}

func increment(action string, times int) {
	if counters == nil {
		return
	}
	counters.With(prometheus.Labels{label: action}).Add(float64(times))
}

func IncrementCreates(times int) {
	increment(create, times)
}

func IncrementUpdates(times int) {
	increment(update, times)
}

func IncrementRemoves(times int) {
	increment(remove, times)
}
