package algorithm

func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		//mid := (l + r) / 2 // 这里 如果是l 和 r过大 则会内存溢出
		mid := (r-l)/2 + l // 这里使用减法
		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
