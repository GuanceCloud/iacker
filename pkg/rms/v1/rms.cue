// import "pkg/template/v1/template.proto";
package v1

import "github.com/GuanceCloud/iacker/pkg/resource/v1"

#Manifest: {
	options?: #Options @protobuf(1,Options)
	resources?: {
		[string]: v1.#Resource
	} @protobuf(2,map[string]pkg.resource.v1.Resource)
	errors?: {
		[string]: v1.#Error
	} @protobuf(3,map[string]pkg.resource.v1.Error)
	// map<string, pkg.template.v1.Manifest> templates = 4;

}

#Options: {
	templates?: [...#TemplateOptions] @protobuf(1,TemplateOptions)
}

// TemplateOptions is a set of options for a template.
#TemplateOptions: {
	template?: string @protobuf(1,string)
	outdir?:   string @protobuf(2,string)
	vars?: {
		[string]: string
	} @protobuf(3,map[string]string)
}
