package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// load data
	err := LoadData()
	if err != nil {
		log.Fatalf("Failed to load dummy data : ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, "Hello World from GO http server")

	})

	http.HandleFunc("/todos", TodosHandler)
	http.HandleFunc("/todos/", TodosByIdHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
