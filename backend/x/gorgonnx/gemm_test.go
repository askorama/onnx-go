package gorgonnx

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/stretchr/testify/assert"
	"gorgonia.org/tensor"
)

var testGemm = []struct {
	name            string
	aT              tensor.Tensor
	bT              tensor.Tensor
	cT              tensor.Tensor
	expectedOutputT tensor.Tensor
}{
	{
		name: "float32-broadcast",
		aT: tensor.New(
			tensor.WithShape(6, 3),
			tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941,
				0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949,
				0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985}),
		),
		bT: tensor.New(
			tensor.WithShape(4, 6),
			tensor.WithBacking([]float32{0.77815676, 0.87001216, 0.9786183, 0.7991586,
				0.46147937, 0.7805292, 0.11827443, 0.639921,
				0.14335328, 0.9446689, 0.5218483, 0.41466194,
				0.2645556, 0.7742337, 0.45615032, 0.56843394,
				0.0187898, 0.6176355, 0.6120957, 0.616934,
				0.94374806, 0.6818203, 0.3595079, 0.43703195}),
		),
		cT: tensor.New(
			tensor.WithShape(1),
			tensor.WithBacking([]float32{0.6976312}),
		),
		expectedOutputT: tensor.New(
			tensor.WithShape(3, 4),
			tensor.WithBacking([]float32{1.3317792, 0.9343705, 0.8733721,
				1.1432098, 1.7855448, 1.2102433,
				1.0507759, 1.5598905, 1.8885028,
				1.1011722, 1.3064879, 1.5622698}),
		),
	},
	{
		name: "float64-broadcast2",
		aT: tensor.New(
			tensor.WithShape(6, 3),
			tensor.WithBacking([]float64{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941,
				0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949,
				0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985}),
		),
		bT: tensor.New(
			tensor.WithShape(4, 6),
			tensor.WithBacking([]float64{0.77815676, 0.87001216, 0.9786183, 0.7991586,
				0.46147937, 0.7805292, 0.11827443, 0.639921,
				0.14335328, 0.9446689, 0.5218483, 0.41466194,
				0.2645556, 0.7742337, 0.45615032, 0.56843394,
				0.0187898, 0.6176355, 0.6120957, 0.616934,
				0.94374806, 0.6818203, 0.3595079, 0.43703195}),
		),
		cT: tensor.New(
			tensor.WithShape(4),
			tensor.WithBacking([]float64{0.6976312, 0.6976312, 0.6976312, 0.6976312}),
		),
		expectedOutputT: tensor.New(
			tensor.WithShape(3, 4),
			tensor.WithBacking([]float64{1.3317792, 0.9343705, 0.8733721,
				1.1432098, 1.7855448, 1.2102433,
				1.0507759, 1.5598905, 1.8885028,
				1.1011722, 1.3064879, 1.5622698}),
		),
	},
	{
		name: "float32-broadcast2",
		aT: tensor.New(
			tensor.WithShape(6, 3),
			tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941,
				0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949,
				0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985}),
		),
		bT: tensor.New(
			tensor.WithShape(4, 6),
			tensor.WithBacking([]float32{0.77815676, 0.87001216, 0.9786183, 0.7991586,
				0.46147937, 0.7805292, 0.11827443, 0.639921,
				0.14335328, 0.9446689, 0.5218483, 0.41466194,
				0.2645556, 0.7742337, 0.45615032, 0.56843394,
				0.0187898, 0.6176355, 0.6120957, 0.616934,
				0.94374806, 0.6818203, 0.3595079, 0.43703195}),
		),
		cT: tensor.New(
			tensor.WithShape(4),
			tensor.WithBacking([]float32{0.6976312, 0.6976312, 0.6976312, 0.6976312}),
		),
		expectedOutputT: tensor.New(
			tensor.WithShape(3, 4),
			tensor.WithBacking([]float32{1.3317792, 0.9343705, 0.8733721,
				1.1432098, 1.7855448, 1.2102433,
				1.0507759, 1.5598905, 1.8885028,
				1.1011722, 1.3064879, 1.5622698}),
		),
	},
	{
		name: "float64-broadcast",
		aT: tensor.New(
			tensor.WithShape(6, 3),
			tensor.WithBacking([]float64{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941,
				0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949,
				0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985}),
		),
		bT: tensor.New(
			tensor.WithShape(4, 6),
			tensor.WithBacking([]float64{0.77815676, 0.87001216, 0.9786183, 0.7991586,
				0.46147937, 0.7805292, 0.11827443, 0.639921,
				0.14335328, 0.9446689, 0.5218483, 0.41466194,
				0.2645556, 0.7742337, 0.45615032, 0.56843394,
				0.0187898, 0.6176355, 0.6120957, 0.616934,
				0.94374806, 0.6818203, 0.3595079, 0.43703195}),
		),
		cT: tensor.New(
			tensor.WithShape(1),
			tensor.WithBacking([]float64{0.6976312}),
		),
		expectedOutputT: tensor.New(
			tensor.WithShape(3, 4),
			tensor.WithBacking([]float64{1.3317792, 0.9343705, 0.8733721,
				1.1432098, 1.7855448, 1.2102433,
				1.0507759, 1.5598905, 1.8885028,
				1.1011722, 1.3064879, 1.5622698}),
		),
	},
}

