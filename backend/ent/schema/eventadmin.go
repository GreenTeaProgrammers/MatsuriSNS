package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EventAdmin スキーマの定義
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
		edge.To("event", Event.Type).
			Unique().
			Required().
			Field("event_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}
