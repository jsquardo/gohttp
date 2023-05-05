package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8000

func main() {
	listenAt := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})
	log.Printf("Open the following URL: http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(listenAt, nil))
}
