package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("year").NotEmpty(),
		field.String("image_url").NotEmpty(),
		field.String("runtimeStr").NotEmpty(),
		field.String("genres").NotEmpty(),
		field.String("imDbRating").NotEmpty(),
		field.String("plot").NotEmpty(),
		field.String("stars").NotEmpty(),
		field.String("metacriticRating").NotEmpty(),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}
