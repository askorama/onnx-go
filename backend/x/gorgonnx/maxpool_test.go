package gorgonnx

import (
	"testing"

	onnx "github.com/owulveryck/onnx-go"
)

//https://stackoverflow.com/questions/37674306/what-is-the-difference-between-same-and-valid-padding-in-tf-nn-max-pool-of-t
//
func TestMaxpool_SAME(t *testing.T) {
	maxpoolOp := &maxpool{}
	operation := onnx.Operation{
		Attributes: map[string]interface{}{
			"auto_pad":     "SAME_UPPER",
			"kernel_shape": []int64{2, 2},
			"strides":      []int64{2, 2},
		},
	}
	err := maxpoolOp.init(operation)
	if err != nil {
		t.Fatal(err)
	}
	maxpoolOp.computePadding([]int{1, 1, 2, 3})
	if maxpoolOp.pad[0] != 0 ||
		maxpoolOp.pad[1] != 0 ||
		maxpoolOp.pad[2] != 1 ||
		maxpoolOp.pad[3] != 0 {
		t.Fail()
	}

}
func TestMaxpool_VALID(t *testing.T) {
	maxpoolOp := &maxpool{}
	operation := onnx.Operation{
		Attributes: map[string]interface{}{
			"auto_pad":     "VALID",
			"kernel_shape": []int64{2, 2},
			"strides":      []int64{2, 2},
		},
	}
	err := maxpoolOp.init(operation)
	if err != nil {
		t.Fatal(err)
	}
	maxpoolOp.computePadding([]int{1, 1, 2, 3})
	if maxpoolOp.pad[0] != 0 ||
		maxpoolOp.pad[1] != 0 ||
		maxpoolOp.pad[2] != 0 ||
		maxpoolOp.pad[3] != 0 {
		t.Fail()
	}
}
