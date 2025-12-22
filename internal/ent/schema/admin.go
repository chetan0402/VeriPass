// Package schema outlines the ORM configuration for Ent-Go
package schema

import (
	"crypto/rand"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Admin defines the schema definition for the Admin entity.
// This entity represents users such as guard/caretaker/wardens.
type Admin struct {
	ent.Schema
}

// Fields defines the columns of the table for Admin.
//
// "id" is a random 26 character string used to uniquely identify an admin.
// This is is subject to change as there is no specific reason for a
// random 26 string to be used for identification purpose.
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

// Edges returns nil as there are no foreign key constraints on the entity Admin
func (Admin) Edges() []ent.Edge {
	return nil
}
