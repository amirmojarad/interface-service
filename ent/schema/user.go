package schema

import (
	"entgo.io/ent/dialect/entsql"
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().Annotations(
			entproto.Field(2),
		).NotEmpty(),
		field.String("email").Unique().Annotations(
			entproto.Field(3),
		).NotEmpty(),
		field.String("password").Annotations(
			entproto.Field(4),
		).NotEmpty(),
		field.String("full_name").Default("").Annotations(
			entproto.Field(5),
		).Optional(),
		field.Time("created_date").Annotations(
			entproto.Field(6),
		).Default(time.Now()),
		field.Time("updated_date").Annotations(
			entproto.Field(7),
		).Default(time.Now()),
		field.Bool("is_admin").Annotations(
			entproto.Field(8),
		).Default(false),
		field.String("image_url").Optional().Annotations(entproto.Field(9)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("favorite_movies", Movie.Type).Annotations(entproto.Field(10), entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("searched_keywords", SearchKeyword.Type).Annotations(entproto.Field(11), entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("favorite_words", Word.Type).Annotations(entproto.Field(12), entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("files", FileEntity.Type).Annotations(entproto.Field(13), entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("collections", Collection.Type).Annotations(entproto.Field(14), entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
