package pb

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
		DataType:   int32(dataType),
		Segment:    (*TensorProto_Segment)(nil),
		FloatData:  floatData,
		Int32Data:  nil,
		StringData: nil,
		Int64Data:  nil,
		Name:       name,
		DocString:  "",
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
func TestNewTensor_float64(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["DOUBLE"])
	floatData := []float64{-0.1615397185087204, -0.4338356554508209, 0.09164135903120041, -0.01685221679508686, -0.06502643972635269, -0.1317378729581833, 0.020417550578713417, -0.1211102306842804}
	name := "testFloat"
	txFloat64 := &TensorProto{
		Dims:       dims,
		DataType:   int32(dataType),
		Segment:    (*TensorProto_Segment)(nil),
		FloatData:  nil,
		Int32Data:  nil,
		StringData: nil,
		Int64Data:  nil,
		Name:       name,
		DocString:  "",
		RawData:    nil,
		DoubleData: floatData,
		Uint64Data: nil,
	}
	tg, err := txFloat64.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]float64) {
		if v != floatData[i] {
			t.Fail()
		}
	}
	var rawData []byte
	for _, v := range floatData {
		b := make([]byte, 8)
		uintElement := math.Float64bits(v)
		binary.LittleEndian.PutUint64(b, uintElement)
		rawData = append(rawData, b...)
	}
	txFloat64.FloatData = nil
	txFloat64.RawData = rawData
	tg, err = txFloat64.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]float64) {
		if v != floatData[i] {
			t.Fail()
		}
	}
}
func TestNewTensor_bool(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["BOOL"])
	boolData := []int32{1, 1, 1, 0, 0, 0, 1, 1}
	name := "testFloat"
	txBool := &TensorProto{
		Dims:       dims,
		DataType:   int32(dataType),
		Segment:    (*TensorProto_Segment)(nil),
		FloatData:  nil,
		Int32Data:  boolData,
		StringData: nil,
		Int64Data:  nil,
		Name:       name,
		DocString:  "",
		RawData:    nil,
		DoubleData: nil,
		Uint64Data: nil,
	}
	tg, err := txBool.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]bool) {
		if v && boolData[i] == 0 {
			t.Fail()
		}
	}
}
