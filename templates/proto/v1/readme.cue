package v1

import (
	"strings"

	template "github.com/GuanceCloud/iacker/pkg/template/v1"
)

for rsname, rsinfo in *inputs.resources | {} {
	outputs: files: "resources/\(strings.ToLower(rsname))/v1/README.md": template.#File & {
		_lowername: strings.ToLower(rsname)

		content: """
		# \(rsinfo.title.en)

		\(rsinfo.description.en)
		"""
	}
}
