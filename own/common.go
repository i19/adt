package own

func min(args ...int) int {
	v := args[0]
	for _, arg := range args {
		if arg < v {
			v = arg
		}
	}
	return v
}
func max(args ...int) int {
	v := args[0]
	for _, arg := range args {
		if arg > v {
			v = arg
		}
	}
	return v
}
