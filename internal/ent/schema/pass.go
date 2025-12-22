package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Pass defines the schema definition for the Pass entity.
// This entity represents an entry/exit of student in/out of the hostel.
type Pass struct {
	ent.Schema
}

// Fields defines the columns of the table for Pass.
// 
// "id" is UUIDv7 because
// 
// 1. Ordered: Since the primary key is physically stored as ordered on
// the disk, having UUIDv7 allows the writes to be sequential.
// 
// 2. Timestamp: Many queries in this application are around
// the create time of the entity, UUIDv7 encodes this information
// on the primary key itself allowing for faster queries
// 
// 3. Randomness: Compared to an autoincrement bigint primary key an UUID
// based provides better randomness.
func (Pass) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, err := uuid.NewV7()
			if err != nil {
				return uuid.Nil
			}
			return id
		}),
		field.String("user_id"),
		field.Enum("type").Values("unspecified", "class", "market", "home", "event").Default("unspecified"),
		field.Time("end_time").Optional(),
	}
}

// Edges defined the foreign key constraints of Pass entity.
// "user_id" is referes to User entity
func (Pass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("passes").Field("user_id").Unique().Required(),
	}
}
