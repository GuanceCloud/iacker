package v1

import (
	"list"
	"strings"
	gotemplate "text/template"

	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	resource "github.com/GuanceCloud/iacker/pkg/resource/v1"
)

#Resource: {
	name: string

	rs: resource.#Resource

	output: string

	_lowername: strings.ToLower(name)

	_template: """
	syntax = "proto3";

	package pkg.resources.\( _lowername ).v1;

	option go_package = "github.com/GuanceCloud/openapi/pkg/resources/\( _lowername )/v1;v1";

	"""

	_model_template: """
		/*
		{{- if .v.title.en }}
		{{ .v.title.en }}
		{{- end }}
		*/
		message {{ .name }} {
			{{- range $v := .properties }}

		{{ $v }}
			{{- end }}
		}
		"""

	_prop_template: """
			/*
			{{ .v.title.en }}

			{{- if .v.description.en }}
			{{ .v.description.en }}
			{{- end }}
			*/
			{{ template "type" .v.schema }} {{ .v.name }} = {{ .index }};

		{{- define "type" }}
		{{- if eq .type "array" -}}
			repeated {{ template "elem" .elem }}
		{{- else if eq .type "object" -}}
			{{ .model }}
		{{- else if eq .type "ref" -}}
			{{ template "optional" . }}string
		{{- else if eq .type "integer" -}}
			{{ template "optional" . }}int64
		{{- else if eq .type "boolean" -}}
			{{ template "optional" . }}bool
		{{- else -}}
			{{ template "optional" . }}{{ .type }}
		{{- end -}}
		{{- end }}

		{{- define "elem" }}
		{{- if eq .type "object" -}}
			{{ .model }}
		{{- else if eq .type "ref" -}}
			string
		{{- else if eq .type "integer" -}}
			int64
		{{- else if eq .type "boolean" -}}
			bool
		{{- else -}}
			{{ .type }}
		{{- end -}}
		{{- end }}

		{{- define "optional" }}
		{{- if .required }}{{ else }}optional {{ end -}}
		{{- end }}
		"""

	_models: [
		for mname, minfo in rs.models {
			gotemplate.Execute(_model_template, {
				"name": mname
				"v":    minfo
				properties: [
					for i, pinfo in minfo.properties {
						gotemplate.Execute(_prop_template, {
							"v":     pinfo
							"index": i + 1
						})
					},
				]
			})
		},
	]

	output: strings.Join(list.Concat([[_template], _models]), "\n")
}

for rsname, rsinfo in *inputs.resources | {} {
	outputs: files: "resources/\(strings.ToLower(rsname))/v1/\(strings.ToLower(rsname)).proto": template.#File & {
		_gen: #Resource & {
			"name": rsname
			"rs":   rsinfo
		}
		content: _gen.output
	}
}
