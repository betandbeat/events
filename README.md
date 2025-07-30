# events

A strongly-typed, idiomatic Go event library for BetandBeat microservices. This library provides a simple, consistent way to define, publish, and consume events using Google Eventarc and CloudEvents.

---

## Quick Start

### 1. Define Your Event

Create a struct for your event in the `event/` package. Implement the `EventType() string` method, returning the event name (see naming convention below).

Then, have all event attributes public and of primitive types (JSON transport).

```go
package event

type UserSignedUp struct {
    ID string `json:"id"`
    At string `json:"at"`
}

func (UserSignedUp) EventType() string { return "user.signed_up" }
```

### 2. Publish an Event

```go
import (
    "context"
    "github.com/betandbeat/events"
    "github.com/betandbeat/events/event"
)

func main() {
    ctx := context.Background()
    publisher, closeFn, err := events.NewEventarc(
        ctx,
        "projects/your-project/locations/your-location/channels/your-channel", // the eventarc advanced bus resource name
        "betandbeat.iam", // the source, which service is publishing events
    )
    if err != nil {
        panic(err)
    }
    defer closeFn()

    evt := event.UserSignedUp{ID: "123", At: "2025-07-30T12:00:00Z"}
    if err := publisher.Publish(ctx, evt); err != nil {
        panic(err)
    }
}
```

---

## Adding a New Event

1. **Create a new struct** in the `event/` package.
2. **Implement the `EventType() string` method** to return the event name (see below).
3. **Follow the naming convention** for the event type string.

Example:
```go
package event

type OrderShipped struct {
    OrderID string `json:"order_id"`
    ShippedAt string `json:"shipped_at"`
}

func (OrderShipped) EventType() string { return "order.shipped" }
```

---

## Event Naming Convention

To ensure consistency across our systems (analytics, webhooks, event streams), we follow a strict event naming convention.

### Pattern
```
<entity>.<action>[_<qualifier>]
```
- **entity** → Singular noun representing the subject of the event
- **action** → Past-tense verb describing what happened
- **qualifier** (optional) → Adds detail if needed for disambiguation

### Rules
1. Use singular entity names
   - ✅ user.signed_up
   - ❌ users.signed_up
2. Use lowercase snake_case
   - ✅ user.signed_up
   - ❌ user.signedUp
   - ❌ user.signed-up
3. Use past tense for completed events
   - ✅ order.shipped
   - ❌ order.ship
   - ❌ order.shipping
4. Optional qualifiers for more detail
   - user.password_reset_requested
   - user.password_reset_completed
   - cart.item_added
   - cart.item_removed
5. Namespace complex domains with sub-entities if needed
   - subscription.payment_failed
   - invoice.payment_succeeded

### Examples
- user.signed_up               # A new user registered
- user.email_verified          # A user successfully verified their email
- order.created                # A new order was placed
- order.shipped                # An order was shipped
- cart.item_added              # A product was added to the user’s cart
- subscription.payment_failed  # A recurring payment attempt did not succeed

### Benefits
- Consistent across services and languages
- Human-readable and machine-parseable
- Future-proof for analytics, event streaming, and external integrations

---

## FAQ

**Q: Do I need to set event metadata (ID, source, content type, etc.)?**
- No. The publisher handles all CloudEvent metadata. You only need to provide your strongly-typed struct and implement `EventType()`.

**Q: Where do I put new event types?**
- In the `event/` package. Group related events in files as needed (e.g., `iam_events.go`, `order_events.go`).

**Q: How do I test my event?**
- Write unit tests in the `event/` package or integration tests using the publisher interface.

**Q: Why is this repository public?**
- Because there's no trade secret in here :)

---

## Contributing
- Follow the event naming convention.
- Keep event payloads minimal and explicit.
- Document new event types with comments.
- PRs welcome!