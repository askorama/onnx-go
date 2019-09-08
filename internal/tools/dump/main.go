package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/owulveryck/onnx-go/internal/pb-onnx"
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

	scs := spew.ConfigState{
		Indent:                  "\t",
		DisablePointerAddresses: true,
	}
	scs.Dump(m)
}
