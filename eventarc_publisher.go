package events

import (
	"context"
	"encoding/json"
	"fmt"

	publishing "cloud.google.com/go/eventarc/publishing/apiv1"
	"cloud.google.com/go/eventarc/publishing/apiv1/publishingpb"
	"github.com/betandbeat/events/utils"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Eventarc is a concrete implementation using Google Eventarc.
type Eventarc struct {
	client     *publishing.PublisherClient
	messageBus string
	source     string
}

// NewEventarc creates a new Eventarc publisher and returns a close function.
func NewEventarc(ctx context.Context, messageBus string, source string) (Publisher, func() error, error) {
	client, err := publishing.NewPublisherClient(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create publisher client: %w", err)
	}
	p := &Eventarc{
		client:     client,
		messageBus: messageBus,
		source:     source,
	}
	closeFunc := func() error {
		return client.Close()
	}
	return p, closeFunc, nil
}

// Publish sends an Event to the configured message bus.
func (p *Eventarc) Publish(ctx context.Context, event Event) error {
	ev := cloudevents.NewEvent()
	ev.SetID(utils.NewEventID())
	ev.SetSource(p.source)
	ev.SetType(event.EventType())
	ev.SetDataContentType("application/json")
	ev.SetSpecVersion("1.0")
	ev.SetData(cloudevents.ApplicationJSON, event)

	stringifiedEvent, err := json.Marshal(ev)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	req := &publishingpb.PublishRequest{
		MessageBus: p.messageBus,
		Format: &publishingpb.PublishRequest_JsonMessage{
			JsonMessage: string(stringifiedEvent),
		},
	}
	_, err = p.client.Publish(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}
	return nil
}
