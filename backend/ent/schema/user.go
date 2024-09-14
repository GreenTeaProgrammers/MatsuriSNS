package schema

import (
	"time"

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
		field.Int("user_id").Unique(),
		field.String("username").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("hashed_password").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),              // Matches the edge in Post
		edge.To("events", Event.Type),            // Matches the edge in Event
		edge.To("event_admins", EventAdmin.Type), // Matches the edge in EventAdmin
	}
}
