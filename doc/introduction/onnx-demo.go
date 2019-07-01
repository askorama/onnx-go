// This file contains specific handler to run the demos

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

func init() {
	http.HandleFunc("/image", imagePostHandler)
	http.HandleFunc("/model", modelPostHandler)
}

type img struct {
	Image string
	Model string `json:"model"`
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

	var res results
	if md, ok := models[payload.Model]; ok {
		img, err := processPicture(payload.Image, md.height, md.width)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
		output, err := process(img, md.height, md.width, md.table)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if md.postProcessing != nil {
			output = md.postProcessing(output)
		}
		res = classify(output, md.table)
	}
	enc := json.NewEncoder(w)
	enc.Encode(res)
}
func modelPostHandler(w http.ResponseWriter, r *http.Request) {
	// reset the backend and the model
	// START_MODEL OMIT
	var b []byte // will hold the content of the `model.onnx file
	var err error
	// Create a backend receiver
	backend = gorgonnx.NewGraph()
	// Create a model and set the execution backend
	model = onnx.NewModel(backend)
	// ...
	b, err = ioutil.ReadAll(r.Body) // OMIT
	defer r.Body.Close()            // OMIT
	if err != nil {                 // OMIT
		http.Error(w, err.Error(), http.StatusInternalServerError) // OMIT
		log.Println(err)                                           // OMIT
		return                                                     // OMIT
	} // OMIT
	err = model.UnmarshalBinary(b)
	// END_MODEL OMIT
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
