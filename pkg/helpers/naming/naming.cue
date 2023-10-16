package naming

import (
	"strings"
)

#Snake: {
	name:       string
	snake:      name
	uppercamel: strings.Join([ for _, token in strings.Split(snake, "_") {strings.ToTitle(strings.ToLower(token))}], "")
	lowercamel: strings.ToCamel(uppercamel)
	lower:      strings.ToLower(lowercamel)
}

#UpperCamel: {
	snake:      "unimplemented"
	name:       string
	uppercamel: name
	lowercamel: strings.ToCamel(uppercamel)
	lower:      strings.ToLower(lowercamel)
}
