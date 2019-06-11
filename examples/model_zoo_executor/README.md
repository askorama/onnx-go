# About

This is a simple utility that runs a model from the model zoo thanks to the Gorgonia backend

## Example

Download a pre-trained [model from the zoo](https://github.com/onnx/models) (for now, only [MNIST](https://github.com/onnx/models/tree/master/mnist) is known to work)

then smply run:

`go run main.go -model /tmp/mnist/model.onnx -input /tmp/mnist/test_data_set_0/input_0.pb -output /tmp/mnist/test_data_set_0/output_0.pb`

The utility evaluates the model and check if the computed output is equal to the expected output (within a delta of 5e-3).
If the result is ok, it displays the result:

`[975.67035 -618.7244 6574.5684 668.0278 -917.27057 -1671.6357 -1952.7606 -61.54949 -777.17645 -1439.5311]`

# Test files

You can test (and bench a model) with `go test`.
Simply download a model from the zoo into a subdirectory of the `./models` dir; then run:

`MODELDIR=./models/YOURMODEL go test`

If you have several models inside the `./models` dir, you can run them all by typing `make`

## Bench

A simple benchmark that measures the execution time of the model is also included; to run it:

`MODELDIR=./models/YOURMODEL go test -bench=.`

### Profiling

You can also use this benchmark for profiling the application:

`MODELDIR=models/YOURMODEL go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out -benchtime=10s`

The use the `pprof tool`:

`go tool pprof -http=localhost:3999 memprofile.out`
or 
`go tool pprof -http=localhost:3999 profile.out`

