package leetcode

func maxArea(arr []int) int {

	max, start, end := 0, 0, len(arr)-1
	for start < end {
		width := end - start
		high := 0
		if arr[start] < arr[end] { // 如果左边 的高度 低于 右边的高度
			high = arr[start] // 那么取出左边的高度  短板 计算一下 当前的面积
			start++
		} else { // 如果 左边的高度 高于 右边的高度
			high = arr[end] //那么去除 右边的高度  计算一下当前的面积
			end--
		}

		tmp := width * high
		if tmp > max {
			max = tmp
		}
	}
	return max
}
