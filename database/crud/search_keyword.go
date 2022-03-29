package crud

import (
	"interface_project/ent"
)

func (crud Crud) AddSearchKeywords(keywords []string, email string) ([]*ent.SearchKeyword, error) {
	bulk := make([]*ent.SearchKeywordCreate, len(keywords))
	u, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	for i, keyword := range keywords {
		bulk[i] = crud.Client.SearchKeyword.Create().SetTitle(keyword).SetUserID(u.ID)
	}
	if keywords, err := crud.Client.SearchKeyword.CreateBulk(bulk...).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return keywords, nil
	}
}
