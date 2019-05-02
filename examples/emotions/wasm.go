// +build js,wasm

package main

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"runtime"
	"time"

	"syscall/js"

	"github.com/vincent-petithory/dataurl"
)

func logInfo(s interface{}) {
	js.Global().Get("document").
		Call("getElementById", infoBoxElementID).
		Set("innerHTML", s)
}

func getModel() ([]byte, error) {
	files := js.Global().Get("document").Call("getElementById", knowledgeFileElementID).Get("files")
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
		return dataURL.Data, err
	}
	return nil, errors.New("too many file in the selector")
}

func getImage() (*image.Gray, error) {
	pic := js.Global().Get("document").Call("getElementById", canvasElementID).Call("toDataURL")
	dataURL, err := dataurl.DecodeString(pic.String())
	if err != nil {
		return nil, err
	}
	if dataURL.ContentType() != "image/png" {
		return nil, errors.New("not a png image")
	}
	m, err := png.Decode(bytes.NewReader(dataURL.Data))
	if err != nil {
		return nil, err
	}

	imgGray, ok := m.(*image.Gray)
	if !ok {
		return nil, errors.New("Not a gray image")
	}
	return imgGray, nil
}

func displayResult(e emotions) {
}

func run() error {
	b, err := getModel()
	if err != nil {
		logInfo(err)
		return err
	}
	err = model.UnmarshalBinary(b)
	if err != nil {
		logInfo(err)
		return err
	}
	// Declare callback
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// handle event
		// Get the picture
		logInfo("processing element")
		img, err := getImage()
		runtime.GC()
		if err != nil {
			logInfo(err.Error())
			return nil

		}
		output, err := process(img)
		runtime.GC()
		if err != nil {
			logInfo(err.Error())
			return nil
		}

		displayResult(output)
		return nil
	})
	// Hook it up with a DOM event
	js.Global().Get("document").
		Call("getElementById", boutonSubmit).
		Call("addEventListener", "click", cb)
	c := make(chan struct{}, 0)
	<-c
	return nil
}
