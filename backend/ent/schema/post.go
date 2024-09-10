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
		field.String("comment").NotEmpty(),
		field.Float("location_x"),
		field.Float("location_y"),
		field.String("video_url").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("posts"),
		edge.From("event", Event.Type).
			Ref("posts"),
		edge.To("images", PostImage.Type),
	}
}
