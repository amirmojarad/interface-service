package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Annotations(entproto.Field(2)),
		field.Time("created_date").Default(time.Now()).Annotations(entproto.Field(3)),
		field.Time("updated_date").Default(time.Now()).Annotations(entproto.Field(4)),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("collections").Annotations(entproto.Field(5)),
		edge.To("collection_words", Word.Type).Annotations(entproto.Field(6)),
	}
}

func (Collection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
