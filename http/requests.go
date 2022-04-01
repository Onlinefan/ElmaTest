package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//GetBody посылает get запрос и возвращает тело ответа
func GetBody(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}
