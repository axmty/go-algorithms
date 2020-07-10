package problems

// AreThereDuplicates determines if there are duplicates in the params.
func AreThereDuplicates(params ...interface{}) bool {
	m := make(map[interface{}]bool)
	for _, p := range params {
		if _, exist := m[p]; exist {
			return true
		}
		m[p] = true
	}
	return false
}
