package iteration

// Repeatrecursive takes a rune and returns a string with that rune n times
func RepeatRecursive(r rune, times int) (s string) {

	// define the function otherwise Go can't handle the recursion
	var stringBuild func(int, rune, string) string
	stringBuild = func(iterations int, r rune, s string) string {
		if iterations == 0 {
			return s
		}
		return s + string(r) + stringBuild(iterations-1, r, s)
	}

	return stringBuild(times, r, "")
}

// RepeatIterative takes a rune and returns a string with that rune n times
func RepeatIterative(r rune, times int) (s string) {

	for i := 0; i < times; i++ {
		s += string(r)
	}
	return
}
