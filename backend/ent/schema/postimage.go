package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// PostImage schema definition.
type PostImage struct {
    ent.Schema
}

func (PostImage) Fields() []ent.Field {
    return []ent.Field{
        field.String("image_url").NotEmpty(),
    }
}

func (PostImage) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("post", Post.Type).
            Ref("images").
            Unique(),
    }
}

