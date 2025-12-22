package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User defines the schema definition for User entity
// This entity represents a student
type User struct {
	ent.Schema
}

// Fields defines the columns of table for User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name"),
		field.String("room"),
		field.String("hostel"),
		field.String("phone").Unique(),
	}
}

// Edges defines the constraint relation with Pass entity.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("passes", Pass.Type),
	}
}
