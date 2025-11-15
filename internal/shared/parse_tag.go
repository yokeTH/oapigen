package shared

import "strings"

func ParseTag(raw, key string) string {
	st := raw
	for st != "" {
		st = strings.TrimLeft(st, " ")
		if !strings.HasPrefix(st, key+`:"`) {
			idx := strings.Index(st, " ")
			if idx < 0 {
				break
			}
			st = st[idx+1:]
			continue
		}
		st = st[len(key)+2:]
		end := strings.Index(st, `"`)
		if end < 0 {
			return ""
		}
		val := st[:end]
		return val
	}
	return ""
}
