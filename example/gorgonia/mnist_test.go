package main

import (
	"fmt"
	"log"
	"testing"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	"gorgonia.org/gorgonia/debugger/dot"
	gorgonnx "gorgonia.org/gorgonia/onnx"
)

func TestMain(t *testing.T) {
	graph := gorgonnx.NewGraph()
	err := onnx.Unmarshal(mnist.GetMnist(), graph)
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Marshal(graph)
	if err != nil {
		log.Fatal(err)
	}
	//  fmt.Println(string(b))
	// create a VM to run the program on
	machine := gorgonnx.NewTapeMachine(graph)

	// Run the program
	err = machine.RunAll()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(graph)
	log.Println("....")
	vals := graph.GetOutputValues()
	fmt.Println(vals)
	for _, v := range vals {
		fmt.Println(v.Data())
	}
}
