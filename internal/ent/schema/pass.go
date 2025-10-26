package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Pass struct {
	ent.Schema
}

func (Pass) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("user_id"),
		field.Enum("type").Values("unspecified", "class", "market", "home", "event").Default("unspecified"),
		field.Time("start_time"),
		field.Time("end_time").Optional(),
	}
}

func (Pass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("passes").Field("user_id").Unique().Required(),
	}
}
