## Emotion fer+

This is an example of onnx go being able to run the emotion fer+ model

### Howto

- Download an extract the model from [the zoo](https://github.com/onnx/models/tree/master/emotion_ferplus)
- take a picture of a face: the picture should be 64X64 in gray mode and in png.
- run `go run main.go -model /path/to/model.onnx -input mypic.png`




