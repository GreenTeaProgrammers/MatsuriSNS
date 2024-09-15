package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event スキーマの定義
type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Text("description").Optional(),
		field.String("map_url").NotEmpty(),
		field.String("qr_code_url").Optional(),
		field.Time("start_time"),
		field.Time("end_time"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Int("creator_id"),
	}
}

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creator", User.Type).
			Unique().
			Required().
			Field("creator_id"),
		edge.From("event_admins", EventAdmin.Type).
			Ref("event"),
		edge.From("posts", Post.Type).
			Ref("event"),
	}
}
