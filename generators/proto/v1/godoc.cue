package v1

import (
	"strings"

	template "github.com/GuanceCloud/iacker/pkg/template/v1"
)

// Generate godoc files
for rsname, rsinfo in *inputs.resources | {} {
	outputs: files: "resources/\(strings.ToLower(rsname))/v1/doc.go": template.#File & {
		_lowername: strings.ToLower(rsname)

		content: """
		/*
		Package v1

		# \(rsinfo.title.en)

		\(rsinfo.description.en)
		*/
		package v1
		"""
	}
}
