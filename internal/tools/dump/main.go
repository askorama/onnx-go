package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/sanity-io/litter"
)

func main() {
	onnxFile := os.Args[1]
	b, err := ioutil.ReadFile(onnxFile)
	if err != nil {
		log.Fatal(err)
	}
	var m pb.ModelProto
	err = m.XXX_Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}
	litter.Dump(m)

}
