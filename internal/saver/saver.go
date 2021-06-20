package saver

import (
	"context"

	"github.com/ozoncp/ocp-howto-api/internal/alarmer"
	"github.com/ozoncp/ocp-howto-api/internal/flusher"
	"github.com/ozoncp/ocp-howto-api/internal/howto"
)

// OnOverflow - alias для описания поведения в случае переполнения очереди на сохранение
type OnOverflow int

const (
	// OnOverflowClearAll - при переполнении удалять всю очередь
	OnOverflowClearAll OnOverflow = iota

	// OnOverflowClearAll - при переполнении удалять только самую старую сущность
	OnOverflowClearOldest
)

// Saver - интерфейс для асинхронного сохранения сущностей
type Saver interface {
	// Save добавляет сущность в очередь на сохранение
	Save(howto.Howto)

	// Init запускает обработку сохранения сущностей
	Init()

	// Close дожидается сохранения всех сущностей и останаливает обработку
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

// NewSaver создает новый экземпляр Saver
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
