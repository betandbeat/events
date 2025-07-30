package events

import (
	"context"
)

// Publisher defines the interface for publishing events.
type Publisher interface {
	Publish(ctx context.Context, event Event) error
}

// Event defines the interface for an event with a type.
type Event interface {
	EventType() string
}
