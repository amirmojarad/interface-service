package crud

import "interface_project/ent"

func (crud Crud) AddWordsToUser(wordCreateBulk []*ent.WordCreate) ([]*ent.Word, error) {
	return crud.Client.Word.CreateBulk(wordCreateBulk...).Save(*crud.Ctx)
}

func (crud Crud) GetUserWords(user *ent.User) ([]*ent.Word, error) {
	return nil, nil
}
