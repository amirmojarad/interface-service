package crud

import "interface_project/ent"


func (crud Crud) AddSearchKeywordToUser(email string, keyword string) ([]*ent.SearchKeyword, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		keywords := make([]string, 1)
		keywords[0] = keyword
		if keywords, err := crud.AddSearchKeywords(keywords, u.Email); err != nil {
			return nil, err
		} else {
			return keywords, nil
		}
	}
}



func (crud Crud) GetUserSearchKeyword(email string) ([]*ent.SearchKeyword, error) {
	if u, err := crud.GetUserByEmail(email); err != nil {
		return nil, err
	} else {
		if keywords, err := u.QuerySearchedKeywords().All(*crud.Ctx); err != nil {
			return nil, err
		} else {
			return keywords, nil
		}
	}

}