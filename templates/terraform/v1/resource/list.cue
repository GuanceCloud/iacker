package resource

import (
	gotemplate "text/template"

	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	naming "github.com/GuanceCloud/iacker/pkg/helpers/naming"
)

_resource_list_template: """
// Code generated by Guance Cloud Code Generation Pipeline. DO NOT EDIT.

package \(provider)

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	{{ range $_, $name := . -}}
	"github.com/GuanceCloud/terraform-provider-guance/internal/resources/{{ .lower }}"
	{{ end }}
)

// Resources defines the resources implemented in the provider.
func (p *guanceProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		{{- range $_, $name := . }}
		{{ .lower }}.New{{ .uppercamel }}Resource,
		{{- end }}
	}
}
"""

outputs: files: "\(provider)/resources.go": template.#File & {
	content: gotemplate.Execute(_resource_list_template, {
		for rsname, rsinfo in inputs.resources {
			if !(*rsinfo.meta.datasource | false) {
				(rsname): naming.#UpperCamel & {
					"name": rsname
				}
			}
		}
	})
}