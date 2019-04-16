package testreport

import "github.com/owulveryck/onnx-go/backend/testbackend"

// Coverage returns (100 - skipped *100 / tested)
func Coverage(tests []*testbackend.TestCase) float64 {
	var skipped, tested float64
	for _, tc := range tests {
		if !tc.Skipped {
			skipped++
		}
		if tc.Tested {
			tested++
		}
	}
	return (float64(100) - skipped*float64(100)/tested)
}
