package events

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEvent struct {
	Name string
}

func (TestEvent) EventType() string {
	return "test.event"
}

func TestReflectionEventMatcher_Handle(t *testing.T) {
	matcher := NewReflectionEventMatcher()

	handled := false
	matcher.Add(TestEvent{}.EventType(), func(ev TestEvent) (any, error) {
		handled = true
		assert.Equal(t, "copilot", ev.Name, "expected Name to be 'copilot'")
		return "ok", nil
	})

	ev := TestEvent{Name: "copilot"}
	data, err := json.Marshal(ev)
	assert.NoError(t, err, "failed to marshal event")

	results, err := matcher.Handle(TestEvent{}.EventType(), data)
	assert.NoError(t, err, "Handle returned error")
	assert.Len(t, results, 1)
	assert.Equal(t, "ok", results[0])
	assert.True(t, handled, "handler was not called")
}

func TestReflectionEventMatcher_MultipleHandlers(t *testing.T) {
	matcher := NewReflectionEventMatcher()
	calls := []string{}

	matcher.Add(TestEvent{}.EventType(), func(ev TestEvent) (any, error) {
		calls = append(calls, "handler1")
		return "first", nil
	})
	matcher.Add(TestEvent{}.EventType(), func(ev TestEvent) (any, error) {
		calls = append(calls, "handler2")
		return "second", nil
	})

	ev := TestEvent{Name: "copilot"}
	data, err := json.Marshal(ev)
	assert.NoError(t, err, "failed to marshal event")

	results, err := matcher.Handle(TestEvent{}.EventType(), data)
	assert.NoError(t, err, "Handle returned error")
	assert.Len(t, results, 2)
	assert.Equal(t, "first", results[0])
	assert.Equal(t, "second", results[1])
	assert.Equal(t, []string{"handler1", "handler2"}, calls, "handlers not called in order")
}

func TestReflectionEventMatcher_NoHandler(t *testing.T) {
	matcher := NewReflectionEventMatcher()
	results, err := matcher.Handle("unknown.event", []byte(`{}`))
	assert.Error(t, err, "expected error for unknown event type")
	assert.Nil(t, results)
}

func TestReflectionEventMatcher_UnmarshalError(t *testing.T) {
	matcher := NewReflectionEventMatcher()
	matcher.Add(TestEvent{}.EventType(), func(ev TestEvent) (any, error) {
		return nil, nil
	})
	_, err := matcher.Handle(TestEvent{}.EventType(), []byte(`not-json`))
	assert.Error(t, err, "expected error for invalid JSON")
}
