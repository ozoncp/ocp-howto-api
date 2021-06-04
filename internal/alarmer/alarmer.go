package alarmer

type Alarmer interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}
