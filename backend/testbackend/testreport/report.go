package testreport

import (
	"io"
	"text/template"

	"github.com/owulveryck/onnx-go/backend/testbackend"
)

// WriteCoverageReport on w
func WriteCoverageReport(w io.Writer, test []*testbackend.TestCase, t *template.Template) {
	t.Execute(w, test)
}

// ByStatus is a wrapper to sort the tests by failure, success, skipped
type ByStatus []*testbackend.TestCase

func (a ByStatus) Len() int      { return len(a) }
func (a ByStatus) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByStatus) Less(i, j int) bool {
	switch {
	case a[i].Failed && !a[j].Failed:
		return true
	case !a[i].Failed && a[j].Skipped:
		return true
	default:
		return false
	}
}
