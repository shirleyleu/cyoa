package main

import (
	"fmt"
	"github.com/shirleyleu/cyoa/cyoa"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.New("CYOA").Parse(`
<!doctype html>

<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
        <li><a href="/{{.Arc}}">{{.OptionText}}</a></li>
        {{end}}
    </ul>
</body>
</html>`))

type storyHandler struct{
	story cyoa.Story
}

func (h storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dest := r.URL.Path[1:]
	if dest == "" {
		dest = "intro"
	}
	if v, ok := h.story[dest]; ok {
		err := tmpl.Execute(w, v)
		if err !=nil {
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
			log.Printf("tmpl.Execute: %s", err)
		}
		return
	}
	http.Error(w, "Chapter not found.",http.StatusNotFound)
}

func main() {
	m, err := cyoa.ParseJSON("../gopher.json")
	if err != nil {
		log.Fatalf("Unable to parse JSON: %s", err)
	}
	http.Handle("/", storyHandler{m})
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
