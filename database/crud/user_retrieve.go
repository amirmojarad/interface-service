package crud

import (
	"interface_project/ent"
	"interface_project/ent/user"
)
func (crud Crud) GetAllUsers() ([]*ent.User, error) {
	if users, err := crud.Client.User.Query().All(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (crud Crud) GetUserByID(id int) (*ent.User, error) {
	if u, err := crud.Client.User.Get(*crud.Ctx, id); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}



func (crud Crud) GetUserByEmail(email string) (*ent.User, error) {
	if u, err := crud.Client.User.Query().Where(user.EmailEQ(email)).First(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}