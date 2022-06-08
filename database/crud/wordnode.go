package crud

import (
	"interface_project/ent"
	"interface_project/ent/fileentity"
	"interface_project/ent/predicate"
	"interface_project/ent/wordnode"
)

func (crud Crud) GetAllWordNodes() ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().All(*crud.Ctx)
}

func (crud Crud) UpdateWordnode(wordnode *ent.WordNode, word *ent.Word) (*ent.WordNode, error) {
	updatedWordnode, err := crud.Client.WordNode.UpdateOne(wordnode).AddOccurence(1).Save(*crud.Ctx)
	return updatedWordnode, err
}

func (crud Crud) CreateWordnode(word *ent.Word, fileID int, isPreposition bool) (*ent.WordNode, error) {
	return crud.Client.WordNode.Create().
		SetFileID(fileID).
		SetTitle(word.Title).
		SetIsPreposition(isPreposition).
		SetOccurence(1).
		AddWords(word).
		Save(*crud.Ctx)
}

func (crud Crud) AddWordNode(wordNodeSchema *ent.WordNode, fileID int) (*ent.WordNode, error) {
	if createdWordNode, err := crud.Client.WordNode.Create().
		SetTitle(wordNodeSchema.Title).
		SetOccurence(wordNodeSchema.Occurence).
		SetIsPreposition(wordNodeSchema.IsPreposition).SetFileID(fileID).Save(*crud.Ctx); err != nil {
		return nil, err
	} else {
		return createdWordNode, nil
	}
}

func (crud Crud) CreateNewWordNode(wordNodeSchema *ent.WordNode, fileID int) *ent.WordNodeCreate {
	return crud.Client.WordNode.Create().
		SetTitle(wordNodeSchema.Title).
		SetOccurence(wordNodeSchema.Occurence).
		SetIsPreposition(wordNodeSchema.IsPreposition).SetFileID(fileID)
}

func (crud Crud) CreateWordNodes(wordNodes []*ent.WordNodeCreate) ([]*ent.WordNode, error) {
	createdWordNodes, err := crud.Client.WordNode.CreateBulk(wordNodes...).Save(*crud.Ctx)
	if err != nil {
		return nil, err
	} else {
		return createdWordNodes, nil
	}
}

func (crud Crud) GetWordNodeByTitle(title string) (*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(wordnode.TitleEQ(title)).First(*crud.Ctx)
}

func (crud Crud) GetFileWordNodes(fileID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(wordnode.HasFileWith(fileentity.IDEQ(fileID))).All(*crud.Ctx)
}

func (crud Crud) GetWordsFromWordNode(wordnodeID int) ([]*ent.Word, error) {
	return crud.Client.WordNode.Query().Where(wordnode.IDEQ(wordnodeID)).QueryWords().All(*crud.Ctx)
}

func (crud Crud) SortByID(fileID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(predicate.WordNode(fileentity.ID(fileID))).Order(ent.Asc(wordnode.FieldID)).All(*crud.Ctx)
}

func (crud Crud) SortByPreposition(fileID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(wordnode.HasFileWith(fileentity.ID(fileID))).Order(ent.OrderFunc(wordnode.IsPreposition(true))).All(*crud.Ctx)
}

func (crud Crud) SortByOccurence(fileID int) ([]*ent.WordNode, error) {
	return crud.Client.WordNode.Query().Where(wordnode.HasFileWith(fileentity.ID(fileID))).Order(ent.Asc(wordnode.FieldOccurence)).All(*crud.Ctx)
}
