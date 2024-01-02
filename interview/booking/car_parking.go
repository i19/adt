package booking

// Given array of boolean representing where cars are parked,
// find out the min length required to cover the roof of k cars,
// all test cases passed
func carParkingForce(parked []bool, k int) int {
	p := 0
	for _, v := range parked {
		if v {
			p++
		}
	}
	if p < k {
		return -1
	}

	ans := len(parked)
	for i, v := range parked {
		if v {
			tk := k
			j := i
			for ; tk > 0 && j < len(parked); j++ {
				if parked[j] {
					tk--
				}
			}
			if tk == 0 {
				ans = min(ans, j-i)
				if ans == k {
					return ans
				}
			}
		}
	}

	return ans
}

func carParkingSlidingWindow(parked []bool, k int) int {
	n := len(parked)
	left := 0
	for ; !parked[left]; left++ {
	}
	right, found := left-1, 0
	for found < k && right < n-1 {
		right++
		if parked[right] {
			found++
		}
	}
	if found < k {
		return -1
	}

	ans := right - left + 1
	for right < n {
		left++
		for ; left < n && !parked[left]; left++ {
		}
		right++
		for ; right < n && !parked[right]; right++ {
		}
		ans = min(ans, right-left+1)
	}
	return ans

}
