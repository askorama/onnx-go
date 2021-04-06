package ir

import (
	"encoding/binary"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTensor_float32(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["FLOAT"])
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

	// Corrupted Data
	txFloat32.RawData = rawData[1:] // remove first slide item to corrupt data
	_, err = txFloat32.Tensor()
	assert.EqualError(t, err, "<nil>: Unable to decode data")
}

func TestNewTensor_float32_noData(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["FLOAT"])
	name := "testFloat"
	txFloat32 := &TensorProto{
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
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txFloat32.Tensor()
	assert.EqualError(t, err, "No data found")
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
	txFloat64.DoubleData = nil
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

	// Corrupted Data
	txFloat64.RawData = rawData[1:] // remove first slide item to corrupt data
	_, err = txFloat64.Tensor()
	assert.EqualError(t, err, "<nil>: Unable to decode data")
}

func TestNewTensor_float64_noData(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["DOUBLE"])
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
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txFloat64.Tensor()
	assert.EqualError(t, err, "No data found")
}

func TestNewTensor_bool(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["BOOL"])
	boolData := []int32{1, 1, 1, 0, 0, 0, 1, 1}
	name := "testBool"
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
	var rawData []byte
	for _, v := range boolData {
		b := make([]byte, 8)
		b[7] = byte(v)
		rawData = append(rawData, b...)
	}
	txBool.Int32Data = nil
	txBool.RawData = rawData
	tg, err = txBool.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]bool) {
		if v && boolData[i] == 0 {
			t.Fail()
		}
	}

	// Corrupted Data
	txBool.RawData = rawData[1:] // remove first slide item to corrupt data
	_, err = txBool.Tensor()
	assert.EqualError(t, err, "<nil>: Unable to decode data")
}

func TestNewTensor_bool_noData(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["BOOL"])
	name := "testBool"
	txBool := &TensorProto{
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
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txBool.Tensor()
	assert.EqualError(t, err, "No data found")
}

func TestNewTensor_UnknownType(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := 999
	name := "testUnknownType"
	txBool := &TensorProto{
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
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txBool.Tensor()
	assert.EqualError(t, err, "Unknown input type: 999")
}

func TestNewTensor_UndefinedType(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["UNDEFINED"])
	name := "testBool"
	txBool := &TensorProto{
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
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txBool.Tensor()
	assert.EqualError(t, err, "This tensor datatype is undefined")
}

func TestNewTensor_NotYetImplementedType(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataTypeStrings := []string{
		"FLOAT16",
	}
	name := "testNotYetImplementedType"
	for _, dataTypeString := range dataTypeStrings {
		dataType := TensorProto_DataType(TensorProto_DataType_value[dataTypeString])
		txNoHandledType := &TensorProto{
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
			DoubleData: nil,
			Uint64Data: nil,
		}
		_, err := txNoHandledType.Tensor()
		assert.EqualError(t, err, "type: "+dataTypeString+": Not Yet Implemented")
	}
}

func TestNewTensor_UnknownNotYetImplementedType(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataTypeStrings := []string{
		"UINT8",
		"INT8",
		"UINT16",
		"INT16",
		"STRING",
		"UINT32",
		"UINT64",
		"COMPLEX64",
		"COMPLEX128",
	}
	name := "testUnknownNotYetImplementedType"
	for _, dataTypeString := range dataTypeStrings {
		dataType := TensorProto_DataType(TensorProto_DataType_value[dataTypeString])
		txNoHandledType := &TensorProto{
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
			DoubleData: nil,
			Uint64Data: nil,
		}
		_, err := txNoHandledType.Tensor()
		assert.EqualError(t, err, "Unknown type "+strings.ToLower(dataTypeString)+": Not Yet Implemented")
	}
}

func TestNewTensor_withSegment(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["BOOL"])
	name := "testBool"
	txBool := &TensorProto{
		Dims:       dims,
		DataType:   int32(dataType),
		Segment:    &TensorProto_Segment{},
		FloatData:  nil,
		Int32Data:  nil,
		StringData: nil,
		Int64Data:  nil,
		Name:       name,
		DocString:  "",
		RawData:    nil,
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txBool.Tensor()
	assert.EqualError(t, err, "This tensor is segmented: Not Yet Implemented")
}

func TestNewTensor_int32(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["INT32"])
	intData := []int32{-1, -1, 1, -1, -1, -1, 1, -1}
	name := "testInt"
	txInt32 := &TensorProto{
		Dims:       dims,
		DataType:   int32(dataType),
		Segment:    (*TensorProto_Segment)(nil),
		FloatData:  nil,
		Int32Data:  intData,
		StringData: nil,
		Int64Data:  nil,
		Name:       name,
		DocString:  "",
		RawData:    nil,
		DoubleData: nil,
		Uint64Data: nil,
	}
	tg, err := txInt32.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]int32) {
		if v != intData[i] {
			t.Fail()
		}
	}
	var rawData []byte
	for _, v := range intData {
		b := make([]byte, 4)
		uintElement := uint32(v)
		binary.LittleEndian.PutUint32(b, uintElement)
		rawData = append(rawData, b...)
	}
	txInt32.Int32Data = nil
	txInt32.RawData = rawData
	tg, err = txInt32.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]int32) {
		if v != intData[i] {
			t.Fail()
		}
	}

	// Corrupted Data
	txInt32.RawData = rawData[1:] // remove first slide item to corrupt data
	_, err = txInt32.Tensor()
	assert.EqualError(t, err, "<nil>: Unable to decode data")
}

