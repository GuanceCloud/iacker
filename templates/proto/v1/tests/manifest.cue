package tests

import (
	"github.com/GuanceCloud/iacker/examples/petstore"
	proto "github.com/GuanceCloud/iacker/templates/proto/v1"
)

"resources": petstore.resources

"templates": {
	"proto": proto
}

"options": "templates": [
	{
		template: "proto"
		outdir:   ".build"
	},
]
