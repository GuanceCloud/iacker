package v1

import (
	"github.com/GuanceCloud/iacker/pkg/resource/v1"
	v1_1 "github.com/GuanceCloud/iacker/pkg/template/v1"
)

#Manifest: {
	options?: #Options @protobuf(1,Options)
	resources?: {
		[string]: v1.#Resource
	} @protobuf(2,map[string]pkg.resource.v1.Resource)
	errors?: {
		[string]: v1.#Error
	} @protobuf(3,map[string]pkg.resource.v1.Error)
}

#Options: {
	templates?: {
		[string]: v1_1.#Options
	} @protobuf(1,map[string]pkg.template.v1.Options)
}
