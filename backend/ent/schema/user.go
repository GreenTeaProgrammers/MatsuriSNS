package schema

import (
    "entgo.io/ent"
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

