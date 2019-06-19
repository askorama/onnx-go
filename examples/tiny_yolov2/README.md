# This is a sample utility that runs tiny-yolo v2

to run this utility you need:

- the onnx definition of tiny yolo v2 from the [model zoo](https://github.com/onnx/models/tree/master/tiny_yolov2).
- a jpeg picture

## Usage

```
$ go run main.go -h
  -h    help
  -img string
        path of an input jpeg image (use - for stdin)
  -model string
        path to the model file (default "model.onnx")
  -output string
        path of an output png file (use - for stdout) (default "output.png")
  -s    silent mode (useful if output is -)
This application is configured via the environment. The following environment
variables can be used:

KEY                          TYPE     DEFAULT    REQUIRED    DESCRIPTION
YOLO_CONFIDENCE_THRESHOLD    Float    0.30       true
YOLO_PROBA_THRESHOLD         Float    0.90       true
```

to run it, simply do a 
`go run main.go -model /path/to/tiny_yolov2/model.onnx -img /path/to/picture.jpg`

if you want to generate an anotated output picture in png, use the `-output` parameter


You can alter the output by playing with the environment variables

- `YOLO_CONFIDENCE_THRESHOLD`: bypass the boxes with a confidence lower than this value
- `YOLO_PROBA_THRESHOLD`: bypass the boxes with a class detection lower than this value
