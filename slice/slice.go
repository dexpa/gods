package slice

func Search(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		if f(i) {
			return i
		}
		i++
	}
	return -1
}

type StringSlice []string

// Contains searches for x in StringSlice and
// returns the index of x in slice if slice contains x, otherwise returns -1
func (s StringSlice) Contains(x string) int {
	return Search(len(s), func(i int) bool { return s[i] == x })
}

type IntSlice []int

// Contains searches for x in StringSlice and
// returns the index of x in slice if slice contains x, otherwise returns -1
func (s IntSlice) Contains(x int) int {
	return Search(len(s), func(i int) bool { return s[i] == x })
}

type Float64Slice []float64

// Contains searches for x in StringSlice and
// returns the index of x in slice if slice contains x, otherwise returns -1
func (s Float64Slice) Contains(x float64) int {
	return Search(len(s), func(i int) bool { return s[i] == x })
}