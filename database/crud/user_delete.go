package crud

import "interface_project/ent"

func (crud Crud) DeleteUserByEmail(email string) (*ent.User, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		err := crud.Client.User.DeleteOne(u).Exec(*crud.Ctx)
		return u, err
	}
}
