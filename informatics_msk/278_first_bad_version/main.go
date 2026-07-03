func firstBadVersion(n int) int {
	rp := n
	lp := 0
	for rp-lp > 1 {
		m := lp + (rp-lp)/2
		if isBadVersion(m) {
			rp = m
		} else {
			lp = m
		}
	}
	return rp
}
