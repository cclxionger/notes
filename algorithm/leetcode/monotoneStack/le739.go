package monotonestack

import (
	"fmt"
	"sync"
)

/*
leet739思路：
每次都要向右寻找比他第一个大的元素，单调递增栈（因为小的进去了，大的就出来了）
栈里面村的是temperatures数组的索引，假设说当前索引是i，然后每次i对应的值先与栈顶元素比较，
大的话就一直（for）把最顶元素pop掉，直到遇到小的话就push i，
*/
// 不使用栈，而是用切片代替栈
func DailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		//stack[len(stack) - 1]相当于栈顶

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1] //必须要在里面，否则每次push的时候，栈顶也改变了
			result[top] = i - top
			stack = stack[:len(stack)-1] //pop
		}
		stack = append(stack, i) //push
	}
	return result

}

// 使用栈
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := NewStack()
	stack.Push(0)
	for i := 1; i < n; i++ {
		//要保证stack.Length() > 0否则会报错,因为没有元素了，还stack.Peek了
		for stack.Length() > 0 && temperatures[i] > temperatures[stack.Peek()] {

			// 弹出栈顶元素，并计算等待天数
			index := stack.Pop()
			result[index] = i - index

		}
		stack.Push(i)

	}
	// for stack.Length() > 0 {
	// 	index := stack.Pop()
	// 	result[index] = 0
	// }
	return result
}

type Item = int

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
		return -1
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
