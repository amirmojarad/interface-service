package crud

import (
	"interface_project/ent"
	"interface_project/ent/user"
	"interface_project/usecases/generators"
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
	if u, err := crud.Client.User.Query().Where(user.Email(email)).First(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

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

func (crud Crud) ChangeFullName(id int, fullName string) (*ent.User, error) {
	if u, err := crud.GetUserByID(id); err != nil {
		return nil, err
	} else {
		if u, err = u.Update().SetFullName(fullName).Save(*crud.Ctx); err != nil {
			return nil, err
		} else {
			return u, nil
		}
	}
}

func (crud Crud) DeleteUserByEmail(email string) (*ent.User, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		err := crud.Client.User.DeleteOne(u).Exec(*crud.Ctx)
		return u, err
	}
}
