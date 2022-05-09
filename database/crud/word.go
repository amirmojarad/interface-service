package crud

import (
	"interface_project/ent"
)

func (crud Crud) AddWordsToUser(words []*ent.WordCreate) ([]*ent.Word, error) {
	if createdWords, err := crud.Client.Word.CreateBulk(words...).Save(*crud.Ctx); err != nil {
		// return nil, errors.New("error while adding words to database")
		return nil, err
	} else {
		return createdWords, nil
	}
}
