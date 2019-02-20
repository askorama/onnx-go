package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gorgonia.org/gorgonia/debugger"
	"gorgonia.org/gorgonia/debugger/dot"
	"gorgonia.org/gorgonia/node"
	gorgonnx "gorgonia.org/gorgonia/onnx"
	"gorgonia.org/tensor"
)

func main() {

	graph := gorgonnx.NewGraph()
	m := onnx.NewModel(graph)
	b := mnist.GetMnist()
	err := m.Unmarshal(b)
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
	debugC := make(chan debugger.DebugMsg, 0)
	//  fmt.Println(string(b))
	// create a VM to run the program on
	machine := gorgonnx.NewTapeMachine(graph,
		gorgonnx.WithDebuggingChannel(debugC))

	go func(c chan debugger.DebugMsg) {
		for msg := range c {
			pload := bytes.NewBuffer(msg.Payload)
			dec := gob.NewDecoder(pload)
			var instr debugger.Instruction
			err := dec.Decode(&instr)
			if err != nil {
				log.Println(err)
			}
			log.Println(instr)
			for err == nil {
				var input tensor.Dense
				err = dec.Decode(&input)
				if err == nil {
					log.Println("Value:", input)
				}
			}
			if err != nil && err != io.EOF {
				log.Println(err)
			}
		}
	}(debugC)

	// Run the program
	err = machine.RunAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range m.Output {
		fmt.Println(graph.Node(v).(node.Node).Value().Data())
	}
	/*
		// DEBUGGING
		mp := &pb.ModelProto{}
		err = proto.Unmarshal(b, mp)
		if err != nil {
			log.Fatal(err)
		}
		w := os.Stdout
		fmt.Fprintf(w, "const dir = \"numpy/\"\n")
		for i, nodeProto := range mp.Graph.Node {
			fmt.Fprintf(w, "func Test%v%v(t *testing.T) {\n", nodeProto.OpType, i)
			fmt.Fprintf(w, "save := dir + \"%v%v\"\n", nodeProto.OpType, i)
			fmt.Fprintf(w, "os.MkdirAll(save, os.ModePerm)\n")
			// Input
			for i := range nodeProto.Input {
				n, ok := m.GetNodeByName(nodeProto.Input[i])
				if !ok {
					log.Fatalf("Node %v not found", nodeProto.Input[i])
				}
				//fmt.Printf("[%v] Input %v: %v\n", nodeProto.OpType, i, n.(node.Node).Value().Dtype())
				writeTensorTo(w, fmt.Sprintf("input%v", i), n.(node.Node))
			}
			// Output
			if len(nodeProto.Output) != 1 {
				log.Fatal("Weird")
			}
			n, ok := m.GetNodeByName(nodeProto.Output[0])
			if !ok {
				log.Fatalf("Node %v not found", nodeProto.Output[0])
			}
			writeTensorTo(w, "output", n.(node.Node))
			//fmt.Printf("[%v] Output: %v\n", nodeProto.OpType, n.(node.Node).Value())
			fmt.Fprintf(w, "}\n")
		}
	*/
}

func writeTensorTo(w io.Writer, name string, n node.Node) error {
	fmt.Fprintf(w, "f%v, err := os.Create(save + \"/%v.np\")\n", name, name)
	fmt.Fprintf(w, "if err != nil {\n")
	fmt.Fprintf(w, "t.Fatal(err)\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "defer f%v.Close()\n", name)
	fmt.Fprintf(w, "%v := tensor.New(\n", name)
	fmt.Fprintf(w, "\ttensor.WithShape%v,\n", n.Value().Shape())
	fmt.Fprintf(w, "\ttensor.Of(tensor.Dtype{reflect.TypeOf(%v(1))}),\n", n.Value().Dtype())
	fmt.Fprintf(w, "\ttensor.WithBacking(%#v),\n", n.Value().Data())
	fmt.Fprintf(w, ")\n")
	fmt.Fprintf(w, "err = %v.WriteNpy(f%v)\n", name, name)
	fmt.Fprintf(w, "if err != nil {\n")
	fmt.Fprintf(w, "t.Fatal(err)\n")
	fmt.Fprintf(w, "}\n")

	return nil
}
