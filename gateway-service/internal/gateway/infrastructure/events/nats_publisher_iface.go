package events

type EventPublisher interface {
	// Request sends a request to NATS and waits for a reply.
	// It publishes the given data to the specified topic and returns the response.
	// 'topic' is the subject to which the request is sent.
	// 'data' is the payload that will be marshaled and sent.
	Request(topic string, data any) ([]byte, error)
}
