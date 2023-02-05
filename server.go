package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	http.ListenAndServe(":3000", r)
}
