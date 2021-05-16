package algorithm

import (
	"testing"
)
func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}
func Test_basic_binarysearch(t *testing.T) {

	arr := []int{1,2,4,5,6,7,8}
	target := 8
	r := 3

	result :=binarySearch(arr,target)

	assertEqual(t, r, result)

}