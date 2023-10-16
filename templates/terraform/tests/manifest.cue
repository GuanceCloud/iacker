package tests

import (
	"github.com/GuanceCloud/iacker/examples/petstore"
	"github.com/GuanceCloud/iacker/templates/terraform"
)

resources: petstore.resources

templates: "terraform": terraform

options: templates: [
	{
		template: "terraform"
		outdir:   ".build"
		vars: {provider: "guance"}
	},
]
