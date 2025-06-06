package schema

import (
	"crypto/rand"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Admin struct {
	ent.Schema
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			return rand.Text()
		}).Unique().MaxLen(26).Immutable(),
		field.String("email").Unique(),
		field.String("name"),
		field.String("hostel"),
		field.Bool("can_add_pass"),
	}
}

func (Admin) Edges() []ent.Edge {
	return nil
}
