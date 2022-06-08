package crud

import (
	"interface_project/ent"
	"interface_project/ent/movie"
)


func (crud Crud) AddMoviesToUser(movieIDs []int, email string) ([]*ent.Movie, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		u, err = u.Update().AddFavoriteMovieIDs(movieIDs...).Save(*crud.Ctx)
		if err != nil {
			return nil, err
		}
		return u.QueryFavoriteMovies().AllX(*crud.Ctx), nil
	}
}

func (crud Crud) GetFavoriteMovies(email string) ([]*ent.Movie, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		if movies, err := u.QueryFavoriteMovies().All(*crud.Ctx); err != nil {
			return nil, err
		} else {
			return movies, nil
		}
	}
}


func (crud Crud) GetFavoriteMovie(email string, movieID int) (*ent.Movie, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		if movie, err := u.QueryFavoriteMovies().Where(movie.ID(movieID)).First(*crud.Ctx); err != nil {
			return nil, err
		} else {
			return movie, nil
		}
	}
}


func (crud Crud) DeleteMovieFromFavorites(email string, movieIDs []int) ([]*ent.Movie, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		if err = u.Update().RemoveFavoriteMovieIDs(movieIDs...).Exec(*crud.Ctx); err != nil {
			return nil, err
		} else {
			if movies, err := u.QueryFavoriteMovies().All(*crud.Ctx); err != nil {
				return nil, err
			} else {
				return movies, nil
			}
		}
	}
}
