package sentences

import (
	"bufio"
	"interface_project/ent"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetSentences(client *ent.Client, file *os.File, user *ent.User) []*ent.WordCreate {
	return makeWordsBulk(client, generateSentences(file), user)
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

func makeWordsBulk(client *ent.Client, sentences []*sentence, user *ent.User) []*ent.WordCreate {
	bulk := []*ent.WordCreate{}
	for _, item := range sentences {
		for _, token := range item.tokens {
			times := strings.Split(item.timeRange, " ")

			start, _ := time.Parse("12:12:12,123", times[0])
			end, _ := time.Parse("12:12:12,123", times[2])

			bulk = append(bulk, client.Word.Create().SetDuration(item.timeRange).SetMeaning("").SetSentence(item.RawSentence).SetTitle(token).SetUser(user).SetEnd(end).SetStart(start))
		}
	}
	return bulk
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
