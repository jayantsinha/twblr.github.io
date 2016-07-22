package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) (ret []string) {

	for key, val := range data {
		for _, intrst := range val {
			if intrst == interest {
				ret = append(ret, key)
			}
		}
	}
	return ret
}

func contains(src []string, elem string) bool {
	for _, s := range src {
		if s == elem {
			return true
		}
	}

	return false
}
