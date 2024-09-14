package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post schema definition.
type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").NotEmpty(),
		field.Float("location_x"),
		field.Float("location_y"),
		field.String("video_url").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type). // Ensure the inverse edge matches in User
						Ref("posts"). // Use plural for consistency
						Unique(),
		edge.From("event", Event.Type). // Ensure the inverse edge matches in Event
						Ref("posts"),
		edge.To("images", PostImage.Type). // Ensure the edge to PostImage is correct
							Unique(),
	}
}
