package main

import (
	"fmt"
	"github.com/shirleyleu/cyoa/storymaker"
	"log"
	"net/http"
)

type storyHandler struct{
	story storymaker.Story
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dest := r.URL.Path[1:]
	if dest == "" {
		dest = "intro"
	}
	v, ok := s.story[dest]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(v)
}

func main() {
	m, err := storymaker.ParseJSON("../gopher.json")
	if err != nil {
		log.Fatalf("Unable to parse JSON: %s", err)
	}
	http.Handle("/", storyHandler{m})
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}




// Create HTML index that starts with intro story
