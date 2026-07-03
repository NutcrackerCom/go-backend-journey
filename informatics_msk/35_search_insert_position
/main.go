func searchInsert(nums []int, target int) int {
	l := -1
	r := len(nums)
	for r-l > 1 {
		m := l + (r-l)/2
		fmt.Println(m)
		if nums[m] < target {
			l = m
		} else {
			r = m
		}
	}
   return r
}
