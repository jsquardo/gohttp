package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = 8000
	listenAt := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})
	log.Printf("Open the following URL: http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(listenAt, nil))
}
