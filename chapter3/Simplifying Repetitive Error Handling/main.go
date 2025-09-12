package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func helperUnmarshal() (person Person) {
	r, err := http.Get("https://swapi.dev/api/people/1")
	check(err, "Calling SW people API")
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	check(err, "Read JSON from response")
	err = json.Unmarshal(data, &person)
	check(err, "Unmarshalling")
	return person
}
func check(err error, msg string) {
	if err != nil {
		log.Println("Error encountered:", msg)
		// do common error-handling stuff}
	}
}
