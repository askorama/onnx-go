// This file contains specific handler to run the demos

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/image", imagePostHandler)
	http.HandleFunc("/model", modelPostHandler)
}

type img struct {
	Image interface{}
	Bla   string `json:"bla"`
}

func imagePostHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var payload img
	err := dec.Decode(&payload)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	fmt.Fprintf(w, `{result: "ok"}`)
}
func modelPostHandler(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

}
