package events

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
