package resource

import (
    "strings"

	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	builder "github.com/GuanceCloud/iacker/templates/terraform/v1/builder"
)

// Generate go types of resource model
for rsname, rsinfo in inputs.resources {
	if !(*rsinfo.meta.datasource | false) {
		outputs: files: "internal/resources/\(strings.ToLower(rsname))/model.go": template.#File & {
			_builder: builder.#StructBuilder & {
				"pkg":  strings.ToLower(rsname)
				"name": rsname
				"rs":   rsinfo
			}
			content: _builder.output
		}
	}
}
