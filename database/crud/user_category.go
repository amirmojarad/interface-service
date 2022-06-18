package crud

import (
	"interface_project/ent"
	"interface_project/ent/category"
)

func (crud Crud) CreateCategory(email, title string) (*ent.Category, error) {
	user, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, nil
	}
	return user.QueryCategories().Where(category.TitleEQ(title)).First(*crud.Ctx)
}
