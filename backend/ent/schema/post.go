package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post スキーマの定義
type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("event_id"),
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
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("event", Event.Type).
			Unique().
			Required().
			Field("event_id"),
		edge.From("images", PostImage.Type).
			Ref("post"),
	}
}
