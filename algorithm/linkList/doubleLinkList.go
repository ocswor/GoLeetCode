package linkList

import "fmt"

// 节点数据
type NodeValue interface{}

// 双向链表节点
type DoubleNode struct {
	Key   NodeValue
	Value NodeValue
	Pre   *DoubleNode
	Next  *DoubleNode
}

// 双链表
type DoubleList struct {
	Size uint
	Head *DoubleNode
	Tail *DoubleNode
}

// 双链表初始化
func NewDoubleList() *DoubleList {

	return &DoubleList{}
}

func (list *DoubleList) Append(node *DoubleNode) bool {

	if node == nil {
		return false
	}
	if list.Size == 0 {
		list.Tail = node
		list.Head = node
		node.Pre = nil
		node.Next = nil
	} else {
		node.Pre = list.Tail
		node.Next = nil
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size++
	return true
}

func (list *DoubleList) Get(index uint) *DoubleNode {

	if list.Size == 0 || index > list.Size-1 {
		return nil
	}
	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}
	return node
}

func (list *DoubleList) Insert(index uint, node *DoubleNode) bool {
	if index > list.Size || node == nil {
		return false
	}
	if index == list.Size {
		return list.Append(node)
	}

	if index == 0 {
		//node.Next = list.Head
		//list.Head = node
		//list.Head.Pre = nil

		list.Head.Pre = node
		node.Next = list.Head
		list.Head = node
		node.Pre = nil

		list.Size++
		return true
	}

	nextNode := list.Get(index) // 插入位置的下一个节点
	preNode := nextNode.Pre     // 插入位置的上一个节点

	// 主要影响三个节点
	//处理上一个节点
	preNode.Next = node
	// 处理当前节点
	node.Next = nextNode
	node.Pre = preNode
	// 处理 下一个节点
	nextNode.Pre = node

	list.Size++
	return true
}

func (list *DoubleList) Delete(index uint) bool {

	if index > list.Size-1 {
		return false
	}

	if index == 0 {
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Pre = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	if index == list.Size-1 {
		list.Tail.Pre.Next = nil
		list.Tail = list.Tail.Pre
		list.Size--
		return true
	}

	node := list.Get(index)

	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	list.Size--
	return true

}

func (list *DoubleList) DeleteNode(node *DoubleNode) bool {

	if list.Head == node {
		// 头结点
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Pre = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	if list.Tail == node {
		// 尾节点
		list.Tail.Pre.Next = nil
		list.Tail = list.Tail.Pre
		list.Size--
		return true
	}
	// 中间节点
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	list.Size--
	return true
}

func (list *DoubleList) MoveToHead(node *DoubleNode) bool {
	list.DeleteNode(node)
	list.Insert(0, node)
	return true
}

/*
正序 打印链表
*/
func (list *DoubleList) Display() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is null or empty")
		return
	}
	ptr := list.Head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Value)
		ptr = ptr.Next
	}
}

/*
倒序 打印链表
*/
func (list *DoubleList) Reverse() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is null or empty")
		return
	}
	ptr := list.Tail
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Value)
		ptr = ptr.Pre
	}
}
