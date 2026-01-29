package strutil

func SplitAndTrim(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			if i > start {
				trimmed := TrimSpace(s[start:i])
				if trimmed != "" {
					result = append(result, trimmed)
				}
			}
			start = i + 1
		}
	}
	if start < len(s) {
		trimmed := TrimSpace(s[start:])
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func TrimSpace(s string) string {
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}
