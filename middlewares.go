package main

import (
	"net/http"
	"regexp"
)

// validPath initializes regexp rule to get a valid URL.
//
// e.g. /edit/lorem-ipsum -> kebab-case
// 		/edit/LoremIpsum -> PascalCase
// 		/edit/loremIpsum -> camelCase
var validPath = regexp.MustCompile("^/(edit|save|view)/([-a-zA-Z0-9]+)$")

// makeHandler used to validate the URLs and send the requested wiki title.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := validPath.FindStringSubmatch(r.URL.Path)
		if path == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, path[2])
	}
}
