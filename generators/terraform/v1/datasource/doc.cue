package datasource

import (
	naming "github.com/GuanceCloud/iacker/pkg/helpers/naming"
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
)

// Generate Terraform data source documentation
for rsname, rsinfo in inputs.resources {
	if (*rsinfo.meta.datasource | false) {
		outputs: files: "internal/datasources/\((naming.#Snake & {name: rsinfo.plural}).lower)/README.md": template.#File & {
			content: *rsinfo.description.en | ""
		}
	}
}
