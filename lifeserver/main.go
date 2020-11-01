package main

import (
	"encoding/json"
	"fmt"
	"lifeserver/life"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LifeResult contains the results of a Game of Life simulation.
type LifeResult struct {
	Width  int     `json:"Width"`
	Result [][]int `json:"Result"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/life", handleLife)

	fmt.Println("Starting server at port 1000")

	log.Fatal(http.ListenAndServe(":1000", router))
}

func handleLife(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	result := life.RunLife(30)
	lifeResult := LifeResult{30, result}
	err := json.NewEncoder(w).Encode(lifeResult.Result)
	if err != nil {
		log.Fatal(err)
	}

}
