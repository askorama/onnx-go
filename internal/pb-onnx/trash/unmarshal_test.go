package pb

import (
	"testing"

	"gorgonia.org/tensor"
)

func TestUnmarshalAttributes(t *testing.T) {
	err := UnmarshalAttributes(nil, nil)
	var receiver string
	if err == nil {
		t.Fail()
	} else {
		_, ok := err.(*InvalidUnmarshalError)
		if !ok {
			t.Fatal("Invalid error type", err)
		}
	}
	err = UnmarshalAttributes(nil, receiver)
	if err == nil {
		t.Fail()
	} else {
		_, ok := err.(*InvalidUnmarshalError)
		if !ok {
			t.Fatal("Invalid error type", err)
		}
	}
	err = UnmarshalAttributes(nil, &receiver)
	if err == nil {
		t.Fail()
	} else {
		_, ok := err.(*InvalidUnmarshalError)
		if !ok {
			t.Fatal("Invalid error type", err)
		}
	}
	type validReceiverType struct {
		I    int64         `attributeName:"I" required:"true"`
		Ints []int64       `attributeName:"INTS" required:"true"`
		T    tensor.Tensor `attributeName:"Tensor" required:"true"`
	}
	var validReceiver validReceiverType
	err = UnmarshalAttributes(nil, &validReceiver)
	if err == nil {
		t.Fail()
	}
	simpleAttributeIName := "I"
	simpleAttributeIType := AttributeProto_INT
	simpleAttributeIValue := int64(1)
	simpleAttributeINTSName := "INTS"
	simpleAttributeINTSType := AttributeProto_INTS
	simpleAttributeINTSValue := []int64{1, 2, 3, 4}
	simpleAttributeTensorName := "Tensor"
	simpleAttributeTensorType := AttributeProto_TENSOR
	dims := []int64{8, 1, 1}
	dataType := TensorProto_DataType(1)
	floatData := []float32{-0.1615397185087204, -0.4338356554508209, 0.09164135903120041, -0.01685221679508686, -0.06502643972635269, -0.1317378729581833, 0.020417550578713417, -0.1211102306842804}
	name := "testFloat"
	simpleAttributeTensorValue := &TensorProto{
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

	simpleAttributes := []*AttributeProto{
		&AttributeProto{
			Name: &simpleAttributeIName,
			Type: &simpleAttributeIType,
			I:    &simpleAttributeIValue,
		},
		&AttributeProto{
			Name: &simpleAttributeINTSName,
			Type: &simpleAttributeINTSType,
			Ints: simpleAttributeINTSValue,
		},
		&AttributeProto{
			Name: &simpleAttributeTensorName,
			Type: &simpleAttributeTensorType,
			T:    simpleAttributeTensorValue,
		},
	}
	err = UnmarshalAttributes(simpleAttributes, &validReceiver)
	if err != nil {
		t.Fatal(err)
	}
	if validReceiver.I != simpleAttributeIValue {
		t.Fail()
	}
	if len(simpleAttributeINTSValue) != len(validReceiver.Ints) {
		t.Fail()
	}
	for i := range simpleAttributeINTSValue {
		if simpleAttributeINTSValue[i] != validReceiver.Ints[i] {
			t.FailNow()
		}
	}
	t.Log(validReceiver)
}
