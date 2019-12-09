package main

import (
	"fmt"
	"net/http"

	"urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/make-school": "https://www.makeschool.com",
		"/bew2.5":      "https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Ecosystems",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the JSONHandler using the mapHandler as the
	// fallback
	json := `
	[
		{
			"path": "",
			"url": ""
		}
	]
`
	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
