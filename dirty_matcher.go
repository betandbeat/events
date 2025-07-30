package events

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ReflectionEventMatcher struct {
	handlers map[string][]handlerEntry
}

type handlerEntry struct {
	handler reflect.Value
	typ     reflect.Type
}

func NewReflectionEventMatcher() *ReflectionEventMatcher {
	return &ReflectionEventMatcher{handlers: make(map[string][]handlerEntry)}
}

// Add registers a handler for a specific event type.
// handler must be a function: func(T) (any, error), where T is the event struct.
func (m *ReflectionEventMatcher) Add(eventType string, handler any) {
	fn := reflect.ValueOf(handler)
	typ := fn.Type().In(0)
	m.handlers[eventType] = append(m.handlers[eventType], handlerEntry{handler: fn, typ: typ})
}

// Handle dispatches the event to all handlers registered for the event type.
// Returns a slice of results and the first error encountered (if any).
func (m *ReflectionEventMatcher) Handle(eventType string, data []byte) ([]any, error) {
	entries, ok := m.handlers[eventType]
	if !ok || len(entries) == 0 {
		return nil, fmt.Errorf("no handler for event type: %s", eventType)
	}
	var results []any
	for _, entry := range entries {
		arg := reflect.New(entry.typ).Interface()
		if err := json.Unmarshal(data, arg); err != nil {
			return results, err
		}
		callResults := entry.handler.Call([]reflect.Value{reflect.ValueOf(arg).Elem()})
		if err, ok := callResults[1].Interface().(error); ok && err != nil {
			return results, err
		}
		results = append(results, callResults[0].Interface())
	}
	return results, nil
}
