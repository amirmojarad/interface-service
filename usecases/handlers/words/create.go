package words

// import (
// 	"container/list"
// 	"interface_project/database/crud"
// 	"interface_project/ent"
// 	"interface_project/ent/fileentity"
// 	"interface_project/ent/wordnode"
// 	"log"

// 	// "interface_project/ent/wordnode"
// 	preposition "interface_project/usecases/handlers/prepositions"
// 	"interface_project/usecases/handlers/sentences"
// 	"os"
// )

// func CreateWordnodes(file *os.File, fileID int, crud *crud.Crud, user *ent.User) ([]*ent.WordNode, error) {
// 	wordCreates := sentences.GetSentences(crud.Client, file, user)
// 	words, err := crud.AddWordsToUser(wordCreates)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// wordnodesList := list.New()
// 	// for _, word := range words {
// 	// 	if wordnode, ok := adt.IsWordNodeExist(word.Title); ok {
// 	// 		adt.UpdateWordNode(wordnode, word.ID)
// 	// 	} else {
// 	// 		adt.AddWordNode(&WordNode{
// 	// 			Occurence:     1,
// 	// 			FileID:        fileID,
// 	// 			Title:         word.Title,
// 	// 			IsPreposition: preposition.IsPreposition(word.Title)})
// 	// 		log.Println(len(adt.Wordnodes))
// 	// 	}
// 	// }
// 	// for _, wordNode := range adt.Wordnodes {
// 	// 	log.Printf("%+v", wordNode)
// 	// }
// 	for _, word := range words {
// 		wordnode, err := crud.GetWordNodeByTitle(word.Title)
// 		if err != nil {
// 			crud.CreateWordnode(word, fileID, preposition.IsPreposition(word.Title))
// 		} else {

// 			crud.UpdateWordnode(wordnode, word)
// 		}
// 	}
// 	return crud.Client.WordNode.Query().Where(wordnode.HasFileWith(fileentity.IDEQ(fileID))).All(*crud.Ctx)
// }
