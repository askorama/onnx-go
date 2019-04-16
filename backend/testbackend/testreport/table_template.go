package testreport

import "text/template"

// ReportTable ...
var ReportTable = template.Must(template.New("report").Parse(ReportTemplate))

// ReportTemplate ...
const ReportTemplate = `
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
{{- define "red"}}![#f03c15](https://placehold.it/15/f03c15/000000?text=+){{ end }}
{{- define "green"}}![#c5f015](https://placehold.it/15/c5f015/000000?text=+){{ end}}
{{- define "orange"}}![#ffa500](https://placehold.it/15/FFA500/000000?text=+){{end}}
`
