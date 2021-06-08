package alarmer

import "time"

type Alarmer interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

type periodAlarmer struct {
	Alarmer
	period time.Duration
	alarm  chan struct{}
	close  chan struct{}
}

func NewPeriodAlarmer(period time.Duration) Alarmer {
	return &periodAlarmer{
		period: period,
		alarm:  make(chan struct{}),
		close:  make(chan struct{}),
	}
}

func (alarmer *periodAlarmer) Alarm() <-chan struct{} {
	return alarmer.alarm
}

func (alarmer *periodAlarmer) Init() {
	go alarmer.poll()
}

func (alarmer *periodAlarmer) Close() {
	alarmer.close <- struct{}{}
}

func (alarmer *periodAlarmer) poll() {
	timer := time.NewTimer(alarmer.period)
	for {
		select {
		case <-alarmer.close:
			close(alarmer.alarm)
			close(alarmer.close)
			return
		case <-timer.C:
			alarmer.tick()
			timer.Reset(alarmer.period)
		}
	}
}

func (alarmer *periodAlarmer) tick() {
	select {
	case alarmer.alarm <- struct{}{}:
	default:
	}
}
