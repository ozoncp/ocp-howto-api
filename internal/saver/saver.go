package saver

import (
	"github.com/ozoncp/ocp-howto-api/internal/alarmer"
	"github.com/ozoncp/ocp-howto-api/internal/flusher"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
)

type OnOverflow int

const (
	OnOverflowClearAll OnOverflow = iota
	OnOverflowClearOldest
)

type Saver interface {
	Save(howto.Howto)
	Init()
	Close()
}

type saver struct {
	Saver
	capacity   uint
	toSave     []howto.Howto
	flusher    flusher.Flusher
	alarmer    alarmer.Alarmer
	onOverflow OnOverflow
}

func NewSaver(
	capacity uint,
	onOverflow OnOverflow,
	flusher flusher.Flusher,
	alarmer alarmer.Alarmer,
) Saver {
	return &saver{
		capacity:   capacity,
		toSave:     make([]howto.Howto, 0, capacity),
		flusher:    flusher,
		alarmer:    alarmer,
		onOverflow: onOverflow,
	}
}

func (saver saver) Save(howto howto.Howto) {
	if saver.isFull() {
		saver.resolveOverflow()
	}
	saver.toSave = append(saver.toSave, howto)
}

func (saver saver) Init() {
	go saver.poll()
}

func (saver saver) Close() {
	saver.save()
}

func (saver saver) isFull() bool {
	return len(saver.toSave) == int(saver.capacity)
}

func (saver saver) resolveOverflow() {
	switch saver.onOverflow {
	case OnOverflowClearOldest:
		saver.toSave = saver.toSave[1:]

	case OnOverflowClearAll:
		fallthrough
	default:
		saver.toSave = make([]howto.Howto, 0, saver.capacity)
	}
}

func (saver saver) poll() {
	saver.waitAlarm()
	saver.save()
}

func (saver saver) waitAlarm() {
	<-saver.alarmer.Alarm()
}

func (saver saver) save() {
	failed := saver.flusher.Flush(saver.toSave)
	saver.toSave = append(saver.toSave, failed...)
}
