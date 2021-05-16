package leetcode

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	// high 是 nums1的长度 是较短的那一个
	for low <= high {
		nums1Mid = low + (high-low)>>1 // 对nums1 进行分割 左侧 默认 num1 长度除以2 就是 nums1的分割线，后面还会调整到满足条件
		nums2Mid = k - nums1Mid        // k 表示两个数组长度之和 除以2 再减去 nums1Mid
		// 目前这两个分割线位置是暂时定义的
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			low = nums1Mid + 1
		} else {
			break
		}
	}
	fmt.Printf("nums1Mid:%d nums2Mid:%d \n",nums1Mid,nums2Mid)
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if len(nums1)+len(nums2)&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
