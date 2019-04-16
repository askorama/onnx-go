package testbackend

import (
	"io"
	"sort"
	"text/template"
)

// WriteCoverageReport on w
func WriteCoverageReport(w io.Writer, test []*TestCase) {
	t := template.Must(template.New("report").Parse(reportTemplate))
	sort.Sort(byStatus(test))
	t.Execute(w, test)
}

const reportTemplate = `
|.| OpType      | Test                 | Tested | Skipped | Failed |
|-|-------------|----------------------|--------|---------|--------|
{{ range . -}}
|{{template "flag" .}}|{{ .OpType }}|{{ .Title }}|{{ .Tested }}|{{ .Skipped }}|{{ .Failed }}|
{{ end }}
{{- define "flag"}}
    {{- if .Failed -}}
	{{- template "red" -}}
    {{- else -}}
	{{- if .Skipped -}}
	    {{- template "orange" -}}
	{{- else -}}
	    {{- template "green" -}}
	{{- end -}}
    {{- end -}}
{{- end -}}
{{define "red"}}![#f03c15](https://placehold.it/15/f03c15/000000?text=+){{ end }}
{{define "green"}}![#c5f015](https://placehold.it/15/c5f015/000000?text=+){{ end}}
{{define "orange"}}![#ffa500](https://placehold.it/15/FFA500/000000?text=+){{end}}
`

type byStatus []*TestCase

func (a byStatus) Len() int      { return len(a) }
func (a byStatus) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byStatus) Less(i, j int) bool {
	switch {
	case a[i].Failed && !a[j].Failed:
		return true
	case !a[i].Failed && a[j].Skipped:
		return true
	default:
		return false
	}
}
