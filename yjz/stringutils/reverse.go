// Package stringutil contains utility functions for working with strings.
package xxx

// 包名随便取,调Reverse方法时,用这个xxx就可以

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/**
从集合里面取值
*/
func GetFromMap(s string) int {

	maps := make(map[string]int)
	maps["张三"] = 20

	return maps[s]
}
