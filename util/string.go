package util

import (
	"fmt"
	"reflect"
	"strings"
)

func isOctalDigit(c byte) bool {
	return c >= '0' && c <= '7'
}

func unescapeOctal(octal string) (byte, error) {
	val := 0
	for i := 0; i < 3; i++ {
		d := octal[i] - '0'
		if d > 7 {
			return 0, fmt.Errorf("invalid octal digit '%c'", octal[i])
		}
		val = val*8 + int(d)
	}
	if val > 255 {
		return 0, fmt.Errorf("octal value %d out of byte range", val)
	}
	return byte(val), nil
}

func unescapeSeq(ch byte) (byte, error) {
	switch ch {
	case 'n':
		return '\n', nil
	case 't':
		return '\t', nil
	case 'r':
		return '\r', nil
	case '\\':
		return '\\', nil
	case '"':
		return '"', nil
	case '\'':
		return '\'', nil
	// add more, like \b, \f 等
	default:
		return 0, fmt.Errorf("unknown escape sequence '\\%c'", ch)
	}
}

// Parse the string literal enclosed in quotation marks like "hello\nworld"
func StringValue(image string) (string, error) {
	if len(image) < 2 || image[0] != '"' || image[len(image)-1] != '"' {
		return "", fmt.Errorf("invalid string literal: missing surrounding quotes")
	}

	inner := image[1 : len(image)-1] // remove begin end <">
	var buf strings.Builder
	pos := 0

	for pos < len(inner) {
		// find next <\>
		idx := strings.IndexByte(inner[pos:], '\\')
		if idx == -1 {
			// no more escape character, write the remaining part and end
			buf.WriteString(inner[pos:])
			break
		}

		actualIdx := pos + idx // The absolute position of the backslash in the original "inner"

		// writing The part before backslash
		if actualIdx > pos {
			buf.WriteString(inner[pos:actualIdx])
		}

		// Check if it is a three-digit octal escape sequence (\\123)
		if actualIdx+3 < len(inner) &&
			isOctalDigit(inner[actualIdx+1]) &&
			isOctalDigit(inner[actualIdx+2]) &&
			isOctalDigit(inner[actualIdx+3]) {

			octal := inner[actualIdx+1 : actualIdx+4]
			ch, err := unescapeOctal(octal)
			if err != nil {
				return "", fmt.Errorf("invalid octal escape at position %d: %w", actualIdx, err)
			}
			buf.WriteByte(ch)
			pos = actualIdx + 4 // Skip the backslash and the three digits

		} else {
			// Common single-character escape sequences (such as \n, \t, \\\, etc.)
			if actualIdx+1 >= len(inner) {
				return "", fmt.Errorf("incomplete escape at end of string")
			}
			ch, err := unescapeSeq(inner[actualIdx+1])
			if err != nil {
				return "", fmt.Errorf("invalid escape sequence at position %d: %w", actualIdx, err)
			}
			buf.WriteByte(ch)
			pos = actualIdx + 2 // Skip the backslashes and escape characters
		}
	}

	return buf.String(), nil
}

// TODO: need check again
func ToString(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	// get actual type
	t := reflect.TypeOf(obj)
	typeName := t.String()

	// Get hash based on address
	val := reflect.ValueOf(obj)
	var hashCode int32
	switch val.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		hashCode = int32(val.Pointer() & 0x7fffffff)
	default:
		hashCode = HashCode(fmt.Sprintf("%v", obj))
	}

	return fmt.Sprintf("%s@%x", typeName, uint32(hashCode))
}

func HashCode(s string) int32 {
	var h int32 = 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int32(s[i])
	}
	return h
}
