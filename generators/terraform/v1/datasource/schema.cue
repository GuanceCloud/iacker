package schema

import (
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	naming "github.com/GuanceCloud/iacker/pkg/helpers/naming"
	builder "github.com/GuanceCloud/iacker/templates/terraform/v1/builder"
)

// Generate Terraform schema
for rsname, rsinfo in inputs.resources {
	if (*rsinfo.meta.datasource | false) {
		outputs: files: "internal/datasources/\((naming.#Snake & {"name": rsinfo.plural}).lower)/schema.go": template.#File & {
			_builder: builder.#SchemaBuilder & {
				"pkg":  (naming.#Snake & {"name": rsinfo.plural}).lower
				"name": rsname
				"rs":   rsinfo
				"isds": true
			}
			content: _builder.output
		}
	}
}
