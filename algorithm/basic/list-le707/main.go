package main

import "fmt"

// 有哨兵的单链表
type List struct {
	next *List
	val  int
}
type MyLinkedList struct {
	dummyHead *List // 虚拟头节点
	size      int   // 链表大小
}

func Constructor() *MyLinkedList {
	sentinel := &List{nil, -1}        // 创建一个空的 List 节点作为虚拟头节点
	return &MyLinkedList{sentinel, 0} // 返回一个指向 MyLinkedList 的指针，并初始化 size 为 0
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.dummyHead.next
	for index > 0 {
		cur = cur.next
		index--
	}
	// 此时 cur 指向目标节点
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	sentinel := this.dummyHead
	newNode := &List{nil, val}
	// 如果是空链表
	if sentinel.next == nil {
		sentinel.next = newNode
	} else {
		// 如果不是空链表
		cur := sentinel.next
		newNode.next = cur
		sentinel.next = newNode
	}
	this.size++

}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}
func main() {
	list := Constructor()
	list.AddAtHead(2)
	list.AddAtHead(2)
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
