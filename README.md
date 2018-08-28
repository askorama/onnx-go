![ONNX Logo](vignettes/imgs/ONNX_logo_main.png) ![Go Logo](vignettes/imgs/Go-Logo_Blue.png)

[![](https://godoc.org/github.com/owulveryck/onnx-go?status.svg)](http://godoc.org/github.com/owulveryck/onnx-go)

This is the Go Interface to [Open Neural Network Exchange (ONNX)](https://onnx.ai/).

# About

This is a compiled version of the ONNX protobuf definition file.

## Using this API

This package is go-gettable.

```shell
go get -v github.com/owulveryck/onnx-go
```

### Example

This reads a ONNX file and decode the Model into the `ModelProto`

```go
import (
        "io/ioutil"
        "log"

        onnx "github.com/owulveryck/onnx-go"
)

func main() {
        b, err := ioutil.ReadFile("/path/to/onnx/file.onnx")
        if err != nil {
                log.Fatal(err)
        }
        model := new(onnx.ModelProto)
        err = model.Unmarshal(b)
        if err != nil {
                log.Fatal(err)
        }
        log.Println(model)
}
```

For more information, please refer to the GoDoc.

## Generating the file

This file is generated from the protobuf definition file of the official [Github repository of ONNX](github.com/onnx/onnx).
You should not need to rebuild it and it can be used as-is.

Anyway if for any reason you want to test some new feature, you need:

* a protobuf compiler with a go extension (see [here](https://github.com/golang/protobuf) for more info;
* the [Protocol Buffers for Go with Gadgets](https://github.com/gogo/protobuf) (aka gogoproto) binaries 

Checkout the repository of ONNX somewhere, then run:

```
protoc --gofast_out=. /onnx/onnx.proto
``` 

from the base directory of the repository.
