package event

type SomethingHappened struct {
	What  string `json:"what"`
	When  string `json:"when"`
	Who   string `json:"who"`
	Where string `json:"where"`
	Why   string `json:"why"`
}

func (s SomethingHappened) EventType() string { return "something.happened" }
