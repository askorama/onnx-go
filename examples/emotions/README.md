## Emotion fer+

This is an example of onnx go being able to run the emotion fer+ model

### Howto

- Download an extract the model from [the zoo](https://github.com/onnx/models/tree/master/emotion_ferplus)
- take a picture of a face: the picture should be 64X64 in gray mode and in png.
- run `go build && ./emotions -model /path/to/model.onnx -input mypic.png`

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
    ./emotions -model /path/to/model.onnx -input -
```

On MacOS good results are obtained with this configuration:

```
> imagesnap # brew install imagesnap
> convert snapshot.jpg -gravity center -crop 300x300+0+0 -resize 64x64 -set colorspace Gray -separate -average -gamma 0.8 png:-  | tee output.png | ./emotions -model model.onnx -input -
```

```
> open snapshot.jpg
> open output.png
```

### Wasm experiment

This utility can be compiled in wasm, but it is highly experimental and may fail in out-of-memory very often...
