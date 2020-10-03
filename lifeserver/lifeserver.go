package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/life", handleLife)

	fmt.Println("Starting server at port 8080")

	http.ListenAndServe(":8080", nil)

}

func handleLife(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}
