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
		field.String("meaning").Annotations(entproto.Field(3)),
		field.Bool("isPreposition").Annotations(entproto.Field(4)),
		field.String("sentence").NotEmpty().Annotations(entproto.Field(5)),
		field.String("duration").NotEmpty().Annotations(entproto.Field(6)),
		field.String("start").NotEmpty().Annotations(entproto.Field(7)),
		field.String("end").NotEmpty().Annotations(entproto.Field(8)),
	}
}

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("favorite_words").Unique().Annotations(entproto.Field(9)),
		edge.From("file", FileEntity.Type).Ref("words").Unique().Annotations(entproto.Field(10)),
		edge.From("collection", Collection.Type).Ref("collection_words").Annotations(entproto.Field(11)),
		edge.From("owner", User.Type).Ref("words").Unique().Annotations(entproto.Field(12)),
	}
}

func (Word) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
