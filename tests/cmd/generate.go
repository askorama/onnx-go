package main

import (
	"flag"
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
		b, err := ioutil.ReadFile(*testdir + file.Name() + "/model.onnx")
		if err != nil {
			log.Fatal(err)
		}
		model := new(onnx.ModelProto)
		err = model.XXX_Unmarshal(b)
		if err != nil {
			log.Fatal(err)
		}
		pretty.Print(model)
	}
}
