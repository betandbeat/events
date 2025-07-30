package events

import cloudevents "github.com/cloudevents/sdk-go/v2"

// CloudEventHeaders defines the headers for a CloudEvent in the context of Huma (huma.rocks).
// As per https://cloud.google.com/eventarc/docs/cloudevents#http-request
// ce-id			Unique identifier for the event	1096434104173400
// ce-source		Identifies the source of the event	//pubsub.googleapis.com/projects/my-project/topics/my-topic
// ce-specversion	The CloudEvents specification version used for this event	1.0
// ce-type			The type of event data	google.cloud.pubsub.topic.v1.messagePublished
// ce-time			Event generation time, in RFC 3339 format (optional)	2020-12-20T13:37:33.647Z
type CloudEventHeaders struct {
	ID          string `header:"ce-id"`
	Source      string `header:"ce-source"`
	SpecVersion string `header:"ce-specversion"`
	Type        string `header:"ce-type"`
	Time        string `header:"ce-time"` // Optional
}

// ToCloudEvent converts the CloudEventHeaders to a CloudEvent with the provided data.
func (ceh CloudEventHeaders) ToCloudEvent(data interface{}) cloudevents.Event {
	ev := cloudevents.NewEvent()
	ev.SetID(ceh.ID)
	ev.SetSource(ceh.Source)
	ev.SetSpecVersion(ceh.SpecVersion)
	ev.SetType(ceh.Type)
	if ceh.Time != "" {
		parsedTime, err := cloudevents.ParseTimestamp(ceh.Time)
		if err == nil && parsedTime != nil {
			ev.SetTime(parsedTime.Time)
		}
	}
	ev.SetData(cloudevents.ApplicationJSON, data)
	return ev
}
