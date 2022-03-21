package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Word holds the schema definition for the Word entity.
type Word struct {
	ent.Schema
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("meaning").Optional(),
		field.String("sentence").Optional(),
		field.String("duration").Optional(),
	}
}

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return nil
}
