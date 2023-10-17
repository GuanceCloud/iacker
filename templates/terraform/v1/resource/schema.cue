package resource

import (
    "strings"

	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	builder "github.com/GuanceCloud/iacker/templates/terraform/v1/builder"
)

// Generate Terraform schema
for rsname, rsinfo in inputs.resources {
	if !(*rsinfo.meta.datasource | false) {
		outputs: files: "internal/resources/\(strings.ToLower(rsname))/schema.go": template.#File & {
			_builder: builder.#SchemaBuilder & {
				"name": rsname
				"pkg":  strings.ToLower(rsname)
				"rs":   rsinfo
			}
			content: _builder.output
		}
	}
}
