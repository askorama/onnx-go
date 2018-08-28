package onnx_test

import (
	"io/ioutil"
	"log"

	onnx "github.com/owulveryck/onnx-go"
)

func ExampleModelProto() {
	b, err := ioutil.ReadFile("/path/to/onnx/file.onnx")
	if err != nil {
		log.Fatal(err)
	}
	model := new(onnx.ModelProto)
	err = model.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(model)
}
