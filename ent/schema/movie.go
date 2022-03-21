package schema

import "entgo.io/ent"

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return nil
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}
