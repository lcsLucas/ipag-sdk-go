package utils

func ExistsIn[K comparable](list []K, element K) bool {
	for _, v := range list {
		if element == v {
			return true
		}
	}

	return false
}
