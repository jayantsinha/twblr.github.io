package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) (ret []int32) {

	for _, val := range vals {
			ret = append(ret, val*val)
	}
	return

}

func filterInts(op filterOperation, vals []int32) (ret []int32) {
	for _, val := range vals {
			if val > 2 {
				ret = append(ret, val)
			}
	}
	return
}

func concatenate(dest []string, newValues ...string) (ret []string) {
	ret = append(ret, dest...)
	for _, val := range newValues {
			ret = append(ret, val)
	}
	return
}

func equals(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
			return false
	}
	if list1 == nil && list2 == nil {
		return false
	}
	if list1 == nil || list2 == nil {
		return false
	}
	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func partialReverse(src []int, from, to int) []int {
	parr := make([]int, to-from+1)
	for i := from; i<=to; i++ {
		parr=append(parr, src[i])
	}
	//reverse
	midpoint := len(parr)/2
	for i := 0; i < midpoint; i++ {
		tmp := parr[i]
		parr[i] = parr[len(parr)-i-1]
		parr[len(parr)-i-1] = tmp
	}
	return parr
}
