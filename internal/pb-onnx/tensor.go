package pb

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"

	"github.com/pkg/errors"

	"gorgonia.org/tensor"
)

// Tensor returns a Gorgonia compatible tensor
func (tx *TensorProto) Tensor() (tensor.Tensor, error) {
	if tx.Segment != nil {
		return nil, errors.Wrap(ErrNotYetImplemented, "This tensor is segmented")
	}
	// Get the data type
	dt, err := TensorProto_DataType(tx.DataType).Dtype()
	if err != nil {
		return nil, err
	}
	var size = make([]int, len(tx.Dims))
	for i := range tx.Dims {
		size[i] = int(tx.Dims[i])
	}
	opts := []tensor.ConsOpt{tensor.WithShape(size...), tensor.Of(dt)}
	switch dt {
	case tensor.Bool:
		switch {
		case tx.Int32Data != nil:
			backing := make([]bool, len(tx.Int32Data))
			for i := 0; i < len(tx.Int32Data); i++ {
				if tx.Int32Data[i] == 1 {
					backing[i] = true
				}
			}
			opts = append(opts, tensor.WithBacking(backing))
		default:
			return nil, errors.New("No data found")
		}

	case tensor.Float32:
		switch {
		case tx.RawData != nil:
			buf := bytes.NewReader(tx.RawData)
			element := make([]byte, 4)
			var err error
			var backing []float32
			for {
				var n int
				n, err = buf.Read(element)
				if err != nil || n != 4 {
					break
				}
				uintElement := binary.LittleEndian.Uint32(element)
				backing = append(backing, math.Float32frombits(uintElement))
			}
			if err != io.EOF {
				return nil, errors.Wrapf(err, "%v", ErrCorruptedData)
			}
			opts = append(opts, tensor.WithBacking(backing))
		case tx.FloatData != nil:
			opts = append(opts, tensor.WithBacking(tx.FloatData))
		default:
			return nil, errors.New("No data found")
		}
	case tensor.Float64:
		switch {
		case tx.DoubleData != nil:
			opts = append(opts, tensor.WithBacking(tx.DoubleData))
		case tx.RawData != nil:
			buf := bytes.NewReader(tx.RawData)
			element := make([]byte, 8)
			var err error
			var backing []float64
			for {
				var n int
				n, err = buf.Read(element)
				if err != nil || n != 8 {
					break
				}
				uintElement := binary.LittleEndian.Uint64(element)
				backing = append(backing, math.Float64frombits(uintElement))
			}
			if err != io.EOF {
				return nil, errors.Wrapf(err, "%v", ErrCorruptedData)
			}
			opts = append(opts, tensor.WithBacking(backing))
		default:
			return nil, errors.New("No data found")
		}
	case tensor.Int64:
		switch {
		case tx.RawData != nil:
			buf := bytes.NewReader(tx.RawData)
			element := make([]byte, 8)
			var err error
			var backing []int64
			for {
				var n int
				n, err = buf.Read(element)
				if err != nil || n != 8 {
					break
				}
				uintElement := binary.LittleEndian.Uint64(element)
				backing = append(backing, int64(uintElement))
			}

			if err != io.EOF {
				return nil, errors.Wrapf(err, "%v", ErrCorruptedData)
			}
			opts = append(opts, tensor.WithBacking(backing))
		case tx.Int64Data != nil:
			opts = append(opts, tensor.WithBacking(tx.Int64Data))
		default:
			return nil, errors.New("No data found")
		}
	case tensor.Int32:
		switch {
		case tx.RawData != nil:
			buf := bytes.NewReader(tx.RawData)
			element := make([]byte, 4)
			var err error
			var backing []int32
			for {
				var n int
				n, err = buf.Read(element)
				if err != nil || n != 4 {
					break
				}
				uintElement := binary.LittleEndian.Uint32(element)
				backing = append(backing, int32(uintElement))
			}
			if err != io.EOF {
				return nil, errors.Wrapf(err, "%v", ErrCorruptedData)
			}
			opts = append(opts, tensor.WithBacking(backing))
		case tx.Int32Data != nil:
			opts = append(opts, tensor.WithBacking(tx.Int32Data))
		default:
			return nil, errors.New("No data found")
		}
	default:
		return nil, errors.Wrapf(ErrNotYetImplemented, "Unknown type %v", dt)

	}

	return tensor.New(opts...), nil
}
