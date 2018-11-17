package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Create struct representing story structure
type Story map[string]Chapter

type Chapter struct{
	Title     string    `json:"title"`
	Paragraphs []string `json:"story"`
	Options   []option  `json:"options"`
}

type option struct{
	OptionText string `json:"text"`
	Arc string `json:"arc"`
}

func ParseJSON(filepath string) (Story, error){
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	m := make(map[string]Chapter)
	if err := json.Unmarshal(data, &m); err !=nil{
		return nil, err
	}
	return m, nil
}
