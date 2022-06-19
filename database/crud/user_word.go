package crud

import (
	"entgo.io/ent/dialect/sql"
	"interface_project/ent"
	"interface_project/ent/word"
)

func (crud Crud) AddFavoriteWordsToUser(words []string, email string) ([]*ent.Word, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	fetchedWordsIDs, err := crud.Client.Word.Query().Where(word.TitleIn(words...)).IDs(*crud.Ctx)
	if err != nil {
		return nil, err
	}
	user, err = user.Update().AddFavoriteWordIDs(fetchedWordsIDs...).Save(*crud.Ctx)
	if err != nil {
		return nil, err
	}
	return crud.Client.Word.Query().Where(word.IDIn(fetchedWordsIDs...)).All(*crud.Ctx)
}

func (crud Crud) GetUserFavoriteWords(email string) ([]*ent.Word, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user.QueryFavoriteWords().Where(func(s *sql.Selector) {
		s.Distinct().OnP(s.Select("title").P()).Select("title").Dialect()
	}).All(*crud.Ctx)
}

func (crud Crud) GetUserWords(email string, titles []string, fileID int) ([]*ent.Word, error) {
	file, err := crud.GetUserFileByID(fileID, email)
	if err != nil {
		return nil, err
	}
	return file.QueryWords().Where(word.TitleIn(titles...)).All(*crud.Ctx)
}

func (crud Crud) GetWordSentences(email, title string) ([]*ent.Word, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user.QueryWords().Where(word.TitleEQ(title)).All(*crud.Ctx)
}
