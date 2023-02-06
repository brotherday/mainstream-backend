package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Comment struct {
	ID   string
	Body string `json:"body"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	r.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		file := r.PostFormValue("file")

		fmt.Fprintf(w, "%s\n", file)
	})

	r.HandleFunc("/retrieve/{epoch}/{num}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		epoch, num := vars["epoch"], vars["num"]

		fmt.Fprintf(w, "You've requested the presentation %s/%s\n", epoch, num)
	})

	r.HandleFunc("/comment/{epoch}/{num}/{commentID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		epoch, num, commentID := vars["epoch"], vars["num"], vars["commentID"]

		fmt.Fprintf(w, "You've requested the comment %s for presentation %s/%s \n", commentID, epoch, num)
	})

	r.HandleFunc("/comment/add/{epoch}/{num}/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		epoch, num := vars["epoch"], vars["num"]

		now := time.Now()

		var comment Comment
		_ = json.NewDecoder(r.Body).Decode(&comment)
		comment.ID = now.String()

		// json.NewEncoder(w).Encode(&comment)
		fmt.Fprintf(w, "The comment %s \n", comment.Body)
		fmt.Fprintf(w, "You've made a comment %s for presentation %s/%s\n", comment, epoch, num)
	})

	http.ListenAndServe(":3000", r)
}
