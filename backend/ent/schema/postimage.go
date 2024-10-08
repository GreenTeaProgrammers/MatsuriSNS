package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PostImage スキーマの定義
type PostImage struct {
	ent.Schema
}

func (PostImage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("post_id"),
		field.String("image_url").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

func (PostImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("post", Post.Type).
			Unique().
			Required().
			Field("post_id"),
	}
}
