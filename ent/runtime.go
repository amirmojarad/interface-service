// Code generated by entc, DO NOT EDIT.

package ent

import (
	"interface_project/ent/fileentity"
	"interface_project/ent/movie"
	"interface_project/ent/schema"
	"interface_project/ent/searchkeyword"
	"interface_project/ent/user"
	"interface_project/ent/word"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fileentityFields := schema.FileEntity{}.Fields()
	_ = fileentityFields
	// fileentityDescPath is the schema descriptor for path field.
	fileentityDescPath := fileentityFields[0].Descriptor()
	// fileentity.PathValidator is a validator for the "path" field. It is called by the builders before save.
	fileentity.PathValidator = fileentityDescPath.Validators[0].(func(string) error)
	// fileentityDescName is the schema descriptor for name field.
	fileentityDescName := fileentityFields[1].Descriptor()
	// fileentity.NameValidator is a validator for the "name" field. It is called by the builders before save.
	fileentity.NameValidator = fileentityDescName.Validators[0].(func(string) error)
	movieFields := schema.Movie{}.Fields()
	_ = movieFields
	// movieDescTitle is the schema descriptor for title field.
	movieDescTitle := movieFields[0].Descriptor()
	// movie.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	movie.TitleValidator = movieDescTitle.Validators[0].(func(string) error)
	searchkeywordFields := schema.SearchKeyword{}.Fields()
	_ = searchkeywordFields
	// searchkeywordDescTitle is the schema descriptor for title field.
	searchkeywordDescTitle := searchkeywordFields[0].Descriptor()
	// searchkeyword.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	searchkeyword.TitleValidator = searchkeywordDescTitle.Validators[0].(func(string) error)
	// searchkeywordDescRate is the schema descriptor for rate field.
	searchkeywordDescRate := searchkeywordFields[1].Descriptor()
	// searchkeyword.DefaultRate holds the default value on creation for the rate field.
	searchkeyword.DefaultRate = searchkeywordDescRate.Default.(uint16)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescFullName is the schema descriptor for full_name field.
	userDescFullName := userFields[3].Descriptor()
	// user.DefaultFullName holds the default value on creation for the full_name field.
	user.DefaultFullName = userDescFullName.Default.(string)
	// userDescCreatedDate is the schema descriptor for created_date field.
	userDescCreatedDate := userFields[4].Descriptor()
	// user.DefaultCreatedDate holds the default value on creation for the created_date field.
	user.DefaultCreatedDate = userDescCreatedDate.Default.(time.Time)
	// userDescUpdatedDate is the schema descriptor for updated_date field.
	userDescUpdatedDate := userFields[5].Descriptor()
	// user.DefaultUpdatedDate holds the default value on creation for the updated_date field.
	user.DefaultUpdatedDate = userDescUpdatedDate.Default.(time.Time)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[6].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
	wordFields := schema.Word{}.Fields()
	_ = wordFields
	// wordDescTitle is the schema descriptor for title field.
	wordDescTitle := wordFields[0].Descriptor()
	// word.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	word.TitleValidator = wordDescTitle.Validators[0].(func(string) error)
	// wordDescSentence is the schema descriptor for sentence field.
	wordDescSentence := wordFields[3].Descriptor()
	// word.SentenceValidator is a validator for the "sentence" field. It is called by the builders before save.
	word.SentenceValidator = wordDescSentence.Validators[0].(func(string) error)
	// wordDescDuration is the schema descriptor for duration field.
	wordDescDuration := wordFields[4].Descriptor()
	// word.DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	word.DurationValidator = wordDescDuration.Validators[0].(func(string) error)
	// wordDescStart is the schema descriptor for start field.
	wordDescStart := wordFields[5].Descriptor()
	// word.StartValidator is a validator for the "start" field. It is called by the builders before save.
	word.StartValidator = wordDescStart.Validators[0].(func(string) error)
	// wordDescEnd is the schema descriptor for end field.
	wordDescEnd := wordFields[6].Descriptor()
	// word.EndValidator is a validator for the "end" field. It is called by the builders before save.
	word.EndValidator = wordDescEnd.Validators[0].(func(string) error)
}
