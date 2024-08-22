package windowSlide

func CanCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	len1 := 0
	for l, r := 0, 0; l < len(gas); l++ {
		for sum >= 0 {
			if len1 == len(gas) {
				return l
			}
			r = (l + len1) % len(gas)
			sum += gas[r] - cost[r]
			len1++
		}
		sum -= gas[l] - cost[l]
		len1--
	}
	return -1
}
