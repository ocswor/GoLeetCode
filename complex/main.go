package main

import "fmt"

// 时间复杂度

/*
O(1) 时间复杂度
*/
func one(n int) {
	fmt.Printf("n:%d", n)
}

/*
O(1) 时间复杂度 常数级
*/
func two(n int) {
	fmt.Printf("n:%d", n)
	fmt.Printf("n:%d", n)
	fmt.Printf("n:%d", n)
}

/*
O(n) 时间复杂度 线性
*/
func three(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("n:%d \n", n)
	}
}

/*
O(n^2) n*n n的平方
*/
func four(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("n:%d \n", n)
		}
	}
}

/*
O(log(n)) 复杂度


*/
func five(n int) {
	for i := 1; i < n; i = i * 2 {
		fmt.Printf("n:%d \n", n)
	}
}

/*
斐波那契数列 之和

O(k^n) 指数级的 非常的慢

*/
func six(n int) int {
	if n < 2 {
		return n
	}
	return six(n-2) + six(n-1)
}

/*
求 斐波那契数列 第n项和  对比 方法6  改成 O(n) 线性的级复杂度
 */
func seven(n int) int {

	if n < 2 {
		return n
	}

	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i-2] + nums[i-1]
	}
	return nums[n]
}

func main() {
	//one(1000)
	//two(1000)
	//three(1000)
	//four(1000)
	//five(1000)

	//s :=six(100) // 一时半会没结果
	//fmt.Printf("s:%d",s)
	s :=seven(1000)
	fmt.Printf("s:%d",s)
}
