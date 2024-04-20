package main

import (
	"fmt"
	"sort"
)

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

type ReverseSortBy struct{ SortBy }

func (r ReverseSortBy) Less(i, j int) bool {
	return r.SortBy.Less(j, i)
}
func main() {

	a := []int{1, 5, 7, 6, 2}
	fmt.Println(a)
	b := make([]float64, len(a)) //想把int类型的切片a转换成float类型的切片b
	for i := 0; i < len(a); i++ {
		b[i] = float64(a[i])
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(b)))
	fmt.Println(b)
	sort.Sort(SortBy(a)) //SortBy(a) 将 a 转换为 SortBy 类型。
	//因为 SortBy 是 []int 的别名，所以这种转换是零成本的，只是告诉编译器你现在要将这个切片当作 SortBy 类型来处理。
	fmt.Println(a)

	sort.Sort(ReverseSortBy{SortBy(a)}) //将a切片再次反转过来
	fmt.Println(a)
	fmt.Println()
	d := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为 0的元素
	d = append(d[:0], d[1:]...) //删除掉第一个元素
	fmt.Println(d)
	//31 32 33 34 35 36 37
	//删除掉索引为2的元素 就是删除掉33
	d = append(d[:2], d[3:]...)
	fmt.Println(d)
	fmt.Println()
}
