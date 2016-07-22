package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) (ret []int32) {

	for _, val := range vals {
			ret = append(ret, val*val)
	}
	return ret

}

func filterInts(op filterOperation, vals []int32) (ret []int32) {
	for _, val := range vals {
			if val > 2 {
				ret = append(ret, val)
			}
	}
	return ret
}

func concatenate(dest []string, newValues ...string) (ret []string) {
	n := dest
	copy(ret, n)
	for _, val := range dest {
			ret = append(ret, val)
	}

	return ret
}

func equals(list1 []string, list2 []string) bool {
	return false
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
