package shared

import "strings"

func IsHTTPMethod(m string) bool {
	switch strings.ToLower(m) {
	case "get", "post", "put", "delete", "patch", "options", "head":
		return true
	default:
		return false
	}
}
