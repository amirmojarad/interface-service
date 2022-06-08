package crud

import (
	"interface_project/ent"
	"interface_project/usecases/generators"
)
func (crud Crud) AddUsers(users []*ent.User) ([]*ent.User, error) {
	bulk := make([]*ent.UserCreate, len(users))
	for i, u := range users {
		var userCreate ent.UserCreate
		hashedPassword, _ := generators.HashPassword(u.Password)
		userCreate.SetEmail(u.Email).SetPassword(hashedPassword).SetUsername(u.Username).Save(*crud.Ctx)
		bulk[i] = &userCreate
	}
	if users, err := crud.Client.User.CreateBulk(bulk...).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}



func (crud Crud) AddUser(userSchema *ent.User) (*ent.User, error) {
	hashedPassword, _ := generators.HashPassword(userSchema.Password)
	if newUser, err := crud.Client.User.Create().SetEmail(userSchema.Email).SetPassword(hashedPassword).SetUsername(userSchema.Username).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return newUser, nil
	}
}