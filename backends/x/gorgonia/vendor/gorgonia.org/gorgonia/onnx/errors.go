package onnx

// ErrNotImplemented is fired when trying to call an operator that is not yet implemented
type ErrNotImplemented struct {
	Operator       string
	AttributeName  string
	AttributeValue interface{}
	Message        string
}

func (e *ErrNotImplemented) Error() string {
	return "operator" + e.Operator + " not implemented"
}
