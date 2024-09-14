package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EventAdmin schema definition.
type EventAdmin struct {
	ent.Schema
}

func (EventAdmin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("event_id"),
		field.Int("user_id"),
		field.Time("created_at").Default(time.Now),
	}
}

func (EventAdmin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type). // Define the inverse edge to Event
						Ref("event_admins"). // Matches the edge defined in Event
						Unique(),
		edge.From("user", User.Type). // Define the inverse edge to User
						Ref("event_admins"). // Matches the edge defined in User
						Unique(),
	}
}
