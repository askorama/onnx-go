package main

import "text/template"

var testTemplate = template.Must(template.New("testCase").Parse(testTmpl))
var testTestTemplate = template.Must(template.New("testTestCase").Parse(testTestTmpl))

type testValue struct {
	OpType         string
	TestName       string
	Description    string
	ModelValue     modelValue
	ModelB         string
	Input          []iO
	ExpectedOutput []iO
}
type iO struct {
	Shape string
	Data  string
}

type modelValue struct {
	TestName     string
	IrVersion    string
	OpsetVersion string
	NodeProto    []nodeProtoValue
	Input        []valueInfoProto
	Output       []valueInfoProto
	ValueInfo    []valueInfoProto
}
type valueInfoProto struct {
	Name     string
	ElemType string
	Dims     []string
}

type nodeProtoValue struct {
	Input         string //Input:     []string{"Parameter193", "Parameter193_reshape1_shape"},
	Output        string //Output:    []string{"Parameter193_reshape1"},
	Name          string //Name:      "Times212_reshape1",
	OpType        string //OpType:    "Reshape",
	AttributeDump string
	Attributes    []attribute
}

type attribute struct {
	Name    string
	Type    string
	Ints    string
	I       string
	S       string
	F       string
	T       string
	Tensors string
	Floats  string
	Strings string
}

const testTestTmpl = `
package onnxtest

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
)

func TestNew{{ .TestName }}(t *testing.T) {
	mytest := New{{ .TestName }}()
	var model ir.ModelProto
	err := proto.Unmarshal(mytest.ModelB, &model)
	if err != nil {
		t.Fatal(err)
	}
	if model.Graph == nil {
		t.Fatal("graph is nil")
	}
	if len(model.Graph.Input) != len(mytest.Input) {
		t.Fatalf("invalid test: model has %v input, but test only provide %v", len(model.Graph.Input), len(mytest.Input))
	}
	if len(model.Graph.Output) != len(mytest.ExpectedOutput) {
		t.Fatalf("invalid test: model has %v input, but test only provide %v", len(model.Graph.Output), len(mytest.ExpectedOutput))
	}
}
`

//Model: {{ template "modelCase" .ModelValue }},

