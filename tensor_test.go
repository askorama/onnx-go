package onnx

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestNewTensor_float32(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(1)
	floatData := []float32{-0.1615397185087204, -0.4338356554508209, 0.09164135903120041, -0.01685221679508686, -0.06502643972635269, -0.1317378729581833, 0.020417550578713417, -0.1211102306842804}
	name := "testFloat"
	txFloat32 := &TensorProto{
		Dims:       dims,
		DataType:   &dataType,
		Segment:    (*TensorProto_Segment)(nil),
		FloatData:  floatData,
		Int32Data:  nil,
		StringData: nil,
		Int64Data:  nil,
		Name:       &name,
		DocString:  nil,
		RawData:    nil,
		DoubleData: nil,
		Uint64Data: nil,
	}
	tg, err := txFloat32.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]float32) {
		if v != floatData[i] {
			t.Fail()
		}
	}
	var rawData []byte
	for _, v := range floatData {
		b := make([]byte, 4)
		uintElement := math.Float32bits(v)
		binary.LittleEndian.PutUint32(b, uintElement)
		rawData = append(rawData, b...)
	}
	txFloat32.FloatData = nil
	txFloat32.RawData = rawData
	tg, err = txFloat32.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]float32) {
		if v != floatData[i] {
			t.Fail()
		}
	}
}
