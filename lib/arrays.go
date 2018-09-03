package lib

// Concat concatenates two arrays of integers
func Concat(a []int, b []int) []int {

	return append(a, b...)
}
