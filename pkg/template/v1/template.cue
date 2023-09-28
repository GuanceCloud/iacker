package v1

import (
	"github.com/GuanceCloud/iacker/pkg/resource/v1"
	v1_1 "github.com/GuanceCloud/iacker/pkg/testing/v1"
)

// Manifest is a template definitions that generates files from templates.
#Manifest: {
	inputs?:  #Inputs  @protobuf(1,Inputs)
	outputs?: #Outputs @protobuf(2,Outputs)
	diagnostics?: [...#Diagnostic] @protobuf(3,Diagnostic)
}

// Inputs is a set of inputs for a template.
#Inputs: {
	options?: #Options @protobuf(1,Options)
	resources?: {
		[string]: v1.#Resource
	} @protobuf(2,map[string]pkg.resource.v1.Resource)
	errors?: {
		[string]: v1.#Error
	} @protobuf(3,map[string]pkg.resource.v1.Error)
	examples?: [...v1_1.#ExampleSuite] @protobuf(4,pkg.testing.v1.ExampleSuite)
}

// Outputs is a set of outputs for a template.
#Outputs: {
	files?: {
		[string]: #File
	} @protobuf(1,map[string]File)
}

// Options is a set of options for a template.
#Options: {
	outdir?: string @protobuf(1,string)
	vars?: {
		[string]: string
	} @protobuf(2,map[string]string)
	provisioners?: [...#Provisioner] @protobuf(3,Provisioner)
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

// Provisioner is a provisioner.
#Provisioner: {
	type?:  string            @protobuf(1,string)
	shell?: #ShellProvisioner @protobuf(2,ShellProvisioner)
}

// ShellProvisioner is a shell provisioner.
#ShellProvisioner: {
	inline?: string @protobuf(1,string)
}
