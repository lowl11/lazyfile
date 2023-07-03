package folderapi

import (
	"fmt"
	"strings"
)

func replaceAllDashes(path, delimiter string) string {
	path = strings.ReplaceAll(path, "/", delimiter)
	path = strings.ReplaceAll(path, "\\", delimiter)
	return path
}

func buildObjectPath(path, name string) string {
	builder := strings.Builder{}
	builder.Grow(len(path) + len(name) + 1)

	_, _ = fmt.Fprintf(&builder, "%s/%s", path, name)
	return builder.String()
}
