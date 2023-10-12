package terraform

import (
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	resource "github.com/GuanceCloud/iacker/templates/terraform/resource"
	datasource "github.com/GuanceCloud/iacker/templates/terraform/datasource"
)

name: "terraform"

inputs: template.#Inputs

_rs: resource & { "inputs": inputs }
_ds: datasource & { "inputs": inputs } 

outputs: files: _rs.outputs.files & _ds.outputs.files

diagnostics: _rs.diagnostics + _ds.diagnostics
