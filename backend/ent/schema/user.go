package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User schema definition.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty(),
		field.String("email").Unique(),
		field.String("password_hash").NotEmpty(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type),
	}
}
