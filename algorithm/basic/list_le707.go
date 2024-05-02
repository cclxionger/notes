package basic

//双链表
type MyLinkedList struct {
	dummy *Node //哨兵
}
type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{rear}
}

func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	for head != this.dummy && index > 0 {
		index--
		head = head.Next
	}
	if index != 0 {
		return -1
	}
	return head.Val

}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy.Next, //这俩步定义的时候就已经增加了一条线
		Pre:  dummy,
	}
	dummy.Next.Pre = node //不能相反
	dummy.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	cur := dummy
	for cur.Next != dummy {
		cur = cur.Next
	} //遍历
	node := &Node{
		Val:  val,
		Pre:  cur,
		Next: dummy,
	}
	cur.Next = node  // 当前尾节点的 Next 指向新节点
	dummy.Pre = node // 哨兵节点的前一个节点指向新节点

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	dummy := this.dummy
	cur := dummy
	for index > 0 && cur.Next != dummy {
		index--
		cur = cur.Next
	}
	if index < 0 {
		return
	}
	node := &Node{
		Val:  val,
		Pre:  cur,
		Next: cur.Next,
	}
	if cur.Next != nil {
		cur.Next.Pre = node
	}
	// cur.Next.Pre = node
	cur.Next = node
	if cur != dummy {
		node.Pre = cur
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	dummy := this.dummy
	cur := dummy
	for cur.Next != dummy && index > 0 { //排序index 是负数的情况
		index--
		cur = cur.Next
	}
	if index != 0 { //排序index超级大的情况
		return
	}
	toDelete := cur.Next
	cur.Next = toDelete.Next
	toDelete.Next.Pre = cur
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
