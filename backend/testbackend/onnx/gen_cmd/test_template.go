package main

import "text/template"

var testTemplate = template.Must(template.New("testCase").Parse(testTmpl))
var testCasesTemplate = template.Must(template.New("testCase").Parse(testCasesTmpl))

//var modelTemplate = template.Must(template.New("modelCase").Parse(modelTmpl))

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
&pb.NodeProto{
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
&pb.ModelProto{
      IrVersion: {{ .IrVersion }},
      OpsetImport: []*pb.OperatorSetIdProto{
	      &pb.OperatorSetIdProto{
		      Domain:  "",
		      Version: {{ .OpsetVersion }},
	      },
      },
      Graph: &pb.GraphProto{
	      Node: []*pb.NodeProto{
		      {{ range .NodeProto }}
		      &pb.NodeProto{
			      Input:     {{ .Input }},
			      Output:    {{ .Output }},
			      Name:      {{ .Name }},
			      OpType:    {{ .OpType }},
			      {{ if .Attributes }}
			      Attribute: []*pb.AttributeProto{
			      {{ range .Attributes }}
				  &pb.AttributeProto{
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
	      Initializer: []*pb.TensorProto{
		      /*
		      &pb.TensorProto{
			      Dims:      []int64{16, 4, 4, 10},
			      DataType:  1,
			      FloatData: parameter193,
			      Name:      "Parameter193",
		      },
		      */
	      },  
	      Input: []*pb.ValueInfoProto{
		      {{ range .Input }}
		      &pb.ValueInfoProto{
			      Name: "{{ .Name }}",
			      Type: &pb.TypeProto{
				      Value: &pb.TypeProto_TensorType{
					      TensorType: &pb.TypeProto_Tensor{
						      ElemType: {{ .ElemType }},
						      Shape: &pb.TensorShapeProto{
							      Dim: []*pb.TensorShapeProto_Dimension{
								      {{ range .Dims }}
								      &pb.TensorShapeProto_Dimension{
									      Value: &pb.TensorShapeProto_Dimension_DimValue{
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
	      Output: []*pb.ValueInfoProto{
		      {{ range .Output }}
		      &pb.ValueInfoProto{
			      Name: "{{ .Name }}",
			      Type: &pb.TypeProto{
				      Value: &pb.TypeProto_TensorType{
					      TensorType: &pb.TypeProto_Tensor{
						      ElemType: {{ .ElemType }},
						      Shape: &pb.TensorShapeProto{
							      Dim: []*pb.TensorShapeProto_Dimension{
								      {{ range .Dims }}
								      &pb.TensorShapeProto_Dimension{
									      Value: &pb.TensorShapeProto_Dimension_DimValue{
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
	      ValueInfo: []*pb.ValueInfoProto{
		      {{ range .ValueInfo }}
		      &pb.ValueInfoProto{
			      Name: "{{ .Name }}",
			      Type: &pb.TypeProto{
				      Value: &pb.TypeProto_TensorType{
					      TensorType: &pb.TypeProto_Tensor{
						      ElemType: {{ .ElemType }},
						      Shape: &pb.TensorShapeProto{
							      Dim: []*pb.TensorShapeProto_Dimension{
								      {{ range .Dims }}
								      &pb.TensorShapeProto_Dimension{
									      Value: &pb.TensorShapeProto_Dimension_DimValue{
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

const testCasesTmpl = `
package onnxtest

import "github.com/owulveryck/onnx-go/backend/testbackend"

func init() {
// Register all the test cases
{{ if . }} {{ range . }}
{{ if .OpType }} testbackend.Register({{ .OpType }},"{{ .Title }}",New{{ .Title }}) {{end}} {{ end }} {{end}}
}
`
