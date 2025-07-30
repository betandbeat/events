package event

type UserSignedUp struct {
	ID string `json:"id"`
	At string `json:"at"`
}

func (UserSignedUp) EventType() string { return "user.signedup" }

type UserSignedIn struct {
	ID string `json:"id"`
	At string `json:"at"`
}

func (UserSignedIn) EventType() string { return "user.signedin" }
