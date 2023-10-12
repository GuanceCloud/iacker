package allinone

import (
	template "github.com/GuanceCloud/iacker/pkg/template/v1"
)

// Pet is the pet resource definition
resources: "Pet": {
	plural: "pets"
	title: {
		zh: "宠物"
		en: "Pet"
	}
	description: {
		zh: "宠物很可爱，包括猫猫狗狗等"
		en: "Pets are cute, including cats, dogs, etc."
	}
	model: "Pet"
}

// Pet model
resources: "Pet": models: "Pet": {
	title: {
		zh: "宠物"
		en: "Pet"
	}
	properties: [
		{
			name: "id"
			title: {
				zh: "ID"
				en: "ID"
			}
			schema: {
				type:     "integer"
				required: true
			}
		},
	]
}

// Template definition
templates: "foo": {
	name: "basic"

	inputs: template.#Inputs

	diagnostics: [...template.#Diagnostic]

	outputs: files: "README.md": {
		content: "Hello, World!"
	}

	layouts: files: "Init.md": {
		content: "This file is only generated when --init is specified!"
	}
}

// Template options
options: templates: [
	{
		template: "foo"
		outdir:   ".build"
	},
]
