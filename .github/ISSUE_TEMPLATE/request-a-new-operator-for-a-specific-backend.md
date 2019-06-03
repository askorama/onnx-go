---
name: Request a new operator for a specific backend
about: Request a new operator for a specific backend
title: Implement operator `XXX` for backend `YYY`
labels: Feature request, Operator, backend
assignees: ''

---

## Why is this operator needed?

<-- provide some context, eg: this operator is needed to run the model `xxx` with the backend `yyy` -->

## Implementation

* [Description of the `XXX` operator in ONNX](https://github.com/onnx/onnx/blob/master/docs/Operators.md#)
<-- You can also give a pointer to a paper if needed -->

### Link to existing material on the backend

<--
* [Godoc of the operator `XXX` on Gorgonia](https://godoc.org/gorgonia.org/gorgonia#XXX)
-->

#### Expected problems?

<-- Please give more details such as: The operator should be broadcastable -->
N/A

### Tests

<-- Give any information to test the operator 
`go test -run=ONNX/TestOperator`
-->
