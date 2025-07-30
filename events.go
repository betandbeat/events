package events

import (
	"context"
)

// Publisher defines the interface for publishing events.
type Publisher interface {
	Publish(ctx context.Context, event Event) error
}

type Event interface {
	EventType() string
}
