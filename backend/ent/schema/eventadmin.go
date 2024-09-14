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
		edge.From("event", Event.Type).
			Ref("event_admins").
			Unique(),
		edge.From("user", User.Type).
			Ref("event_admins").
			Unique(),
	}
}
