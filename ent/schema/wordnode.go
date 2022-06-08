package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WordNode holds the schema definition for the WordNode entity.
type WordNode struct {
	ent.Schema
}

// Fields of the WordNode.
func (WordNode) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Unique().Annotations(entproto.Field(2)),
		field.Bool("is_preposition").Annotations(entproto.Field(3)),
		field.Int("occurence").Optional().Annotations(entproto.Field(4)),
	}
}

// Edges of the WordNode.
func (WordNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("words", Word.Type).Annotations(entproto.Field(5)),
		edge.From("file", FileEntity.Type).Unique().Required().Ref("wordnodes").Annotations(entproto.Field(6)),
	}
}

func (WordNode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
