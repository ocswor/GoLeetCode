package main


import "fmt"

func main() {
	// 00000100   向右移动一位 就是 00000010
	// 00000011   向右移动一位 就是 00000001
	// 右移一位相当于 除以 2
	fmt.Printf("4 右移以为 %d\n",4>>1)
	fmt.Printf("3 右移以为 %d\n",3>>1)
	fmt.Printf("% 08b  \n",3)

	// 00000100   向左移动一位 就是 00001000
	// 00000011   向左移动一位 就是 00000110
	fmt.Printf("4 左移以为 %d\n",4<<1)
	fmt.Printf("3 左移以为 %d\n",3<<1)
	fmt.Printf("% 08b \n",6)
}