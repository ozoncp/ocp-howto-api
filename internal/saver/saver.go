package saver

import (
	"context"

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
	queue      []howto.Howto
	flusher    flusher.Flusher
	onOverflow OnOverflow
	alarm      <-chan struct{}
	add        chan howto.Howto
	close      chan struct{}
	done       chan struct{}
	context    context.Context
}

func NewSaver(
	ctx context.Context,
	capacity uint,
	onOverflow OnOverflow,
	flusher flusher.Flusher,
	alarmer alarmer.Alarmer,
) Saver {
	return &saver{
		capacity:   capacity,
		queue:      make([]howto.Howto, 0, capacity),
		flusher:    flusher,
		onOverflow: onOverflow,
		alarm:      alarmer.Alarm(),
		add:        make(chan howto.Howto),
		close:      make(chan struct{}),
		done:       make(chan struct{}),
		context:    ctx,
	}
}

func (saver *saver) Save(howto howto.Howto) {
	saver.add <- howto
}

func (saver *saver) Init() {
	go saver.poll()
}

func (saver *saver) Close() {
	saver.close <- struct{}{}
	<-saver.done
}

func (saver *saver) poll() {
	for {
		select {
		case howto := <-saver.add:
			saver.addToQueue(howto)
		case <-saver.alarm:
			saver.flush()
		case <-saver.close:
			close(saver.add)
			close(saver.close)
			saver.flush()
			saver.done <- struct{}{}
			return
		}
	}
}

func (saver *saver) flush() {
	if len(saver.queue) > 0 {
		saver.queue = saver.flusher.Flush(saver.context, saver.queue)
	}
}

func (saver *saver) addToQueue(howto howto.Howto) {
	if saver.isFull() {
		saver.resolveOverflow()
	}
	saver.queue = append(saver.queue, howto)
}

func (saver *saver) isFull() bool {
	return len(saver.queue) == int(saver.capacity)
}

func (saver *saver) resolveOverflow() {
	switch saver.onOverflow {
	case OnOverflowClearOldest:
		saver.queue = saver.queue[1:]

	case OnOverflowClearAll:
		fallthrough
	default:
		saver.queue = saver.queue[:0]
	}
}