func TestNewTensor_int32_noData(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["INT32"])
	name := "testInt"
	txInt32 := &TensorProto{
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
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txInt32.Tensor()
	assert.EqualError(t, err, "No data found")
}

func TestNewTensor_int64(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["INT64"])
	intData := []int64{-1, -1, 1, -1, -1, -1, 1, -1}
	name := "testInt"
	txInt64 := &TensorProto{
		Dims:       dims,
		DataType:   int32(dataType),
		Segment:    (*TensorProto_Segment)(nil),
		FloatData:  nil,
		Int32Data:  nil,
		StringData: nil,
		Int64Data:  intData,
		Name:       name,
		DocString:  "",
		RawData:    nil,
		DoubleData: nil,
		Uint64Data: nil,
	}
	tg, err := txInt64.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]int64) {
		if v != intData[i] {
			t.Fail()
		}
	}
	var rawData []byte
	for _, v := range intData {
		b := make([]byte, 8)
		uintElement := uint64(v)
		binary.LittleEndian.PutUint64(b, uintElement)
		rawData = append(rawData, b...)
	}
	txInt64.Int64Data = nil
	txInt64.RawData = rawData
	tg, err = txInt64.Tensor()
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range tg.Data().([]int64) {
		if v != intData[i] {
			t.Fail()
		}
	}

	// Corrupted Data
	txInt64.RawData = rawData[1:] // remove first slide item to corrupt data
	_, err = txInt64.Tensor()
	assert.EqualError(t, err, "<nil>: Unable to decode data")
}

func TestNewTensor_int64_noData(t *testing.T) {
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(TensorProto_DataType_value["INT64"])
	name := "testInt"
	txInt64 := &TensorProto{
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
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txInt64.Tensor()
	assert.EqualError(t, err, "No data found")
}

func TestNewTensor_float32_raw_empty(t *testing.T) {
	dims := []int64{0}
	dataType := TensorProto_DataType(TensorProto_DataType_value["FLOAT"])
	rawData := []byte{}
	name := "testFloat"
	txFloat32 := &TensorProto{
		Dims:       dims,
		DataType:   int32(dataType),
		Segment:    (*TensorProto_Segment)(nil),
		FloatData:  nil,
		Int32Data:  nil,
		StringData: nil,
		Int64Data:  nil,
		Name:       name,
		DocString:  "",
		RawData:    rawData,
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txFloat32.Tensor()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewTensor_float64_raw_empty(t *testing.T) {
	dims := []int64{0}
	dataType := TensorProto_DataType(TensorProto_DataType_value["DOUBLE"])
	rawData := []byte{}
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
		RawData:    rawData,
		DoubleData: nil,
		Uint64Data: nil,
	}
	_, err := txFloat64.Tensor()
	if err != nil {
		t.Fatal(err)
	}
}
