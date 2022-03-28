package crud

import (
	"interface_project/ent"
	"interface_project/ent/movie"
)

func (crud *Crud) AddMovies(movies []*ent.MovieCreate) ([]*ent.Movie, error) {
	if movies, err := crud.Client.Movie.CreateBulk(movies...).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return movies, nil
	}
}

func (crud *Crud) GetAllMovies() ([]*ent.Movie, error) {
	if movies, err := crud.Client.Movie.Query().Order(ent.Asc(movie.FieldID)).All(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return movies, nil
	}
}

func (crud *Crud) GetMovie(movieID int) (*ent.Movie, error) {
	if fetchedMovie, err := crud.Client.Movie.Get(*crud.Ctx, movieID); err != nil {
		return nil, err
	} else {
		return fetchedMovie, nil
	}
}

func (crud *Crud) DeleteMovie(movieID int) (bool, error) {
	if err := crud.Client.Movie.DeleteOneID(movieID).Exec(*crud.Ctx); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
