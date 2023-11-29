package datasource

import (
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
)

inputs: template.#Inputs

diagnostics: [...template.#Diagnostic]

provider: inputs.vars.provider
