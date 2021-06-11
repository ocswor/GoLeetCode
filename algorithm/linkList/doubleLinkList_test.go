package linkList

import (
	"fmt"
	"testing"
)


//双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向直接后继和直接前驱。
//所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前驱结点和后继结点。

func TestDoubleList_Append(t *testing.T) {

	list := NewDoubleList()

	for i := 0; i < 100; i++ {
		node := DoubleNode{
			Value: i,
		}
		list.Append(&node)
	}

	list.Display()
	list.Reverse()
}

func TestDoubleList_Insert(t *testing.T) {

	list := NewDoubleList()

	for i := 0; i < 100; i++ {
		node := DoubleNode{
			Value: i,
		}
		list.Append(&node)
	}
	list.Insert(89, &DoubleNode{Value: 201})
	list.Display()

}

func TestDoubleList_Delete(t *testing.T) {
	list := NewDoubleList()

	for i := 0; i < 100; i++ {
		node := DoubleNode{
			Value: i,
		}
		list.Append(&node)
	}
	list.Delete(89)
	rs := list.Delete(98)
	r := list.Delete(0)
	fmt.Println(rs)
	fmt.Println(r)
	list.Display()
	fmt.Println(list.Size)
}


func TestDoubleList_MoveToHead(t *testing.T) {
	list := NewDoubleList()

	for i := 0; i < 100; i++ {
		node := DoubleNode{
			Value: i,
		}
		list.Append(&node)
	}
	node := list.Get(98)
	list.MoveToHead(node)
	list.Display()
}