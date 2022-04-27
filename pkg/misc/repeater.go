package misc

func Repeater(s string, n int) string {
	val := ""
	for i := 0; i < n; i++ {
		val += s
	}
	return val
}
