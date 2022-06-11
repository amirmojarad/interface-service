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

// func getStartAndEnd(duration string) (time.Time, time.Time) {
// 	strs := strings.Split(duration, " ")
// 	startAndEnd := make([]time.Time, 2)
// 	for i, item := range strs {
// 	if i == 2 {
// 	continue
// 	}
// 	if i == 1 {
// 		hmsString := strings.Split(item, ",")[0]
// 		millsecond := strings.Split(item, ",")[1]
// 		hmsList := strings.Split(hmsString, ":")
// 		hour, _ := strconv.Atoi(hmsList[0])
// 		min, _ := strconv.Atoi(hmsList[1])
// 		sec, _ := strconv.Atoi(hmsList[2])

// 		startAndEnd[0] = time.Date(0,0,0, hour,min,sec, )
// 	}

// 	}
// 	return nil, nil
// }

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
		//times := strings.Split(item.timeRange, " ")
		//start, _ := time.Parse("12:12:12,123", times[0])
		//end, _ := time.Parse("12:12:12,123", times[2])
		for _, token := range item.tokens {
			if checkWordTitleIsValid(token) {
				if strings.Contains(token, "...") {
					token = strings.Replace(token, "...", "", -1)
				}
				wordBulk = append(wordBulk, client.Word.Create().
					SetTitle(token).
					SetMeaning("").
					SetFile(file).
					SetIsPreposition(preposition.IsPreposition(token)).
					SetDuration(item.timeRange).
					SetSentence(item.RawSentence))
				//wordBulk = append(wordBulk, client.Word.Create().SetDuration(item.timeRange).SetMeaning("").SetSentence(item.RawSentence).SetTitle(token).SetUser(user).SetEnd(end).SetStart(start))
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
