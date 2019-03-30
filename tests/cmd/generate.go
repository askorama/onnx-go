package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/kr/pretty"
	onnx "github.com/owulveryck/onnx-go/internal/pb-onnx"
)

func main() {
	testdir := flag.String("testpath", ".", "path to the onnx test directory")
	op := flag.String("op", "", "the operator who needs tests")
	flag.Parse()
	if *op == "" {
		flag.Usage()
		os.Exit(0)
	}
	// locate all the directories with the pattern test_op_...
	files, err := ioutil.ReadDir(*testdir)
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile("^test_" + *op + "(_*)(.*)")
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		elements := re.FindAllStringSubmatch(file.Name(), -1)
		if len(elements) == 0 {
			continue
		}
		fmt.Printf("var %v = tests.TestCase{", file.Name())
		b, err := ioutil.ReadFile(*testdir + file.Name() + "/model.onnx")
		if err != nil {
			log.Fatal(err)
		}
		model := new(onnx.ModelProto)
		err = model.XXX_Unmarshal(b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Model:")
		pretty.Print(model)
		fmt.Println(",")
		// There should be only one node
		if len(model.GetGraph().GetNode()) > 1 {
			log.Fatal("Not supported")
		}
		node := model.GetGraph().GetNode()[0]
		fmt.Printf("Inputs: []pb.TensorProto{\n")
		for i := range node.GetInput() {
			// Open the tensorproto sample file
			filename := fmt.Sprintf("%v%v/test_data_set_0/input_%v.pb", *testdir, file.Name(), i)
			b, err = ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			sampleTestData := new(onnx.TensorProto)
			err = sampleTestData.XXX_Unmarshal(b)
			if err != nil {
				log.Fatal(err)
			}
			//t, err := sampleTestData.Tensor()
			//if err != nil {
			//	log.Fatal(err)
			//}
			pretty.Print(sampleTestData)
			fmt.Println(",")
		}
		pretty.Printf("}\n")
		fmt.Printf("Output: []pb.TensorProto{\n")
		for i := range node.GetOutput() {
			// Open the tensorproto sample file
			filename := fmt.Sprintf("%v%v/test_data_set_0/output_%v.pb", *testdir, file.Name(), i)
			b, err = ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			sampleTestData := new(onnx.TensorProto)
			err = sampleTestData.XXX_Unmarshal(b)
			if err != nil {
				log.Fatal(err)
			}
			//t, err := sampleTestData.Tensor()
			//if err != nil {
			//		log.Fatal(err)
			//}
			pretty.Print(sampleTestData)
			fmt.Println(",")
		}
		pretty.Printf("}\n")
	}
}
