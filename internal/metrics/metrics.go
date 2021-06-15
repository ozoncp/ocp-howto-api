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
			Help: "Number of howto CUD actions",
		},
		[]string{label},
	)

	prometheus.MustRegister(counters)
}

func increment(action string, times uint64) {
	if counters == nil {
		return
	}
	counters.With(prometheus.Labels{label: action}).Add(float64(times))
}

func IncrementCreates(times uint64) {
	increment(create, times)
}

func IncrementUpdates(times uint64) {
	increment(update, times)
}

func IncrementRemoves(times uint64) {
	increment(remove, times)
}
