package _8_Heap_and_Priority_Queue

import "testing"

func TestArray(t *testing.T) {

	arr := NewArray(5)
	arr.Append("1")
	arr.Append("2")
	arr.Append("3")
	arr.Add(0,"4")
	arr.Add(1,"5")
	arr.Add(1,"6")
	arr.Remove(1)
	arr.RemoveElement("1")

	println(arr.ToString())
}