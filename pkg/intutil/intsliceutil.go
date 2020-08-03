package intutil

func InArray(num int, slice []int) bool {
	for _, v := range slice {
		if num == v {
			return true
		}
	}
	return false
}
