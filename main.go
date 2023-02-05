package main

import (
	"log"
	"net/http"
)

var notes map[string]string

func main() {
	populateTemplates()
	notes = map[string]string{
		"tears": "Tears in heaven",
		"bad":   "Yes I'm everybody's Mr Bad Guy",
	}

	http.HandleFunc("/note/", noteHandler)
	http.Handle("/static/", http.FileServer(http.Dir("public")))

	log.Fatalln(http.ListenAndServe(":5000", nil))
}

func noteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		title := r.URL.Path[len("/note/"):]
		if title == "" {
			err := noteListTemplate.Execute(w, noteListData{notes})
			if err != nil {
				log.Fatalln(err)
			}
			return
		}

		note, ok := notes[title]
		if ok {
			err := noteTemplate.Execute(w, noteData{title, note})
			if err != nil {
				log.Fatalln(err)
			}
			return
		} else {
			http.NotFound(w, r)
			return
		}
	}
	if r.Method == "POST" {
		if r.URL.Path != "/note/" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")
		notes[title] = content
		http.Redirect(w, r, "/note/"+title, http.StatusFound)
	}
}
