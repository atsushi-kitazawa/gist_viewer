package gist

import (
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
)

var gist []*Gist

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

func NewGist(username string) ([]*Gist) {
	response := getGistList(username)
	if err := json.Unmarshal(response, &gist); err != nil {
		log.Fatal(err)
	}
	return gist
}

func GetRawUrl(filename string) string {
    for _, g := range gist {
	if v, ok := g.Files[filename]; ok {
	    return v.RawUrl
	}
    }
    return ""
}

func GetUrl(filename string) string {
    for _, g := range gist {
	if _, ok := g.Files[filename]; ok {
	    return g.Url
	}
    }
    return ""
}

func GetId(filename string) string {
    for _, g := range gist {
	if _, ok := g.Files[filename]; ok {
	    return g.Id
	}
    }
    return ""
}
func GetContent(rawurl string) string {
    response ,err := http.Get(rawurl)
    if err != nil {
	log.Fatal(err)
    }
    bytes, _ := ioutil.ReadAll(response.Body)
    return string(bytes)
}
