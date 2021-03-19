package gist

import (
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
)

type Gist struct {
	Url   string               `json:"url"`
	Id    string               `json:"id"`
	Files map[string]FilesInfo `json:"files"`
}
type FilesInfo struct {
	Filename string `json:"filename"`
	RawUrl   string `json:"raw_url"`
}

func getGistList(username string) []byte {
	url := "https://api.github.com/users/" + username + "/gists"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(response.Body)

	return bytes
}

func NewGist(username string) ([]Gist) {
	response := getGistList(username)
	var gist []Gist
	if err := json.Unmarshal(response, &gist); err != nil {
		log.Fatal(err)
	}
	return gist
}


