package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	dir := http.Dir(".")

	fileServer := http.FileServer(dir)

	hideHiddenFiles := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/.") {
			http.NotFound(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	})
	
	http.Handle("/", hideHiddenFiles)

	log.Println("starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
