package testbackend

import (
	"testing"
)

func TestRun(t *testing.T) {
	for _, tc := range GetOpTypeTests("Conv") {
		tc := tc // capture range variable
		t.Run(tc().GetInfo(), tc().RunTest(testBackend, true))
	}
}
