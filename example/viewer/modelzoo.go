package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gogo/protobuf/proto"
	"github.com/kr/pretty"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
)

func main() {
	b, err := ioutil.ReadFile("/Users/olivier.wulveryck/Documents/squeezenet/model.onnx")
	if err != nil {
		log.Fatal(err)
	}
	model := &pb.ModelProto{}
	err = proto.Unmarshal(b, model)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%# v", pretty.Formatter(model))

}