const testTmpl = `
package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
    "gorgonia.org/tensor"
  "github.com/owulveryck/onnx-go/backend/testbackend"
)

func init() {
testbackend.Register("{{ .OpType }}","{{ .TestName }}",New{{ .TestName }})
}


// New{{ .TestName }} {{ .Description }}
func New{{ .TestName }}() *testbackend.TestCase{
return &testbackend.TestCase{
	OpType: "{{ .OpType }}",
	Title: "{{ .TestName }}",
	ModelB: {{ .ModelB }},
	{{ template "doc" .ModelValue }}
	Input: []tensor.Tensor{
		{{ range .Input }}
		tensor.New(
			tensor.WithShape{{ .Shape }},
			tensor.WithBacking({{ .Data }}),
		),
		{{ end }}
	},
	ExpectedOutput: []tensor.Tensor{
		{{ range .ExpectedOutput }}
		tensor.New(
			tensor.WithShape{{ .Shape }},
			tensor.WithBacking({{ .Data }}),
		),
		{{ end }}
	},
}
}

{{ define "doc" }}
/*
{{ range .NodeProto }}
&ir.NodeProto{
  Input:     {{ .Input }},
  Output:    {{ .Output }},
  Name:      {{ .Name }},
  OpType:    {{ .OpType }},
  Attributes: {{ .AttributeDump }},
},
{{ end }}

*/
{{ end }} 

{{ define "modelCase" }}
&ir.ModelProto{
      IrVersion: {{ .IrVersion }},
      OpsetImport: []*ir.OperatorSetIdProto{
	      &ir.OperatorSetIdProto{
		      Domain:  "",
		      Version: {{ .OpsetVersion }},
	      },
      },
      Graph: &ir.GraphProto{
	      Node: []*ir.NodeProto{
		      {{ range .NodeProto }}
		      &ir.NodeProto{
			      Input:     {{ .Input }},
			      Output:    {{ .Output }},
			      Name:      {{ .Name }},
			      OpType:    {{ .OpType }},
			      {{ if .Attributes }}
			      Attribute: []*ir.AttributeProto{
			      {{ range .Attributes }}
				  &ir.AttributeProto{
						Name: {{ .Name }},
						Type: {{ .Type }},
						{{ if .Ints }}Ints: {{ .Ints }}, {{ end }}
						{{ if .Tensors }}Tensors: {{ .Tensors }}, {{ end }}
						{{ if .Floats }}Floats: {{ .Floats }}, {{ end }}
						{{ if .String }}Strings: {{ .Strings }}, {{ end }}
						{{ if .S }}S: {{ .S }}, {{ end }}
						{{ if .I }}I: {{ .I }},{{ end }}
						{{ if .F }}F: {{ .F }},{{ end }}
						{{ if .T }}T: {{ .T }},{{ end }}
					},

			      {{ end }}
			      },
			      {{ end }}
		      },
		      {{ end }}
	      },
	      Initializer: []*ir.TensorProto{
		      /*
		      &ir.TensorProto{
			      Dims:      []int64{16, 4, 4, 10},
			      DataType:  1,
			      FloatData: parameter193,
			      Name:      "Parameter193",
		      },
		      */
	      },  
	      Input: []*ir.ValueInfoProto{
		      {{ range .Input }}
		      &ir.ValueInfoProto{
			      Name: "{{ .Name }}",
			      Type: &ir.TypeProto{
				      Value: &ir.TypeProto_TensorType{
					      TensorType: &ir.TypeProto_Tensor{
						      ElemType: {{ .ElemType }},
						      Shape: &ir.TensorShapeProto{
							      Dim: []*ir.TensorShapeProto_Dimension{
								      {{ range .Dims }}
								      &ir.TensorShapeProto_Dimension{
									      Value: &ir.TensorShapeProto_Dimension_DimValue{
									      DimValue: {{ . }},
									      },
								      },
								      {{ end }} 
							      },
						      },
					      },
				      },
			      },
		      },
		      {{ end }}
	      },
	      Output: []*ir.ValueInfoProto{
		      {{ range .Output }}
		      &ir.ValueInfoProto{
			      Name: "{{ .Name }}",
			      Type: &ir.TypeProto{
				      Value: &ir.TypeProto_TensorType{
					      TensorType: &ir.TypeProto_Tensor{
						      ElemType: {{ .ElemType }},
						      Shape: &ir.TensorShapeProto{
							      Dim: []*ir.TensorShapeProto_Dimension{
								      {{ range .Dims }}
								      &ir.TensorShapeProto_Dimension{
									      Value: &ir.TensorShapeProto_Dimension_DimValue{
									      DimValue: {{ . }},
									      },
								      },
								      {{ end }} 
							      },
						      },
					      },
				      },
			      },
		      },
		      {{ end }}
	      },
	      {{ if .ValueInfo }}
	      ValueInfo: []*ir.ValueInfoProto{
		      {{ range .ValueInfo }}
		      &ir.ValueInfoProto{
			      Name: "{{ .Name }}",
			      Type: &ir.TypeProto{
				      Value: &ir.TypeProto_TensorType{
					      TensorType: &ir.TypeProto_Tensor{
						      ElemType: {{ .ElemType }},
						      Shape: &ir.TensorShapeProto{
							      Dim: []*ir.TensorShapeProto_Dimension{
								      {{ range .Dims }}
								      &ir.TensorShapeProto_Dimension{
									      Value: &ir.TensorShapeProto_Dimension_DimValue{
									      DimValue: {{ . }},
									      },
								      },
								      {{ end }} 
							      },
						      },
					      },
				      },
			      },
		      },
		      {{ end }}
	      },
	      {{ end }}
      },
} {{ end }}
`

type testCases struct {
	OpType      string
	Title       string
	Constructor string
}
