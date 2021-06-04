package alarmer

import "time"

type Alarmer interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

type periodAlarmer struct {
	Alarmer
	periodMs uint
	alarm    chan struct{}
}

func NewPeriodAlarmer(periodMs uint) Alarmer {
	return &periodAlarmer{
		periodMs: periodMs,
		alarm:    make(chan struct{}),
	}
}

func (alarmer periodAlarmer) Alarm() <-chan struct{} {
	return alarmer.alarm
}

func (alarmer periodAlarmer) Init() {
	go alarmer.poll()
}

func (alarmer periodAlarmer) Close() {
	close(alarmer.alarm)
}

func (alarmer periodAlarmer) poll() {
	alarmer.waitPeriod()
	alarmer.alarm <- struct{}{}
}

func (alarmer periodAlarmer) waitPeriod() {
	time.Sleep(time.Millisecond * time.Duration(alarmer.periodMs))
}
