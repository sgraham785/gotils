package str

// LastString return the last string in slice of string
// example:
// s1 := lastString(strings.Split("example.txt", "."))
// returns txt
func LastString(ss []string) string {
	return ss[len(ss)-1]
}
