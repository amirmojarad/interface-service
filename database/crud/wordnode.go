package crud

import (
	"interface_project/ent"
	"interface_project/ent/movie"
	"interface_project/ent/wordnode"
)

func (crud Crud) AddWordNode(wordNodeSchema *ent.WordNode) (*ent.WordNode, error) {
	if createdWordNode, err := crud.Client.WordNode.Create().
		SetTitle(wordNodeSchema.Title).
		SetOccurence(wordNodeSchema.Occurence).
		SetIsPreposition(wordNodeSchema.IsPreposition).
		Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return createdWordNode, nil
	}
}

func (crud Crud) CreateNewWordNode(wordNodeSchema *ent.WordNode) *ent.WordNodeCreate {
	return crud.Client.WordNode.Create().
		SetTitle(wordNodeSchema.Title).
		SetOccurence(wordNodeSchema.Occurence).
		SetIsPreposition(wordNodeSchema.IsPreposition)
}

func (crud Crud) CreateWordNodes(wordNodes []*ent.WordNodeCreate) ([]*ent.WordNode, error) {
	createdWordNodes, err := crud.Client.WordNode.CreateBulk(wordNodes...).Save(*crud.Ctx)
	if err != nil {
		return nil, err
	} else {
		return createdWordNodes, nil
	}
}

func (crud Crud) GetMovieWordnodes(movieID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(wordnode.HasMovieWordnodeWith(movie.ID(movieID))).All(*crud.Ctx)
}

func (crud Crud) GetWordsFromWordNode(movieID int) ([]*ent.Word, error) {
	return crud.Client.WordNode.Query().Where(wordnode.HasMovieWordnodeWith(movie.ID(movieID))).QueryWords().All(*crud.Ctx)
}

func (crud Crud) SortByID(movieID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Order(ent.Asc(wordnode.FieldID)).All(*crud.Ctx)
}

func (crud Crud) SortByPreposition(movieID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(wordnode.HasMovieWordnodeWith(movie.ID(movieID))).Order(ent.OrderFunc(wordnode.IsPreposition(true))).All(*crud.Ctx)
}

func (crud Crud) SortByOccurence(movieID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(wordnode.HasMovieWordnodeWith(movie.ID(movieID))).Order(ent.Asc(wordnode.FieldOccurence)).All(*crud.Ctx)
}
