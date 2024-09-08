package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
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
    }
}

func (Post) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).
            Ref("posts").
            Unique(),
        edge.From("event", Event.Type).
            Ref("posts").
            Unique(),
    }
}

