package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Report schema definition.
type Report struct {
    ent.Schema
}

func (Report) Fields() []ent.Field {
    return []ent.Field{
        field.Text("reason").NotEmpty(),
    }
}

func (Report) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("post", Post.Type).
            Ref("reports").
            Unique(),
        edge.From("reported_by", User.Type).
            Ref("reports").
            Unique(),
    }
}

