package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	actions  *prometheus.CounterVec
	requests *prometheus.CounterVec
)

const (
	label  string = "action"
	create string = "create"
	update string = "update"
	remove string = "remove"
)

func Register() {
	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "howto_requests",
			Help: "Number of requests for howto CUD actions.",
		},
		[]string{label},
	)
	actions = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "howto_successful_actions",
			Help: "Number of howto CUD actions successfully performed.",
		},
		[]string{label},
	)

	prometheus.MustRegister(requests, actions)
}

func incrementRequests(req string, times int) {
	if requests == nil {
		return
	}
	requests.With(prometheus.Labels{label: req}).Add(float64(times))
}

func incrementActions(act string, times int) {
	if actions == nil {
		return
	}
	actions.With(prometheus.Labels{label: act}).Add(float64(times))
}

func IncrementCreateRequests(times int) {
	incrementRequests(create, times)
}

func IncrementUpdateRequests(times int) {
	incrementRequests(update, times)
}

func IncrementRemoveRequests(times int) {
	incrementRequests(remove, times)
}

func IncrementCreates(times int) {
	incrementActions(create, times)
}

func IncrementUpdates(times int) {
	incrementActions(update, times)
}

func IncrementRemoves(times int) {
	incrementActions(remove, times)
}
