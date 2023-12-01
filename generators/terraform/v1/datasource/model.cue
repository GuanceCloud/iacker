package datasource

import (
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	naming "github.com/GuanceCloud/iacker/pkg/helpers/naming"
	builder "github.com/GuanceCloud/iacker/templates/terraform/v1/builder"
)

// Generate go types of resource model
for rsname, rsinfo in inputs.resources {
	if (*rsinfo.meta.datasource | false) {
		outputs: files: "internal/datasources/\((naming.#Snake & {"name": rsinfo.plural}).lower)/model.go": template.#File & {
			_builder: builder.#StructBuilder & {
				"pkg":  (naming.#Snake & {"name": rsinfo.plural}).lower
				"name": rsname
				"rs":   rsinfo
				"isds": true
			}
			content: _builder.output
		}
	}
}
