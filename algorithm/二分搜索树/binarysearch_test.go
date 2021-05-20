package algorithm

import (
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func assertEqualB(t *testing.B, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}
func Test_basic_binarysearch(t *testing.T) {

	arr := []int{1, 2, 4, 5, 6, 7, 8}
	target := 8
	r := 6
	result := binarySearch(arr, target)

	assertEqual(t, r, result)
}

func BenchmarkNewTestflight(b *testing.B) {
	arr := []int{1, 2, 4, 5, 6, 7, 8}
	target := 8
	r := 6
	for i := 0; i < b.N; i++ {
		result := binarySearch(arr, target)
		assertEqualB(b, r, result)
	}
}
