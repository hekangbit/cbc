package utils

func Align(n, alignment int64) int64 {
	return (n + alignment - 1) / alignment * alignment
}
