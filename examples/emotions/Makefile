all: wasm/emotions.wasm wasm/model.onnx

wasm/emotions.wasm: *.go
	GOOS=js GOARCH=wasm go build -o wasm/emotions.wasm

wasm/model.onnx:
	curl https://onnxzoo.blob.core.windows.net/models/opset_8/emotion_ferplus/emotion_ferplus.tar.gz | tar xzf - --strip=1 -C wasm emotion_ferplus/model.onnx

