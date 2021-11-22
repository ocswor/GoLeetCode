package _8_Heap_and_Priority_Queue

import (
	"errors"
	"fmt"
	"strings"
)

type Array struct {
	data []interface{}
	size int
}

func NewArray(capacity int) Array {
	return Array{
		data: make([]interface{}, capacity),
	}
}

/*
获取数组的容量
*/
func (arr *Array) getCapacity() int {
	return len(arr.data)
}

/*
获取数组中的元素个数
*/
func (arr *Array) GetSize() int {
	return arr.size
}

/*
判断数组是否为空
*/
func (arr *Array) isEmpty() bool {
	return arr.size == 0
}

func (arr *Array) resize(capacity int) {
	newData := make([]interface{}, capacity)
	for i := 0; i < arr.size; i++ {
		newData[i] = arr.data[i]
	}
	arr.data = newData

}

/*
向数组中index位置 插入一个元素
*/
func (arr *Array) Add(index int, e interface{}) error {
	if index < 0 || index > arr.size {
		return errors.New("Add failed. Require index >= 0 and index <= size.")
	}
	if arr.size == len(arr.data) {
		arr.resize(2 * len(arr.data))
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i] // index 位置后面的元素后移
	}
	arr.data[index] = e
	arr.size++
	return nil
}

/*
数组末尾增加一个元素
*/
func (arr *Array) AddLast(e interface{}) error {

	return arr.Add(arr.size, e)
}

/*
数组首位 添加一个元素
*/
func (arr *Array) AddFirst(e interface{}) error {

	return arr.Add(0, e)
}

/*
数组末尾 添加一个元素
*/
func (arr *Array) Append(e interface{}) error {

	return arr.Add(arr.size, e)
}

/*
获取index 位置的元素
*/
func (arr *Array) Get(index int) (interface{}, error) {
	if index < 0 || index >= arr.size {
		return nil, errors.New("Get failed. Index is illegal.")
	}
	return arr.data[index], nil
}

/*
修改索引位置的元素
*/
func (arr *Array) Set(index int, e interface{}) error {
	if index < 0 || index >= arr.size {
		return errors.New("Set failed. Index is illegal.")
	}
	arr.data[index] = e
	return nil
}

/*
查找数组中是否有元素e
*/
func (arr *Array) Contains(e interface{}) bool {

	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

/*
查找数组中是否有元素e 如果不存在元素e，则返回-1
这里 == 还需要验证一下
*/
func (arr *Array) Find(e interface{}) int {

	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

/*
删除 index 位置的元素
*/
func (arr *Array) Remove(index int) (interface{}, error) {
	if index < 0 || index >= arr.size {
		return nil, errors.New("Remove failed. Index is illegal.")
	}
	e := arr.data[index]
	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	arr.data[arr.size] = nil //原有的最后一个元素置空

	// 缩减 和 扩展的规则 可以不一样
	if arr.size == len(arr.data)/4 && len(arr.data)/2 != 0 {
		arr.resize(len(arr.data) / 2)
	}
	return e, nil
}

/*
删除数组第一个元素
*/
func (arr *Array) RemoveFirst() (interface{}, error) {

	return arr.Remove(0)
}

/*
删除数组最后一个元素
*/
func (arr *Array) RemoveLast() (interface{}, error) {

	return arr.Remove(arr.size - 1)
}

/*
删除数组元素 e
*/

func (arr *Array) RemoveElement(e interface{}) (interface{}, error) {
	index := arr.Find(e)
	if index != -1 {
		return arr.Remove(index)
	}
	return nil, errors.New("not find element in arr")

}

/*
交换数组中的两个元素
*/
func (arr *Array) swap(i int, j int) error {

	if i < 0 || i >= arr.size || j < 0 || j >= arr.size {
		return errors.New("Remove failed. Index is illegal.")
	}
	t := arr.data[i]
	arr.data[i] = arr.data[j]
	arr.data[j] = t
	return nil
}

func (arr *Array) ToString() string {
	var b strings.Builder
	rs := fmt.Sprintf("Array: size = %d , capacity = %d \n", arr.size, len(arr.data))
	b.WriteString(rs)
	b.WriteString("[")
	for i := 0; i < arr.size; i++ {
		b.WriteString(arr.data[i].(string))
		if i != arr.size-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("] ")
	return b.String()
}
