package _8_Heap_and_Priority_Queue

import (
	"testing"
)

func TestMaxHeap_Add(t *testing.T) {

	//a := []string{"1", "2", "3", "4", "45", "6"}
	a := []string{"1"}
	maxHeap := NewMaxHeap(20)
	for i := 0; i < len(a); i++ {
		v := a[i]
		println(v)
		maxHeap.Add(v)
	}
	println(maxHeap.ToString())
}
