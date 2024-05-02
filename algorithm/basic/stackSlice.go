package basic

import (
	"fmt"
	"sync"
)
//go中没有栈
// 利用slice进行实现栈
// all types
type Item interface{}

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// New
func NewStack() *ItemStack {
	s := &ItemStack{}
	s.items = []Item{} //初始化了，Item不再是nil了
	return s
}

// 重写String
func (stack *ItemStack) String() string {
	return fmt.Sprintf("%+v", stack.items) //%+v  还是%#v看情况
}

// Push
func (s *ItemStack) Push(it Item) {
	s.lock.Lock()
	s.items = append(s.items, it)
	s.lock.Unlock()
}

// Pop
func (s *ItemStack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1] //删除掉最后一个元素
	return item                           //返回删除的元素
}

// peek (返回栈顶元素)
func (s *ItemStack) Peek() Item {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.items[len(s.items)-1]
}

func (s *ItemStack) Length() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items)
}
