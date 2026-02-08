package util

import "strconv"

const vtab = 0x0B // 垂直制表符 \v

// DumpString 将字符串转义为可打印格式（默认UTF-8编码）
func DumpString(str string) string {
	return DumpStringWithEncoding(str, "UTF-8")
}

// DumpStringWithEncoding 将字符串转义为可打印格式（指定编码）
func DumpStringWithEncoding(s string, encoding string) string {
	// Go 字符串本身就是 UTF-8，如果需要其他编码需要转换
	// 这里简化处理，直接使用 UTF-8
	src := []byte(s)

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
			// 八进制转义
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
