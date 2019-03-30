package onnx

func testReadme() {
	convOperator := Operation{
		Name: "Conv",
		Attributes: map[string]interface{}{
			"auto_pad":  "NOTSET",
			"dilations": []int64{1, 1},
			"group":     1,
			"pads":      []int64{1, 1},
			"strides":   []int64{1, 1},
		},
	}
	_ = convOperator
}
