package array

import (
	"fmt"
	"testing"
)

type para1 struct {
	nums   []int
	target int
}

type question1 struct {
	para1
	ans1
}

type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},
	}
	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v 【output】:%v\n", p, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
