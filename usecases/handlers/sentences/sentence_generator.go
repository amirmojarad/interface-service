package sentences

import (
	"bufio"
	"interface_project/ent"
	preposition "interface_project/usecases/handlers/prepositions"
	"os"
	"strconv"
	"strings"
)

func GetSentences(client *ent.Client, file *os.File, user *ent.User, fileEntity *ent.FileEntity) []*ent.WordCreate {
	return makeWordsBulk(client, generateSentences(file), user, fileEntity)
}

func checkWordTitleIsValid(wordTitle string) bool {
	invalidChars := []string{
		"-",
		"a",
	}
	for _, ic := range invalidChars {
		if ic == wordTitle || strings.Contains(wordTitle, ic) {
			return false
		}
	}
	return true
}

func makeWordsBulk(client *ent.Client, sentences []*sentence, user *ent.User, file *ent.FileEntity) []*ent.WordCreate {
	wordBulk := []*ent.WordCreate{}
	for _, item := range sentences {
		times := strings.Split(item.timeRange, " ")
		for _, token := range item.tokens {
			if checkWordTitleIsValid(token) {
				if strings.Contains(token, "...") {
					token = strings.Replace(token, "...", "", -1)
				}
				wordBulk = append(wordBulk, client.Word.Create().
					SetTitle(token).
					SetMeaning("").
					SetFile(file).SetEnd(times[2]).SetStart(times[0]).
					SetIsPreposition(preposition.IsPreposition(token)).
					SetDuration(item.timeRange).SetOwner(user).
					SetSentence(item.RawSentence))
			}
		}
	}

	return wordBulk
}

func generateSentences(file *os.File) []*sentence {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	sentences := []*sentence{}
	for fileScanner.Scan() {
		if _, err := strconv.Atoi(fileScanner.Text()); err == nil {
			fileScanner.Scan()
			timeRanges := fileScanner.Text()
			fileScanner.Scan()
			for {
				if text := fileScanner.Text(); len(text) != 0 && !strings.Contains(text, "<b>") {
					fileScanner.Scan()
					sentences = append(sentences, generateSentence(timeRanges, text))
				} else {
					break
				}
			}
		}
	}
	return sentences
}
