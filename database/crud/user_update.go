package crud

import (
	"interface_project/ent"
	"time"
)

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

func (crud Crud) ChangePassword(userID int, password string) (*ent.User, error) {
	return nil, nil
}

func (crud Crud) UpdateUser(userID int, user *ent.User) (*ent.User, error) {
	updatedUser, err := crud.Client.User.
		UpdateOneID(userID).
		SetNillableImageURL(&user.ImageURL).
		SetNillableFullName(&user.FullName).
		SetUpdatedDate(time.Now()).Save(*crud.Ctx)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (crud Crud) AddImageUrlToUser(user *ent.User, imageURL string) (*ent.User, error) {
	if updatedUser, err := user.Update().SetImageURL(imageURL).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return updatedUser, err
	}
}
