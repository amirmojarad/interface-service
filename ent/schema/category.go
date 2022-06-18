package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Annotations(entproto.Field(2)),
		field.Time("created_date").Default(time.Now()).Annotations(entproto.Field(3)),
		field.Time("updated_date").Default(time.Now()).Annotations(entproto.Field(4)),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("categories").Annotations(entproto.Field(5)),
		edge.To("category_words", Word.Type).Annotations(entproto.Field(6)),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
