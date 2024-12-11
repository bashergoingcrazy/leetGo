package problems

import (
	"strings"
)

func AddSpaces(s string, spaces []int) string {
	n , m := len(s) , len(spaces)
	i, j := 0, 0
	var builder strings.Builder
	builder.Grow(n+m)

	for i < n {
		if j < m && i == spaces[j] {
			builder.WriteByte(' ')
			j++
		}
		builder.WriteByte(s[i])
		i++
	}

	return builder.String()
}

