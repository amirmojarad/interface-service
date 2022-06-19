package crud

import (
	"interface_project/ent"
	"interface_project/ent/collection"
)

func (crud Crud) CreateCollection(email, title string, words []*ent.Word) (*ent.Collection, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return crud.Client.Collection.Create().AddUser(user).SetTitle(title).AddCollectionWords(words...).Save(*crud.Ctx)
}

func (crud Crud) GetCollections(email string) ([]*ent.Collection, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user.QueryCollections().WithCollectionWords().All(*crud.Ctx)
}

func (crud Crud) GetCollection(email string, collectionID int) (*ent.Collection, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user.QueryCollections().Where(collection.IDEQ(collectionID)).WithCollectionWords().First(*crud.Ctx)
}
