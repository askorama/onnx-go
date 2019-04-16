package testbackend

import (
	"io"
	"text/template"
)

// WriteCoverageReport on w
func WriteCoverageReport(w io.Writer, test []*TestCase) {
	t := template.Must(template.New("report").Parse(reportTemplate))
	t.Execute(w, test)
}

const reportTemplate = `
|.| OpType      | Test                 | Tested | Skipped | Failed |
|-|-------------|----------------------|--------|---------|--------|
{{- range . -}}
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