func TestGemm(t *testing.T) {
	for _, tst := range testGemm {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			aT := tst.aT
			bT := tst.bT
			cT := tst.cT
			expectedOutputT := tst.expectedOutputT

			g := NewGraph()
			a := g.NewNode()
			g.AddNode(a)
			b := g.NewNode()
			g.AddNode(b)
			c := g.NewNode()
			g.AddNode(c)
			output := g.NewNode()
			g.AddNode(output)
			g.SetWeightedEdge(g.NewWeightedEdge(output, a, 0))
			g.SetWeightedEdge(g.NewWeightedEdge(output, b, 1))
			g.SetWeightedEdge(g.NewWeightedEdge(output, c, 2))
			a.(*Node).SetTensor(aT)
			b.(*Node).SetTensor(bT)
			c.(*Node).SetTensor(cT)
			g.ApplyOperation(onnx.Operation{
				Name: "Gemm",
				Attributes: map[string]interface{}{
					"alpha":  float32(0.5),
					"beta":   float32(0.5),
					"transA": int64(1),
					"transB": int64(1),
				},
			}, output)
			err := g.Run()
			if err != nil {
				t.Fatal(err)
			}
			outputT := output.(*Node).GetTensor()
			assert.InDeltaSlice(t, expectedOutputT.Data(), outputT.Data(), 1e-6, "the two tensors should be equal.")
		})
	}
}

func BenchmarkGemm(b *testing.B) {
	aT := tensor.New(
		tensor.WithShape(6, 3),
		tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941,
			0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949,
			0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985}),
	)

	bT := tensor.New(
		tensor.WithShape(4, 6),
		tensor.WithBacking([]float32{0.77815676, 0.87001216, 0.9786183, 0.7991586,
			0.46147937, 0.7805292, 0.11827443, 0.639921,
			0.14335328, 0.9446689, 0.5218483, 0.41466194,
			0.2645556, 0.7742337, 0.45615032, 0.56843394,
			0.0187898, 0.6176355, 0.6120957, 0.616934,
			0.94374806, 0.6818203, 0.3595079, 0.43703195}),
	)

	cT := tensor.New(
		tensor.WithShape(1, 1),
		tensor.WithBacking([]float32{0.6976312}),
	)

	expectedOutputT := tensor.New(
		tensor.WithShape(3, 4),
		tensor.WithBacking([]float32{1.3317792, 0.9343705, 0.8733721,
			1.1432098, 1.7855448, 1.2102433,
			1.0507759, 1.5598905, 1.8885028,
			1.1011722, 1.3064879, 1.5622698}),
	)

	g := NewGraph()
	a := g.NewNode()
	g.AddNode(a)
	bb := g.NewNode()
	g.AddNode(bb)
	c := g.NewNode()
	g.AddNode(c)
	output := g.NewNode()
	g.AddNode(output)
	g.SetWeightedEdge(g.NewWeightedEdge(output, a, 0))
	g.SetWeightedEdge(g.NewWeightedEdge(output, bb, 1))
	g.SetWeightedEdge(g.NewWeightedEdge(output, c, 2))
	a.(*Node).SetTensor(aT)
	bb.(*Node).SetTensor(bT)
	c.(*Node).SetTensor(cT)
	for i := 0; i < b.N; i++ {
		g.ApplyOperation(onnx.Operation{
			Name: "Gemm",
			Attributes: map[string]interface{}{
				"alpha":  float32(0.5),
				"beta":   float32(0.5),
				"transA": int64(1),
				"transB": int64(1),
			},
		}, output)
		err := g.Run()
		if err != nil {
			b.Fatal(err)
		}
		outputT := output.(*Node).GetTensor()
		assert.InDeltaSlice(b, expectedOutputT.Data(), outputT.Data(), 1e-6, "the two tensors should be equal.")

	}
}
