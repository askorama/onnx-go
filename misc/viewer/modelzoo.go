package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/kr/pretty"
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
)

func main() {
	b, err := os.ReadFile("/Users/olivier.wulveryck/Documents/squeezenet/model.onnx")
	if err != nil {
		log.Fatal(err)
	}
	model := &ir.ModelProto{}
	err = proto.Unmarshal(b, model)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%# v", pretty.Formatter(model))

}
