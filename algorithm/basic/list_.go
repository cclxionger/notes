package basic

//go中有list.List，但是自己练习模仿一个 在官网example中
//这个List 是可以头插入，可以尾插入的单链表
type element[T any] struct {
	next *element[T]
	val  T
}
type List[T any] struct {
	head, tail *element[T]
}

func (lt *List[T]) Push(v T){
	if lt.tail == nil {
		lt.head = &element[T]{val: v}
		lt.tail = lt.head
	} else {
		lt.tail.next = &element[T]{val: v}
		lt.tail = lt.tail.next
	}
}

func (lt List[T]) GetAll() []T{
	var elems []T
	for e := lt.head; e != nil; e = e.next {
		elems = append(elems, e.val)
		
	}
	return elems
}
