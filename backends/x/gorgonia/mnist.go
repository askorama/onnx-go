// +build ignore

package main

import (
	"log"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gorgonia.org/gorgonia/debugger/dot"
	"gorgonia.org/gorgonia/node"
	gorgonnx "gorgonia.org/gorgonia/onnx"
)

func main() {

	graph := gorgonnx.NewGraph()
	m := onnx.NewModel(graph)
	b := mnist.GetMnist()
	err := m.Decode(b)
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Marshal(graph)
	if err != nil {
		log.Fatal(err)
	}

	sampleTestData := new(pb.TensorProto)
	err = sampleTestData.XXX_Unmarshal(mnist.GetTest1Input0())
	if err != nil {
		log.Fatal(err)
	}
	t, err := sampleTestData.Tensor()
	if err != nil {
		log.Fatal(err)
	}

	if len(m.Input) != 1 {
		log.Fatal("Expected only one input")
	}
	err = gorgonnx.Let(graph.Node(m.Input[0]).(node.Node), t)
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
		log.Println(graph.Node(v).(node.Node).Value().Data())
	}
}
