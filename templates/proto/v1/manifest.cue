package v1

import (
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
)

name: "proto"

inputs: template.#Inputs

outputs: template.#Outputs

diagnostics: [...template.#Diagnostic]
