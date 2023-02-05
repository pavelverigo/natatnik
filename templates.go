package main

import "html/template"

var (
	noteTemplate     *template.Template
	noteListTemplate *template.Template
)

func populateTemplates() {
	t := template.Must(template.ParseFiles("note.html", "notelist.html"))
	noteTemplate = t.Lookup("note.html")
	noteListTemplate = t.Lookup("notelist.html")
}

type noteData struct {
	Title   string
	Content string
}

type noteListData struct {
	Notes map[string]string
}
