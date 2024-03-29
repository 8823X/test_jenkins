package arithmetic

func SingleNumber(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, n := range nums {
			if 1<<i&n > 0 {
				count++
			}
		}
		if count%3 != 0 {
			res |= 1 << i
		}
	}
	return res
}
