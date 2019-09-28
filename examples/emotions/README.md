## Emotion fer+

This is an example of onnx go being able to run the emotion fer+ model

### Howto

- Download an extract the model from [the zoo](https://github.com/onnx/models/tree/master/vision/body_analysis/emotion_ferplus)
- take a picture of a face: the picture should be 64X64 in gray mode and in png.
- run `go run main.go -model /path/to/model.onnx -input mypic.png`

### pre-processing with ImageMagick

The tool assumes that the input image is in the correct format.
You can check the format with this command:

```
> identify input.png
me.png PNG 64x64 64x64+0+0 8-bit LinearGray 256c 2665B 0.010u 0:00.000
```

You can convert a picture to the expected format with:

```
> convert source.jpg -resize 64 -set colorspace Gray -separate -average dest.png
```

You can also pipe the result directly into the command:

```
convert ~/Downloads/download.png -resize 64 -set colorspace Gray -separate -average png:- | \
    go run main.go -model /path/to/model.onnx -input -
```
