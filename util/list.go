package util

func ContainsAll(list1, list2 []string) bool {
	elements := make(map[string]bool)

	for _, item := range list1 {
		elements[item] = true
	}

	for _, item := range list2 {
		if !elements[item] {
			return false
		}
	}

	return true
}
