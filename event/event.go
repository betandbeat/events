package event

// Event defines the interface for an event with a type.
type Event interface {
	EventType() string
}

func ListAll() []Event {
	var IAM_EVENTS = []Event{
		&UserSignedIn{},
		&UserSignedUp{},
	}

	var TEST_EVENTS = []Event{
		&SomethingHappened{},
	}

	var ALL_EVENTS []Event
	ALL_EVENTS = append(ALL_EVENTS, IAM_EVENTS...)
	ALL_EVENTS = append(ALL_EVENTS, TEST_EVENTS...)
	return ALL_EVENTS
}
