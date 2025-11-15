package shared

import "strings"

func ExtractColonPathParam(path string) []string {
	result := []string{}
	parts := strings.SplitSeq(path, "/")
	for p := range parts {
		if after, ok := strings.CutPrefix(p, ":"); ok {
			result = append(result, after)
		}
	}
	return result
}
