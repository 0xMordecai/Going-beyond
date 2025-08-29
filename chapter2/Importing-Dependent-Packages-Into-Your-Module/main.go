package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{text}", name)
	log.Fatal(http.ListenAndServe(":8080", r))
}
func name(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, _ := template.ParseFiles("static/text.xhtml")
	t.Execute(w, vars["text"])
}
