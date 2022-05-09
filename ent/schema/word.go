package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Word holds the schema definition for the Word entity.
type Word struct {
	ent.Schema
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Annotations(entproto.Field(2)),
		field.String("meaning").Optional().Annotations(entproto.Field(3)),
		field.String("sentence").Optional().Annotations(entproto.Field(4)),
		field.String("duration").Optional().Annotations(entproto.Field(5)),
		field.Time("start").Optional().Annotations(entproto.Field(6)),
		field.Time("end").Optional().Annotations(entproto.Field(7)),
	}
}

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("movie", Movie.Type).Unique().Annotations(entproto.Field(8)),
		edge.From("user", User.Type).Ref("favorite_words").Unique().Annotations(entproto.Field(9)),
	}
}

func (Word) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
