package schema

import (
	"entgo.io/ent"
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
		field.String("title").NotEmpty(),
		field.Bool("is_preposition"),
		field.Int("occurence").Optional(),
	}
}

// Edges of the WordNode.
func (WordNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("words", Word.Type),
	}
}
