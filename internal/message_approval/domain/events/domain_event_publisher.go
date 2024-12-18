package events

type DomainEventPublisher interface {
	Publish(event interface{})
}
