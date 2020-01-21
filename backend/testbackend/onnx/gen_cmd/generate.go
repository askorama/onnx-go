package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
)

var (
	testdir   *string
	outputdir *string
)

func main() {
	testdir = flag.String("testpath", ".", "path to the onnx test directory")
	outputdir = flag.String("outputdir", "", "path to the outputdir")
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
	testcases := make([]testCases, 0)
	for _, file := range files {
		if !file.IsDir() {
			return
		}
		elements := re.FindAllStringSubmatch(file.Name(), -1)
		if len(elements) == 0 {
			continue
		}
		log.Println("-->", file.Name())
		optype, testtitle, err := processFile(file)
		if err != nil {
			log.Println(err)
			continue
		}
		// Add the file to the testcases
		testcases = append(testcases, testCases{
			optype, testtitle, ""})
	}
	/*
		output := os.Stdout
		if *outputdir != "" {
			output, err = os.Create(filepath.Join(*outputdir, "onnx_register_testcases.go"))
			if err != nil {
				log.Fatal(err)
			}
			defer output.Close()
		}
		processTemplate(testCasesTemplate, testcases, output)
	*/
}

// returns the optype and the title
func processFile(file os.FileInfo) (string, string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	var tv testValue
	var mv modelValue
	tv.TestName = toCamelCase(file.Name())
	b, err := ioutil.ReadFile(*testdir + file.Name() + "/model.onnx")
	if err != nil {
		return "", "", err
	}
	tv.ModelB = fmt.Sprintf("%#v", b)
	model := new(ir.ModelProto)
	err = model.XXX_Unmarshal(b)
	if err != nil {
		return "", "", err
	}
	tv.Description = fmt.Sprintf("version: %v. %v", model.GetIrVersion(), model.GetDocString())
	mv.TestName = fmt.Sprintf("%#v", file.Name())
	mv.IrVersion = fmt.Sprintf("%v", model.IrVersion)
	mv.OpsetVersion = fmt.Sprintf("%v", model.OpsetImport[0].Version)
	mv.NodeProto = make([]nodeProtoValue, len(model.Graph.Node))
	for i := range model.Graph.Node {
		mv.NodeProto[i] = nodeProtoValue{
			Input:         fmt.Sprintf("%#v", model.Graph.Node[i].Input),
			Output:        fmt.Sprintf("%#v", model.Graph.Node[i].Output),
			Name:          fmt.Sprintf("%#v", model.Graph.Node[i].Name),
			OpType:        fmt.Sprintf("%#v", model.Graph.Node[i].OpType),
			AttributeDump: spew.Sdump(model.Graph.Node[i].Attribute),
		}
	}
	tv.OpType = model.Graph.Node[0].OpType
	// There should be only one node
	if len(model.GetGraph().GetNode()) > 1 {
		return "", "", fmt.Errorf("graph with more than one node not supported by this utility")
	}
	node := model.GetGraph().GetNode()[0]

	processModelGraphInput(model, &mv)

	err = processModelGraphNodeInput(file.Name(), node, &tv)
	if err != nil {
		return "", "", err
	}

	processModelGraphOutput(model, &mv)

	err = processModelGraphNodeOutput(file.Name(), node, &tv)
	if err != nil {
		return "", "", err
	}

	processModelGraphValueInfo(model, &mv)

	// TestTemplate
	output := os.Stdout
	outputTest := os.Stdout
	if *outputdir != "" {
		output, err = os.Create(filepath.Join(*outputdir, "onnx_"+file.Name()+".go"))
		if err != nil {
			return "", "", err
		}
		defer output.Close()
		outputTest, err = os.Create(filepath.Join(*outputdir, "onnx_"+file.Name()+"_test.go"))
		if err != nil {
			return "", "", err
		}
		defer outputTest.Close()
	}
	tv.ModelValue = mv
	err = processTemplate(testTemplate, tv, output)
	if err != nil {
		return "", "", err
	}
	err = processTemplate(testTestTemplate, tv, outputTest)
	if err != nil {
		return "", "", err
	}
	/*
		err = processTemplate(modelTemplate, mv, output)
		if err != nil {
			return err
		}
	*/
	return mv.NodeProto[0].OpType, tv.TestName, nil
}

func processModelGraphInput(model *ir.ModelProto, mv *modelValue) {
	mv.Input = make([]valueInfoProto, len(model.Graph.Input))
	for i := range model.Graph.Input {
		mv.Input[i] = valueInfoProto{
			Name:     model.Graph.Input[i].Name,
			ElemType: fmt.Sprintf("%v", model.Graph.Input[i].Type.GetTensorType().ElemType),
			Dims:     make([]string, len(model.Graph.Input[i].Type.GetTensorType().Shape.Dim)),
		}
		for j, v := range model.Graph.Input[i].Type.GetTensorType().Shape.Dim {
			mv.Input[i].Dims[j] = fmt.Sprintf("%v", v.GetValue().(*ir.TensorShapeProto_Dimension_DimValue).DimValue)
		}
	}
}

