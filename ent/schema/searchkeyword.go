package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SearchKeyword holds the schema definition for the SearchKeyword entity.
type SearchKeyword struct {
	ent.Schema
}

// Fields of the SearchKeyword.
func (SearchKeyword) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Annotations(entproto.Field(2)),
		field.Uint16("rate").Default(0).Annotations(entproto.Field(3)),
	}
}

// Edges of the SearchKeyword.
func (SearchKeyword) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("searched_keywords").Unique().Annotations(entproto.Field(4)),
	}
}

func (SearchKeyword) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
