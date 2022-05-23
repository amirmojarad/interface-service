package crud_test

import (
	"interface_project/ent"
	"interface_project/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieCrud(t *testing.T) {
	testClient := test.GetTestClientAndContext(t)
	defer testClient.CallCancelAndClose()
	addMovieTest(*testClient)
}

func addMovieTest(tc test.TestClient) {
	movieCreateSchemas := make([]*ent.MovieCreate, 1)
	movieCreateSchemas = append(movieCreateSchemas, tc.Crud.Client.Movie.Create().
	SetGenres("Genres").
	SetImDbRating("ImDbRating").
	SetImageURL("Image").
	SetPlot("Plot").
	SetStars("Stars").
	SetRuntimeStr("RuntimeStr").
	SetTitle("Title").
	SetYear("Description").
	SetMetacriticRating("MetacriticRating"))
	movies, err := tc.Crud.AddMovies(movieCreateSchemas)
	if err != nil {
		tc.T.Fatal(err)
	}
	assert.Equal(&tc.T, 1, len(movies))
}
