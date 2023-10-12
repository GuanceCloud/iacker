package v1

import "github.com/GuanceCloud/iacker/pkg/resource/v1"

// Manifest is a template definitions that generates files from templates.
#Manifest: {
	name?:    string   @protobuf(1,string)
	inputs?:  #Inputs  @protobuf(2,Inputs)
	outputs?: #Outputs @protobuf(3,Outputs)
	diagnostics?: [...#Diagnostic] @protobuf(4,Diagnostic)
	layouts?: #Layouts @protobuf(5,Layouts)
}

// Inputs is a set of inputs for a template.
#Inputs: {
	vars?: {
		[string]: string
	} @protobuf(1,map[string]string)
	resources?: {
		[string]: v1.#Resource
	} @protobuf(2,map[string]pkg.resource.v1.Resource)
	errors?: {
		[string]: v1.#Error
	} @protobuf(3,map[string]pkg.resource.v1.Error)
}

// Outputs is a set of outputs for a template.
#Outputs: {
	files?: {
		[string]: #File
	} @protobuf(1,map[string]File)
}

// Layouts is a set of layouts for a template.
#Layouts: {
	files?: {
		[string]: #File
	} @protobuf(1,map[string]File)
}

// Diagnostic is a diagnostic message.
#Diagnostic: {
	message?:  string @protobuf(1,string)
	severity?: int32  @protobuf(2,int32)
}

// Output is a template output.
#File: {
	content?: string @protobuf(1,string)
}
