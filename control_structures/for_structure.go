package control_structures

func collatzChainLength(n int) (itr int) {
	
	for itr =0; ;itr++ {
		if n%2==0 {
			n = n/2
		} else {
			n = (3*n)+1
		}
		if n==1 {
			break
		}
	}

	return itr+1
}
