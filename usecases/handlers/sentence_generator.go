package handlers

import (
	"bufio"
	"fmt"
	"interface_project/ent"
	"os"
	"strconv"
)

func MakeWordsBulk(client *ent.Client, sentences []*sentence) []*ent.WordCreate {
	bulk := []*ent.WordCreate{}
	for _, item := range sentences {
		for _, token := range item.tokens {
			bulk = append(bulk, client.Word.Create().SetDuration(item.timeRange).SetMeaning("").SetSentence(item.RawSentence).SetTitle(token))
		}
	}
	return bulk
}

func GenerateSentences(file *os.File) []*sentence {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	sentences := []*sentence{}
	for fileScanner.Scan() {
		if _, err := strconv.Atoi(fileScanner.Text()); err == nil {
			fileScanner.Scan()
			timeRanges := fileScanner.Text()
			fileScanner.Scan()
			for {
				if text := fileScanner.Text(); len(text) != 0 {
					fileScanner.Scan()
					sentences = append(sentences, generateSentence(timeRanges, text))
				} else {
					break
				}
			}

		}
	}
	for _, sentence := range sentences {
		fmt.Printf("- %+v\n", sentence)
	}
	return sentences
}
