// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"interface_project/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// CreatedDate holds the value of the "created_date" field.
	CreatedDate time.Time `json:"created_date,omitempty"`
	// UpdatedDate holds the value of the "updated_date" field.
	UpdatedDate time.Time `json:"updated_date,omitempty"`
	// IsAdmin holds the value of the "is_admin" field.
	IsAdmin bool `json:"is_admin,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// FavoriteMovies holds the value of the favorite_movies edge.
	FavoriteMovies []*Movie `json:"favorite_movies,omitempty"`
	// SearchedKeywords holds the value of the searched_keywords edge.
	SearchedKeywords []*SearchKeyword `json:"searched_keywords,omitempty"`
	// FavoriteWords holds the value of the favorite_words edge.
	FavoriteWords []*Word `json:"favorite_words,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// FavoriteMoviesOrErr returns the FavoriteMovies value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FavoriteMoviesOrErr() ([]*Movie, error) {
	if e.loadedTypes[0] {
		return e.FavoriteMovies, nil
	}
	return nil, &NotLoadedError{edge: "favorite_movies"}
}

// SearchedKeywordsOrErr returns the SearchedKeywords value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SearchedKeywordsOrErr() ([]*SearchKeyword, error) {
	if e.loadedTypes[1] {
		return e.SearchedKeywords, nil
	}
	return nil, &NotLoadedError{edge: "searched_keywords"}
}

// FavoriteWordsOrErr returns the FavoriteWords value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FavoriteWordsOrErr() ([]*Word, error) {
	if e.loadedTypes[2] {
		return e.FavoriteWords, nil
	}
	return nil, &NotLoadedError{edge: "favorite_words"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[3] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsAdmin:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldEmail, user.FieldPassword, user.FieldFullName:
			values[i] = new(sql.NullString)
		case user.FieldCreatedDate, user.FieldUpdatedDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				u.FullName = value.String
			}
		case user.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_date", values[i])
			} else if value.Valid {
				u.CreatedDate = value.Time
			}
		case user.FieldUpdatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_date", values[i])
			} else if value.Valid {
				u.UpdatedDate = value.Time
			}
		case user.FieldIsAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_admin", values[i])
			} else if value.Valid {
				u.IsAdmin = value.Bool
			}
		}
	}
	return nil
}

// QueryFavoriteMovies queries the "favorite_movies" edge of the User entity.
func (u *User) QueryFavoriteMovies() *MovieQuery {
	return (&UserClient{config: u.config}).QueryFavoriteMovies(u)
}

// QuerySearchedKeywords queries the "searched_keywords" edge of the User entity.
func (u *User) QuerySearchedKeywords() *SearchKeywordQuery {
	return (&UserClient{config: u.config}).QuerySearchedKeywords(u)
}

// QueryFavoriteWords queries the "favorite_words" edge of the User entity.
func (u *User) QueryFavoriteWords() *WordQuery {
	return (&UserClient{config: u.config}).QueryFavoriteWords(u)
}

// QueryFiles queries the "files" edge of the User entity.
func (u *User) QueryFiles() *FileQuery {
	return (&UserClient{config: u.config}).QueryFiles(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", full_name=")
	builder.WriteString(u.FullName)
	builder.WriteString(", created_date=")
	builder.WriteString(u.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", updated_date=")
	builder.WriteString(u.UpdatedDate.Format(time.ANSIC))
	builder.WriteString(", is_admin=")
	builder.WriteString(fmt.Sprintf("%v", u.IsAdmin))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
