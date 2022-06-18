package crud

import (
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

func (crud Crud) AddWordsToUser(wordCreateBulk []*ent.WordCreate) ([]*ent.Word, error) {
	return crud.Client.Word.CreateBulk(wordCreateBulk...).Save(*crud.Ctx)
}

func (crud Crud) GetUserWords(user *ent.User) ([]*ent.Word, error) {
	return nil, nil
}
