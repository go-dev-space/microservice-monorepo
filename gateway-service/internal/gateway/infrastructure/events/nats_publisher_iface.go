package events

type EventPublisher interface {
	Request(topic string, data any) ([]byte, error)
}
