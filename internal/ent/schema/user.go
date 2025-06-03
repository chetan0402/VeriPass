package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("name"),
		field.String("hostel"),
		field.Bool("can_add_pass"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
