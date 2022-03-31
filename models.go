package main

import "io/ioutil"

// Page model used to construct the content of a web page.
type Page struct {
	Title string
	Body  []byte
}

// save writes file .txt to data/ directory and
// return error if there is a problem while writing.
func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage reads file .txt in data/ directory and
// return error if the file does not exist.
func loadPage(title string) (*Page, error) {
	var (
		filename  = "data/" + title + ".txt"
		body, err = ioutil.ReadFile(filename)
	)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
