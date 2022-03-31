package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// GET /view/title:string
	http.HandleFunc("/view/", makeHandler(viewHandler))

	// GET /edit/title:string
	http.HandleFunc("/edit/", makeHandler(editHandler))

	// POST /save/title:string
	http.HandleFunc("/save/", makeHandler(saveHandler))

	fmt.Println("Server listening at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
