package _8_Heap_and_Priority_Queue

import (
	"errors"
	"fmt"
	"strconv"
)

type MaxHeap struct {
	data Array
}

func NewMaxHeap(capacity int) *MaxHeap {

	return &MaxHeap{
		NewArray(capacity),
	}
}

/*
// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
*/

func (mp *MaxHeap) parent(index int) (int, error) {

	if index == 0 {
		return -1, errors.New("index-0 doesn't have parent.")
	}
	return (index - 1) / 2, nil
}

/*
返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
*/
func (mp *MaxHeap) leftChild(index int) int {

	return index*2 + 1
}

/*
 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
*/
func (mp *MaxHeap) rigthChild(index int) int {

	return index*2 + 2
}

func compare(a interface{}, b interface{}) int {

	//if a.(int) > b.(int) {
	//	return 1
	//} else if a.(int) < b.(int) {
	//	return -1
	//} else {
	//	return 0
	//}

	switch a.(type) {
	case string:
		// 将interface转为string字符串类型
		a_s, _ := a.(string)
		b_s, _ := b.(string)

		a_int, _ := strconv.Atoi(a_s)
		b_int, _ := strconv.Atoi(b_s)
		if a_int > b_int {
			return 1
		} else if a_int < b_int {
			return -1
		} else {
			return 0
		}
	case int32:
		// 将interface转为int32类型
		op, ok := a.(int32)
		fmt.Println(op, ok)
	case int:
		// 将interface转为int64类型
		a_int, _ := a.(int)
		b_int, _ := b.(int)

		if a_int > b_int {
			return 1
		} else if a_int < b_int {
			return -1
		} else {
			return 0
		}

	default:
		fmt.Println("unknown")
	}
	return -1
}
func compareInt(a int, b int) int {

	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}

}

func (mp *MaxHeap) shiftUp(k int) error {
	var p_index int
	var err error

	for k > 0 {
		if p_index, err = mp.parent(k); err != nil {
			return err
		}
		p_v, err := mp.data.Get(p_index)

		if err != nil {
			return err
		}
		k_v, err := mp.data.Get(k)
		if err != nil {
			return err
		}

		if compare(p_v, k_v) < 0 {
			if err := mp.data.swap(k, p_index); err != nil {
				return err
			}
		} else {
			k = p_index
		}
	}

	return nil
}

func (mp *MaxHeap) shiftDown(k int) error {

	for mp.leftChild(k) < mp.data.GetSize() {
		j := mp.leftChild(k)
		jp_v, err := mp.data.Get(j + 1)

		if err != nil {
			return err
		}
		j_v, err := mp.data.Get(j)
		if err != nil {
			return err
		}
		//  data[j] 是 leftChild 和 rightChild 中的最大值
		if j+1 < mp.data.GetSize() && compare(jp_v, j_v) > 0 {
			j++
		}
		k_v, err := mp.data.Get(k)
		j_v, err = mp.data.Get(j)

		if compare(k_v, j_v) >= 0 {
			break
		}
		mp.data.swap(k, j)
		k = j

	}
	return nil
}

func (mp *MaxHeap) Add(e interface{}) error {

	if err := mp.data.AddLast(e); err != nil {
		return err
	}
	if err := mp.shiftUp(mp.data.GetSize() - 1); err != nil {
		return err
	}
	return nil
}

func (mp *MaxHeap) Pop() error {


	if err := mp.shiftUp(mp.data.GetSize() - 1); err != nil {
		return err
	}
	return nil
}

func (mp *MaxHeap) ToString() string {

	return mp.data.ToString()
}
