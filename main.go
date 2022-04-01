package main

import (
	"ElmaTest/url_words_counter"
)

func main() {
	urls, word, routines := getTestData()
	counterStruct := url_words_counter.TotalWordsUrlCounter{Urls: urls, Word: word, TotalRoutines: routines}
	counterStruct.CountTotalWords()
}

//getTestData тестовые данные
func getTestData() ([]string, string, int) {
	return []string{
			"https://golang.org", "https://vk.com", "https://www.google.ru", "https://myshows.me/",
			"https://www.jetbrains.com/ru-ru/go/", "https://golang.org", "https://vk.com", "https://www.google.ru",
			"https://myshows.me/", "https://www.jetbrains.com/ru-ru/go/", "https://www.jetbrains.com/ru-ru/go/",
		},
		"go",
		5
}
