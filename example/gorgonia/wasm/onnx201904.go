// +build wasm

package main

import (
	"bytes"
	"image"
	"log"
	"math"
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

func main() {
	graph = gorgonnx.NewGraph()
	model = onnx.NewModel(graph)
	files := js.Global().Get("document").Call("getElementById", "knowledgeFile").Get("files")
	//fmt.Println("file", files)
	//fmt.Println("Length", files.Length())
	if files.Length() == 1 {
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
			log.Println(err)
			return
		}
		err = model.Unmarshal(dataURL.Data)
		if err != nil {
			log.Fatal(err)
		}

		// modelonnx = dataURL.Data
	}
	// Declare callback
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// handle event
		// Get the picture
		pic := js.Global().Get("document").Call("getElementById", "canvasBox").Call("toDataURL")
		dataURL, err := dataurl.DecodeString(pic.String())
		if err != nil {
			log.Println(err)
			return nil
		}
		var output int
		if dataURL.ContentType() == "image/png" {
			m, _, err := image.Decode(bytes.NewReader(dataURL.Data))
			if err != nil {
				log.Println(err)
				return nil
			}
			output, err = analyze(m)
			if err != nil {
				log.Println(err)
				return nil
			}
		}

		js.Global().Get("document").
			Call("getElementById", "guess").
			Set("innerHTML", output)
		return nil
	})
	// Hook it up with a DOM event
	js.Global().Get("document").
		Call("getElementById", "btnSubmit").
		Call("addEventListener", "click", cb)
	c := make(chan struct{}, 0)
	<-c
}

func analyze(m image.Image) (int, error) {
	// resize the image
	img := imaging.Resize(m, size, 0, imaging.Lanczos)
	t := make([]float32, size*size)
	for i := 0; i < size*size*4; i += 4 {
		t[i/4] = float32(img.Pix[i])
	}
	T := tensor.New(tensor.WithBacking(t), tensor.WithShape(1, 1, size, size))
	err := gorgonnx.Let(graph.Node(model.Input[0]).(node.Node), T)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// create a VM to run the program on
	machine := gorgonnx.NewTapeMachine(graph)

	// Run the program
	err = machine.RunAll()
	if err != nil {
		log.Println(err)
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
