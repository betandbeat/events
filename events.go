package events

import (
	"context"

	"github.com/betandbeat/events/event"
)

// Publisher defines the interface for publishing events.
type Publisher interface {
	Publish(ctx context.Context, event event.Event) error
}
