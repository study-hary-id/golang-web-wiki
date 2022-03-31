package main

import (
	"html/template"
	"net/http"
)

// templates initialize available html template in tmpl/ directory.
var templates = template.Must(template.ParseGlob("tmpl/*"))

// TODO: Delete this function if not used.
// resBadRequest serves response with Status Bad Request format.
func resBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400 status bad request"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//resInternalServerError serves response with Status Internal Server Error format.
func resInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// renderTemplate used to render template page in html form.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		resInternalServerError(w, err)
		return
	}
}