func processModelGraphOutput(model *ir.ModelProto, mv *modelValue) {
	mv.Output = make([]valueInfoProto, len(model.Graph.Output))
	for i := range model.Graph.Output {
		mv.Output[i] = valueInfoProto{
			Name:     model.Graph.Output[i].Name,
			ElemType: fmt.Sprintf("%v", model.Graph.Output[i].Type.GetTensorType().ElemType),
			Dims:     make([]string, len(model.Graph.Output[i].Type.GetTensorType().Shape.Dim)),
		}
		for j, v := range model.Graph.Output[i].Type.GetTensorType().Shape.Dim {
			mv.Output[i].Dims[j] = fmt.Sprintf("%v", v.GetValue().(*ir.TensorShapeProto_Dimension_DimValue).DimValue)
		}
	}
}

func processModelGraphValueInfo(model *ir.ModelProto, mv *modelValue) {
	mv.ValueInfo = make([]valueInfoProto, len(model.Graph.ValueInfo))
	for i := range model.Graph.ValueInfo {
		mv.ValueInfo[i] = valueInfoProto{
			Name:     model.Graph.ValueInfo[i].Name,
			ElemType: fmt.Sprintf("%v", model.Graph.ValueInfo[i].Type.GetTensorType().ElemType),
			Dims:     make([]string, len(model.Graph.ValueInfo[i].Type.GetTensorType().Shape.Dim)),
		}
		for j, v := range model.Graph.ValueInfo[i].Type.GetTensorType().Shape.Dim {
			mv.ValueInfo[i].Dims[j] = fmt.Sprintf("%v", v.GetValue().(*ir.TensorShapeProto_Dimension_DimValue).DimValue)

		}
	}
}

func processModelGraphNodeInput(filename string, node *ir.NodeProto, tv *testValue) error {
	tv.Input = make([]iO, len(node.GetInput()))
	for i := range node.GetInput() {
		// Open the tensorproto sample file
		filepath := fmt.Sprintf("%v%v/test_data_set_0/input_%v.pb", *testdir, filename, i)
		b, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}
		sampleTestData := new(ir.TensorProto)
		err = sampleTestData.XXX_Unmarshal(b)
		if err != nil {
			return err
		}
		t, err := sampleTestData.Tensor()
		if err != nil {
			return err
		}
		data := fmt.Sprintf("%#v", t.Data())
		shape := fmt.Sprintf("%#v", t.Shape())
		if len(t.Shape()) == 1 && t.Shape()[0] == 1 {
			data = fmt.Sprintf("[]float32{%v}", t.Data())
		}
		if len(t.Shape()) == 0 {
			data = fmt.Sprintf("[]float32{%v}", t.Data())
			shape = "(1)"
		}
		tv.Input[i] = iO{
			Shape: shape,
			//Data:  fmt.Sprintf("%#v", t.Data()),
			Data: data,
		}
	}
	return nil
}

func processModelGraphNodeOutput(filename string, node *ir.NodeProto, tv *testValue) error {
	tv.ExpectedOutput = make([]iO, len(node.GetOutput()))
	for i := range node.Output {
		// Open the tensorproto sample file
		filepath := fmt.Sprintf("%v%v/test_data_set_0/output_%v.pb", *testdir, filename, i)
		b, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}
		sampleTestData := new(ir.TensorProto)
		err = sampleTestData.XXX_Unmarshal(b)
		if err != nil {
			return err

		}
		t, err := sampleTestData.Tensor()
		if err != nil {
			return err
		}
		shape := fmt.Sprintf("%#v", t.Shape())
		data := fmt.Sprintf("%#v", t.Data())
		if len(t.Shape()) == 1 && t.Shape()[0] == 1 {
			data = fmt.Sprintf("[]float32{%v}", t.Data())
		}
		if len(t.Shape()) == 0 {
			data = fmt.Sprintf("[]float32{%v}", t.Data())
			shape = "(1)"
		}
		tv.ExpectedOutput[i] = iO{
			Shape: shape,
			Data:  data,
		}
	}
	return nil
}

func processTemplate(t *template.Template, v interface{}, output io.Writer) error {
	var buf bytes.Buffer
	if err := t.Execute(&buf, v); err != nil {
		log.Fatal(err)
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("Cannot format", err)
	}
	_, err = output.Write(p)
	return err
}

var link = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z0-9])")

func toCamelCase(str string) string {
	return link.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}
