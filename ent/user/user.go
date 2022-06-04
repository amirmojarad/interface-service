// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldCreatedDate holds the string denoting the created_date field in the database.
	FieldCreatedDate = "created_date"
	// FieldUpdatedDate holds the string denoting the updated_date field in the database.
	FieldUpdatedDate = "updated_date"
	// FieldIsAdmin holds the string denoting the is_admin field in the database.
	FieldIsAdmin = "is_admin"
	// EdgeFavoriteMovies holds the string denoting the favorite_movies edge name in mutations.
	EdgeFavoriteMovies = "favorite_movies"
	// EdgeSearchedKeywords holds the string denoting the searched_keywords edge name in mutations.
	EdgeSearchedKeywords = "searched_keywords"
	// EdgeFavoriteWords holds the string denoting the favorite_words edge name in mutations.
	EdgeFavoriteWords = "favorite_words"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FavoriteMoviesTable is the table that holds the favorite_movies relation/edge. The primary key declared below.
	FavoriteMoviesTable = "user_favorite_movies"
	// FavoriteMoviesInverseTable is the table name for the Movie entity.
	// It exists in this package in order to avoid circular dependency with the "movie" package.
	FavoriteMoviesInverseTable = "movies"
	// SearchedKeywordsTable is the table that holds the searched_keywords relation/edge.
	SearchedKeywordsTable = "search_keywords"
	// SearchedKeywordsInverseTable is the table name for the SearchKeyword entity.
	// It exists in this package in order to avoid circular dependency with the "searchkeyword" package.
	SearchedKeywordsInverseTable = "search_keywords"
	// SearchedKeywordsColumn is the table column denoting the searched_keywords relation/edge.
	SearchedKeywordsColumn = "user_searched_keywords"
	// FavoriteWordsTable is the table that holds the favorite_words relation/edge.
	FavoriteWordsTable = "words"
	// FavoriteWordsInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	FavoriteWordsInverseTable = "words"
	// FavoriteWordsColumn is the table column denoting the favorite_words relation/edge.
	FavoriteWordsColumn = "user_favorite_words"
	// FilesTable is the table that holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "user_files"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldPassword,
	FieldFullName,
	FieldCreatedDate,
	FieldUpdatedDate,
	FieldIsAdmin,
}

var (
	// FavoriteMoviesPrimaryKey and FavoriteMoviesColumn2 are the table columns denoting the
	// primary key for the favorite_movies relation (M2M).
	FavoriteMoviesPrimaryKey = []string{"user_id", "movie_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultFullName holds the default value on creation for the "full_name" field.
	DefaultFullName string
	// DefaultCreatedDate holds the default value on creation for the "created_date" field.
	DefaultCreatedDate time.Time
	// DefaultUpdatedDate holds the default value on creation for the "updated_date" field.
	DefaultUpdatedDate time.Time
	// DefaultIsAdmin holds the default value on creation for the "is_admin" field.
	DefaultIsAdmin bool
)
