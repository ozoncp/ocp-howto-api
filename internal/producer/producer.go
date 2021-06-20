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
	sender  sarama.SyncProducer
	brokers []string
	topic   string
	events  chan Event
	close   chan struct{}
	done    chan struct{}
}

func New(brokers []string, topic string, capacity int) Producer {

	prod := producer{
		sender:  nil,
		brokers: brokers,
		topic:   topic,
		events:  make(chan Event, capacity),
		close:   make(chan struct{}),
		done:    make(chan struct{}),
	}

	_, err := prod.getSender()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to create producer sender")
	}

	return &prod
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
			p.done <- struct{}{}
			return
		}
	}
}

func (p *producer) flush() {
	for event := range p.events {
		p.send(event)
	}
	if p.sender != nil {
		p.sender.Close()
	}
}

func (p *producer) getSender() (sarama.SyncProducer, error) {
	if p.sender != nil {
		return p.sender, nil
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	sender, err := sarama.NewSyncProducer(p.brokers, config)
	if err != nil {
		return nil, err
	}

	p.sender = sender
	return sender, nil
}

func (p *producer) send(event Event) {

	sender, err := p.getSender()
	if err != nil {
		log.Error().Err(err).Msg("Failed to send event")
		return
	}

	json, err := json.Marshal(event)
	if err != nil {
		return
	}
	message := sarama.ProducerMessage{
		Topic:     p.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(json),
	}
	_, _, err = sender.SendMessage(&message)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send event")
	}
}
