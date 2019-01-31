package main

import (
	"fmt"
	"log"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	"gorgonia.org/gorgonia/debugger/dot"
	"gorgonia.org/gorgonia/node"
	gorgonnx "gorgonia.org/gorgonia/onnx"
)

func main() {

	graph := gorgonnx.NewGraph()
	m := onnx.NewModel(graph)
	err := m.Unmarshal(mnist.GetMnist())
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Marshal(graph)
	if err != nil {
		log.Fatal(err)
	}

	if len(m.Input) != 1 {
		log.Fatal("Expected only one input")
	}
	err = gorgonnx.Let(graph.Node(m.Input[0]).(node.Node), nil)
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
	for _, v := range m.Output {
		fmt.Println(graph.Node(v).(node.Node).Value().Data())
	}
}
