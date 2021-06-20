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

// Register создает и регистрирует необходимые счетчики
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

// IncrementCreateRequests увеливает счетчик запросов на создание Howto
func IncrementCreateRequests(times int) {
	incrementRequests(create, times)
}

// IncrementUpdateRequests увеливает счетчик запросов на обновление Howto
func IncrementUpdateRequests(times int) {
	incrementRequests(update, times)
}

// IncrementRemoveRequests увеливает счетчик запросов на удаление Howto
func IncrementRemoveRequests(times int) {
	incrementRequests(remove, times)
}

// IncrementCreates увеливает счетчик успешных созданий Howto
func IncrementCreates(times int) {
	incrementActions(create, times)
}

// IncrementUpdates увеливает счетчик успешных обновлений Howto
func IncrementUpdates(times int) {
	incrementActions(update, times)
}

// IncrementRemoves увеливает счетчик успешных удалений Howto
func IncrementRemoves(times int) {
	incrementActions(remove, times)
}
