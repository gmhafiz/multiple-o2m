package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type MultipleMany struct {
	ent.Schema
}

func (MultipleMany) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id"),
		field.Int("user_1"),
		field.Int("user_2"),
	}
}

func (MultipleMany) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user1", User.Type).Ref("multiple_many_1").Unique().Required().Field("user_1"),
		edge.From("user2", User.Type).Ref("multiple_many_2").Unique().Required().Field("user_2"),
	}
}
