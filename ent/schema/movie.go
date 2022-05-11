package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Annotations(entproto.Field(2)),
		field.String("year").Annotations(entproto.Field(3)),
		field.String("image_url").Annotations(entproto.Field(4)),
		field.String("runtimeStr").Annotations(entproto.Field(5)),
		field.String("genres").Annotations(entproto.Field(6)),
		field.String("imDbRating").Annotations(entproto.Field(7)),
		field.String("plot").Annotations(entproto.Field(8)),
		field.String("stars").Annotations(entproto.Field(9)),
		field.String("metacriticRating").Annotations(entproto.Field(10)),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("favorite_movies").Annotations(entproto.Field(11)),
		edge.From("word_nodes", WordNode.Type).Ref("movie_wordnode").Required().Unique().Annotations(entproto.Field(12)),
	}
}

func (Movie) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
