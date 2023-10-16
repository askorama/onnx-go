package main

import (
	"log"
	"os"

	"github.com/owulveryck/onnx-go/internal/onnx/ir"
	"github.com/sanity-io/litter"
)

func main() {
	onnxFile := os.Args[1]
	b, err := os.ReadFile(onnxFile)
	if err != nil {
		log.Fatal(err)
	}
	var m ir.ModelProto
	err = m.XXX_Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}
	litter.Dump(m)

}
