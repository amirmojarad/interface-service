package crud

import (
	"interface_project/ent"
	"interface_project/ent/searchkeyword"
	"interface_project/ent/user"
)

func (crud Crud) AddSearchKeywords(keywords []string, email string) ([]*ent.SearchKeyword, error) {

	bulk := []*ent.SearchKeywordCreate{}
	u, err := crud.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	for _, keyword := range keywords {
		if fetchedKeyword := crud.Client.
			SearchKeyword.Query().
			Where(searchkeyword.HasUserWith(user.EmailEQ(email))).
			Where(searchkeyword.TitleContainsFold(keyword)).
			FirstX(*crud.Ctx); err != nil {
			return nil, err
		} else {
			if fetchedKeyword == nil {
				bulk = append(bulk, crud.Client.SearchKeyword.Create().SetTitle(keyword).SetUser(u))
			} else {
				crud.Client.SearchKeyword.UpdateOne(fetchedKeyword).AddRate(1).Save(*crud.Ctx)
			}
		}
	}
	if addedKeywords, err := crud.Client.SearchKeyword.CreateBulk(bulk...).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return addedKeywords, nil
	}
}
