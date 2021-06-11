package lru

import (
	"GoLeetCode/algorithm/linkList"
)

type LRUCache struct {
	capacity   uint
	doubleList *linkList.DoubleList
	nodeTable  map[string]*linkList.DoubleNode
}

func NewLRUCache(capacity uint) *LRUCache {

	return &LRUCache{
		capacity:   capacity,
		doubleList: linkList.NewDoubleList(),
		nodeTable:  map[string]*linkList.DoubleNode{},
	}
}

func (cache *LRUCache) Get(key string) interface{} {

	node, ok := cache.nodeTable[key]
	if !ok {
		return nil
	}
	//
	////如果存在，则需要移动Node节点到表头
	cache.doubleList.MoveToHead(node)
	return node.Value
}

func (c *LRUCache) Put(key string, value interface{}) {
	node, ok := c.nodeTable[key]

	if !ok {
		//如果Node不在表中，代表缓存中并没有
		//fmt.Println(len(c.nodeTable))
		if uint(len(c.nodeTable)) == c.capacity {
			//超过容量了 ,首先移除尾部的节点
			delete(c.nodeTable, c.doubleList.Tail.Key.(string))
			c.doubleList.Delete((c.capacity - 1))
		}
		node = &linkList.DoubleNode{
			Value: value,
			Key:   key,
		}
	}
	//如果存在，则需要移动Node节点到表头
	c.doubleList.Insert(0, node)
	c.nodeTable[key] = node
}
