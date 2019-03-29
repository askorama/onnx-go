![ONNX Logo](vignettes/imgs/ONNX_logo_main.png) ![Go Logo](vignettes/imgs/Go-Logo_Blue.png)

[![](https://godoc.org/github.com/owulveryck/onnx-go?status.svg)](http://godoc.org/github.com/owulveryck/onnx-go)

# onnx-go

This is a Go Interface to [Open Neural Network Exchange (ONNX)](https://onnx.ai/).

## Overview [![GoDoc](https://godoc.org/github.com/owulveryck/onnx-go?status.svg)](https://godoc.org/github.com/owulveryck/onnx-go) [![Go Report Card](https://goreportcard.com/badge/github.com/owulveryck/onnx-go)](https://goreportcard.com/report/github.com/owulveryck/onnx-go)

onnx-go contains primitives to decode a onnx binary model into a computation backend, and use it like any other library in your go code.
for more information about onnx, please visit [onnx.ai](https://onnx.ai).

## Install

```
go get github.com/owulveryck/onnx-go
```

## Example

[embedmd]:# (example_test.go /\/\/ Create/ /model.Decode.*/)
```go
// Create a backend receiver
	backend := simple.NewSimpleGraph()
	// Create a model and set the execution backend
	model := onnx.NewModel(backend)

	// read the onnx model
	b, _ := ioutil.ReadFile("model.onnx")
	// Decode it into the model
	err := model.Decode(b)
```


## Contributing

ToDo.

## Author

[Olivier Wulveryck](https://about.me/owulveryck/getstarted)

## License

MIT.

