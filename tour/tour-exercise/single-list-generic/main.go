package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type T interface{}
type List struct {
	next *List
	val  T
}

// type sentinel List
type myList struct {
	dummyHead *List // 虚拟头节点
	size      int   // 链表大小
}

func NewList() myList {
	// 忽略头节点的val
	dummyHead := &List{next: nil}
	return myList{dummyHead, 0}
}
func (l myList) GetAllList() []T {
	cur := l.dummyHead.next // 第一个节点
	if cur == nil {
		fmt.Println("空链表")
		return nil
	}
	arr := make([]T, 0)
	for {
		arr = append(arr, cur.val)
		cur = cur.next
		if cur == nil {
			break
		}

	}
	return arr
}

func (this myList) addTail(val T) {
	// 想要在尾部插入元素，那么需要查找到最后一个不是nil的节点
	pre := this.dummyHead //哨兵节点
	node := &List{nil, val}
	// 如果是空链表
	if pre.next == nil {
		pre.next = node
		return
	}
	// 如果不是空链表,需要得到nil的前一个节点（也就是最后一个节点）
	for {
		if pre.next == nil {
			pre.next = node
			break
		}
		pre = pre.next
	}

}

func main() {
	list := NewList()
	list.addTail(6)
	list.addTail("lc")
	fmt.Println(list.GetAllList())
}
