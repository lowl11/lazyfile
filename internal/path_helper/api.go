package path_helper

import "strings"

func Build(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	builder := strings.Builder{}
	for index, item := range args {
		builder.WriteString(item)

		if index < len(args)-1 {
			builder.WriteString("/")
		}
	}

	return builder.String()
}
