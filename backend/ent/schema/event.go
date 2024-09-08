package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event schema definition.
type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Text("description").Optional(),
		field.String("map_url").NotEmpty(),
		field.String("qr_code_url").Optional(),
	}
}

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("created_by", User.Type).
			Ref("events").
			Unique(),
	}
}
