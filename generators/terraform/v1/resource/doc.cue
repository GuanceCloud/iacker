package resource

import (
    naming "github.com/GuanceCloud/iacker/pkg/helpers/naming"
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
)

// Generate Terraform resource documentation
for rsname, rsinfo in inputs.resources {
	if !(*rsinfo.meta.datasource | false) {
		outputs: files: "internal/resources/\((naming.#UpperCamel & {name: rsname}).lower)/README.md": template.#File & {
			content: *rsinfo.description.en | ""
		}
	}
}
