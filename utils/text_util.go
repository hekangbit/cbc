package utils

import (
	"strconv"
	"strings"
)

const vtab = 0x0B

// TODO: check correctness, simulate Java String.compareTo，return -1, 0, 1
func CompareStrings(a, b string) int {
	return strings.Compare(a, b)
}

// TODO: need check encoding, does different encoding affect cbc dump function?
func DumpString(str string) string {
	src := []byte(str)

	var buf []byte
	buf = append(buf, '"')

	for _, b := range src {
		c := int(toUnsigned(b))
		switch {
		case c == '"':
			buf = append(buf, '\\', '"')
		case isPrintable(c):
			buf = append(buf, byte(c))
		case c == '\b':
			buf = append(buf, '\\', 'b')
		case c == '\t':
			buf = append(buf, '\\', 't')
		case c == '\n':
			buf = append(buf, '\\', 'n')
		case c == vtab:
			buf = append(buf, '\\', 'v')
		case c == '\f':
			buf = append(buf, '\\', 'f')
		case c == '\r':
			buf = append(buf, '\\', 'r')
		default:
			// oct mode
			buf = append(buf, '\\')
			buf = append(buf, strconv.FormatInt(int64(c), 8)...)
		}
	}

	buf = append(buf, '"')
	return string(buf)
}

func toUnsigned(b byte) byte {
	return b
}

func isPrintable(c int) bool {
	return ' ' <= c && c <= '~'
}
