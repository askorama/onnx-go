package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Clip", "TestClipDefaultMax", NewTestClipDefaultMax)
}

// NewTestClipDefaultMax version: 3.
func NewTestClipDefaultMax() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Clip",
		Title:  "TestClipDefaultMax",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x66, 0xa, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x4, 0x43, 0x6c, 0x69, 0x70, 0x2a, 0xd, 0xa, 0x3, 0x6d, 0x61, 0x78, 0x15, 0x0, 0x0, 0x0, 0x0, 0xa0, 0x1, 0x1, 0x12, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x70, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Clip",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000133000)(name:"max" type:FLOAT )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{-0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732, 0.01050002, 1.7858706, 0.12691209, 0.40198937, 1.8831507, -1.347759, -1.270485, 0.9693967, -1.1731234, 1.9436212, -0.41361898, -0.7474548, 1.922942, 1.4805148, 1.867559, 0.90604466, -0.86122566, 1.9100649, -0.26800337, 0.8024564, 0.947252, -0.15501009, 0.61407936, 0.9222067}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{-0.67246044, -0.35955316, -0.8131463, -1.7262826, 0, -0.40178093, -1.6301984, 0, -0.9072984, 0, 0, 0, 0, -1.2348258, 0, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0, -1.1651498, 0, 0, -1.5362437, 0, 0, 0, -0.17992483, -1.0707526, 0, -0.40317693, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1.347759, -1.270485, 0, -1.1731234, 0, -0.41361898, -0.7474548, 0, 0, 0, 0, -0.86122566, 0, -0.26800337, 0, 0, -0.15501009, 0, 0}),
			),
		},
	}
}
