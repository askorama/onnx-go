package onnx

import (
	"testing"
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
		I    int64   `attribute:"I" required:"true"`
		Ints []int64 `attribute:"INTS" required:"true"`
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
	simpleAttributes := []AttributeProto{
		AttributeProto{
			Name: &simpleAttributeIName,
			Type: &simpleAttributeIType,
			I:    &simpleAttributeIValue,
		},
		AttributeProto{
			Name: &simpleAttributeINTSName,
			Type: &simpleAttributeINTSType,
			Ints: simpleAttributeINTSValue,
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
}
