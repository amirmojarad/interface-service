package word_node

import (
	"interface_project/ent"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	wordsMap := make(map[string][]*ent.Word, 10)
	for i := 0; i < 12; i++ {
		word := ent.Word{
			ID:       1,
			Title:    "",
			Meaning:  "asd",
			Sentence: "asd",
			Duration: "Asd",
			Start:    time.Now(),
			End:      time.Now(),
		}
		Add(&word, wordsMap)
	}
	assert.Equal(t, 1, len(wordsMap))
	assert.Equal(t, 12, len(wordsMap[""]))

}

func TestIsExist(t *testing.T) {

	wordsMap := make(map[string][]*ent.Word, 10)
	word := ent.Word{
		ID:       1,
		Title:    "title",
		Meaning:  "asd",
		Sentence: "asd",
		Duration: "Asd",
		Start:    time.Now(),
		End:      time.Now(),
	}
	Add(&word, wordsMap)
	assert.Equal(t, true, IsExist("title", wordsMap))

}

func TestGet(t *testing.T) {

	wordsMap := make(map[string][]*ent.Word, 10)
	word := ent.Word{
		ID:       1,
		Title:    "title",
		Meaning:  "asd",
		Sentence: "asd",
		Duration: "Asd",
		Start:    time.Now(),
		End:      time.Now(),
	}
	Add(&word, wordsMap)
	wordslist := Get("title", wordsMap)
	assert.Equal(t, 1, len(wordslist))

}
