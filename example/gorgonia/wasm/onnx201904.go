// +build wasm

package main

import (
	"bytes"
	"errors"
	"image"
	"math"
	"runtime"
	"syscall/js"
	"time"

	"github.com/disintegration/imaging"
	"github.com/owulveryck/onnx-go"
	"github.com/vincent-petithory/dataurl"
	"gorgonia.org/gorgonia/node"
	gorgonnx "gorgonia.org/gorgonia/onnx"
	"gorgonia.org/tensor"
)

const size = 28

var (
	graph *gorgonnx.Graph
	model *onnx.Model
)

func logInfo(s interface{}) {
	js.Global().Get("document").
		Call("getElementById", "info").
		Set("innerHTML", s)
}

func loadFile() error {
	graph = gorgonnx.NewGraph()
	model = onnx.NewModel(graph)
	files := js.Global().Get("document").Call("getElementById", "knowledgeFile").Get("files")
	//fmt.Println("file", files)
	//fmt.Println("Length", files.Length())
	if files.Length() == 1 {
		logInfo("Reading the model from the memory of the browser")
		//fmt.Println("Reading from uploaded file")
		reader := js.Global().Get("FileReader").New()
		reader.Call("readAsDataURL", files.Index(0))
		for reader.Get("readyState").Int() != 2 {
			//fmt.Println("Waiting for the file to be uploaded")
			time.Sleep(1 * time.Second)
		}
		content := reader.Get("result").String()
		dataURL, err := dataurl.DecodeString(content)
		if err != nil {
			logInfo(err.Error())
			return err
		}
		err = model.Unmarshal(dataURL.Data)
		if err != nil {
			logInfo("Fatal... cannot decode model, please reload the page")
			return err
		}
		// modelonnx = dataURL.Data
	}
	return nil
}

func main() {
	logInfo("Starting the WASM program")
	err := loadFile()
	if err != nil {
		logInfo(err.Error())
		return
	}
	runtime.GC()
	logInfo("Ready to process!")
	// Declare callback
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// handle event
		// Get the picture
		logInfo("processing element")
		t, err := processImgToArray()
		runtime.GC()
		if err != nil {
			logInfo(err.Error())
			return nil

		}
		output, err := analyze(t)
		runtime.GC()
		if err != nil {
			logInfo(err.Error())
			return nil
		}

		logInfo(output)
		return nil
	})
	// Hook it up with a DOM event
	js.Global().Get("document").
		Call("getElementById", "btnSubmit").
		Call("addEventListener", "click", cb)
	c := make(chan struct{}, 0)
	<-c
}

func processImgToArray() ([]float32, error) {
	pic := js.Global().Get("document").Call("getElementById", "canvasBox").Call("toDataURL")
	dataURL, err := dataurl.DecodeString(pic.String())
	if err != nil {
		return nil, err
	}
	if dataURL.ContentType() != "image/png" {
		return nil, errors.New("not a png image")
	}
	m, _, err := image.Decode(bytes.NewReader(dataURL.Data))
	if err != nil {
		return nil, err
	}

	img := imaging.Resize(m, size, 0, imaging.Lanczos)
	t := make([]float32, size*size)
	for i := 0; i < size*size*4; i += 4 {
		t[i/4] = float32(img.Pix[i])
	}
	return t, nil
}

func analyze(t []float32) (int, error) {
	T := tensor.New(tensor.WithBacking(t), tensor.WithShape(1, 1, size, size))
	err := gorgonnx.Let(graph.Node(model.Input[0]).(node.Node), T)
	if err != nil {
		return 0, err
	}
	// create a VM to run the program on
	machine := gorgonnx.NewTapeMachine(graph)

	// Run the program
	err = machine.RunAll()
	if err != nil {
		return 0, nil
	}
	val := float32(-math.MaxFloat32)
	res := 0
	for i, v := range graph.Node(model.Output[0]).(node.Node).Value().Data().([]float32) {
		if v > val {
			res = i
			val = v
		}
	}
	return res, nil
}
