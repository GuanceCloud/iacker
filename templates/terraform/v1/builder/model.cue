package builder

import (
	"strings"
	gotemplate "text/template"

	resource "github.com/GuanceCloud/iacker/pkg/resource/v1"
	naming "github.com/GuanceCloud/iacker/pkg/helpers/naming"
)

#StructBuilder: {
	name: string

	pkg: string

	rs: resource.#Resource

	isds: *false | bool

	_resource_template: """
		// Code generated by Guance Cloud Code Generation Pipeline. DO NOT EDIT.

		package {{ .pkg }}

		import (
			"github.com/hashicorp/terraform-plugin-framework/types"

			"github.com/GuanceCloud/terraform-provider-guance/internal/consts"
		)

		// {{ .name.lowercamel }}ResourceModel maps the resource schema data.
		type {{ .name.lowercamel }}ResourceModel struct {
			ID types.String `tfsdk:"id"`
			CreatedAt types.String `tfsdk:"created_at"`
			{{- range $v := .properties }}
			{{ $v }}
			{{- end }}
		}

		// GetId returns the ID of the resource.
		func (m *{{ .name.lowercamel }}ResourceModel) GetId() string {
			return m.ID.ValueString()
		}

		// SetId sets the ID of the resource.
		func (m *{{ .name.lowercamel }}ResourceModel) SetId(id string) {
			m.ID = types.StringValue(id)
		}

		// GetResourceType returns the type of the resource.
		func (m *{{ .name.lowercamel }}ResourceModel) GetResourceType() string {
			return consts.TypeName{{ .name.uppercamel }}
		}

		// SetCreatedAt sets the creation time of the resource.
		func (m *{{ .name.lowercamel }}ResourceModel) SetCreatedAt(t string) {
			m.CreatedAt = types.StringValue(t)
		}
		"""

	_data_source_template: _resource_template + """
		// {{ .name.lowercamel }}DataSourceModel maps the resource schema data.
		type {{ .name.lowercamel }}DataSourceModel struct {
			Items []*{{ .name.lowercamel }}ResourceModel `tfsdk:"items"`
			Filters []*sdk.Filter `tfsdk:"filters"`
			MaxResults types.Int64 `tfsdk:"max_results"`
			ID      types.String   `tfsdk:"id"`
		}
		"""

	_model_template: """
		// {{ .name.uppercamel }} maps the resource schema data.
		type {{ .name.uppercamel }} struct {
			{{- range $v := .properties }}
			{{ $v }}
			{{- end }}
		}
		"""

	_prop_template: """
		{{ .name.uppercamel }} {{ template "type" .v.schema }} `tfsdk:"{{ .name.snake }}"`

		{{- define "type" }}
		{{- if eq .type "array" -}}
			[]{{ template "elem" .elem }}
		{{- else if eq .type "object" -}}
			*{{ .model }}
		{{- else if eq .type "ref" -}}
			types.String
		{{- else -}}
			{{ template "primitive" . }}
		{{- end -}}
		{{- end }}

		{{- define "elem" }}
		{{- if eq .type "object" -}}
			*{{ .model }}
		{{- else if eq .type "ref" -}}
			types.String
		{{- else -}}
			{{ template "primitive" . }}
		{{- end -}}
		{{- end }}

		{{- define "primitive" }}
		{{- if eq .type "integer" -}}
			types.Int64
		{{- else if eq .type "boolean" -}}
			types.Bool
		{{- else if eq .type "float" -}}
			types.Float64
		{{- else if eq .type "string" -}}
			types.String
		{{- else -}}
			unsupported type: {{ .type }}
		{{- end -}}
		{{- end }}
		"""

	_resource: gotemplate.Execute(_resource_template, {
		"name": naming.#UpperCamel & {"name": name}
		"pkg":  pkg
		"v":      rs
		properties: [
			for i, pinfo in rs.models[name].properties {
				gotemplate.Execute(_prop_template, {
					"name":  naming.#Snake & {"name": pinfo.name}
					"v":     pinfo
					"index": i + 1
				})
			},
		]
	})

	_data_source: gotemplate.Execute(_data_source_template, {
		"name": naming.#UpperCamel & {"name": name}
		"pkg":  pkg
		"v":      rs
		properties: [
			for i, pinfo in rs.models[name].properties {
				gotemplate.Execute(_prop_template, {
					"name":  naming.#Snake & {"name": pinfo.name}
					"v":     pinfo
					"index": i + 1
				})
			},
		]
	})

	_models: [
		for mname, minfo in (*rs.models | {}) if mname != name {
			gotemplate.Execute(_model_template, {
				"name": naming.#UpperCamel & {"name": mname}
				"v":    minfo
				properties: [
					for i, pinfo in minfo.properties {
						gotemplate.Execute(_prop_template, {
							"name":  naming.#Snake & {"name": pinfo.name}
							"v":     pinfo
							"index": i + 1
						})
					},
				]
			})
		},
	]

	_block_sep: "\n" * 2
	if !isds {
		output: _resource + _block_sep + strings.Join(_models, _block_sep)
	}
	if isds {
		output: _data_source + _block_sep + strings.Join(_models, _block_sep)
	}
}
