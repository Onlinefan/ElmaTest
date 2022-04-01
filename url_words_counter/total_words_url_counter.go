package url_words_counter

import (
	httpHelper "ElmaTest/http"
	"fmt"
	"strings"
	"sync"
)

//TotalWordsUrlCounter объявление структуры
type TotalWordsUrlCounter struct {
	Urls          []string
	Word          string
	TotalRoutines int
}

//CountTotalWords основной метод подсчета слов
func (counterStruct *TotalWordsUrlCounter) CountTotalWords() {
	countRoutineChannel := make(chan int, counterStruct.TotalRoutines)
	var waitGroup sync.WaitGroup
	total := 0

	for _, url := range counterStruct.Urls {
		waitGroup.Add(1)
		countRoutineChannel <- 1
		go counterStruct.countWordsByUrl(url, countRoutineChannel, &waitGroup, &total)
	}

	waitGroup.Wait()
	fmt.Println("Total:", total)
}

//countWordsByUrl подсчет слов по конкретному урлу
func (counterStruct *TotalWordsUrlCounter) countWordsByUrl(url string, countRoutineChannel <-chan int, waitGroup *sync.WaitGroup, total *int) {
	defer waitGroup.Done()
	body := httpHelper.GetBody(url)
	count := strings.Count(string(body), counterStruct.Word)
	fmt.Println("Count for", url, ":", count)
	defer func() {
		<-countRoutineChannel
		*total += count
	}()
}
