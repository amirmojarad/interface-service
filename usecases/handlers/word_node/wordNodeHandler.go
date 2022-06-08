package word_node

import "interface_project/ent"

func Add(word *ent.Word, wordsMap map[string][]*ent.Word) {
	wordsMap[word.Title] = append(wordsMap[word.Title], word)
}

func Get(title string, wordsMap map[string][]*ent.Word) []*ent.Word {
	return wordsMap["title"]
}

func IsExist(title string, wordsMap map[string][]*ent.Word) bool {
	return wordsMap["title"] != nil
}
