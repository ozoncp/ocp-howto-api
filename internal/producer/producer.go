package producer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	log "github.com/rs/zerolog/log"
)

type EventType = uint64

const (
	EventTypeCreated EventType = iota
	EventTypeUpdated
	EventTypeRemoved
)

type Event struct {
	Type EventType
	Body map[string]interface{}
}

type Producer interface {
	Init()
	SendEvent(Event)
	Close()
}

type producer struct {
	Producer
	prod   sarama.SyncProducer
	topic  string
	events chan Event
	close  chan struct{}
	done   chan struct{}
}

func New(broker string, topic string, capacity uint) (Producer, error) {

	brokers := []string{broker}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		prod:   prod,
		topic:  topic,
		events: make(chan Event, capacity),
		close:  make(chan struct{}),
		done:   make(chan struct{}),
	}, nil
}

func (p *producer) Init() {
	go p.poll()
}

func (p *producer) SendEvent(event Event) {
	p.events <- event
}

func (p *producer) Close() {
	p.close <- struct{}{}
	<-p.done
}

func (p *producer) poll() {
	for {
		select {
		case event := <-p.events:
			p.send(event)
		case <-p.close:
			close(p.events)
			p.flush()
			p.prod.Close()
			p.done <- struct{}{}
			return
		}
	}
}

func (p *producer) flush() {
	for event := range p.events {
		p.send(event)
	}
}

func (p *producer) send(event Event) {

	json, err := json.Marshal(event)
	if err != nil {
		return
	}
	message := sarama.ProducerMessage{
		Topic:     p.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(json),
	}
	_, _, err = p.prod.SendMessage(&message)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message")
	}
}
