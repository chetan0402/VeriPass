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
		field.String("id").StorageKey("sch_id").Unique(),
		field.String("name"),
		field.String("room"),
		field.String("hostel"),
		field.String("phone"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
