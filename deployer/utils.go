package deployer

import (
	"fmt"
	"strings"
)

func IncludesFormat(path string, includesString string) string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	if includesString == "" {
		return path
	}
	var includes []string
	for _, p := range strings.Split(includesString, "\n") {
		includes = append(includes, path + strings.TrimSpace(p))
	}
	return strings.Join(includes, " ")
}

func ExcludesFormat(path string, excludesString string) string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	if excludesString == "" {
		return path
	}
	var excludes []string
	for _, p := range strings.Split(excludesString, "\n") {
		excludes = append(excludes, "--exclude=" + path + strings.TrimSpace(p))
	}
	return fmt.Sprintf(" %s %s ", strings.Join(excludes, " "), path)
}
