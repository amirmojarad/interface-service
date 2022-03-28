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
		field.String("year"),
		field.String("image_url"),
		field.String("runtimeStr"),
		field.String("genres"),
		field.String("imDbRating"),
		field.String("plot"),
		field.String("stars"),
		field.String("metacriticRating"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}
