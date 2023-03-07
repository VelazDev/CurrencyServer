package utils

import (
	"GolangCurrencyMS/internal/payload"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Get Data from server using API request
func GetDataFromAPI[Req payload.Request](url string, format Req) Req {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &format)
	fmt.Printf("%+v\n", format)

	return format
}

// Get all Symbol names from Server
func GetAllKeyValues(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	c := make(map[string]json.RawMessage)
	e := json.Unmarshal([]byte(body), &c)
	if e != nil {
		panic(e)
	}

	IdList := make([]string, len(c))
	i := 0
	for s, _ := range c {
		IdList[i] = s
		i++
	}

	return IdList
}
