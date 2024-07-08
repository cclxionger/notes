package main

import "fmt"

// n为1-n，k为选取k个数字
func combine(n int, k int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	dfs(n, k, path, &res, 1)
	return res
}
func dfs(n int, k int, path []int, res *[][]int, startIndex int) {
	if len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := startIndex; i <= n; i++ {
		// 进行剪枝
		// 需要的数字个数为 k-len(path) 剩余个数为 n-i
		// 假如5(n)选3(k)，那么到4的时候就不能要了，3是刚好，
		// 当i为3的时候，3-0=3  5-3=2 
		// 当i为4的时候，4-0=4  5-3=2 退出循环
		// 假如6(n)选3(k)，那么到5的时候就不能要了，4是刚好
		// 当i为4的时候，3-0=3  6-4=2
		// 当i为5的时候，3-0=3  6-5=1
		if k-len(path) > n-i + 1  {
			break
		}
		path = append(path, i)
		dfs(n, k, path, res, i+1)
		path = path[:len(path)-1]
	}
}
func main() {
	a := combine(3, 2)
	fmt.Println(a)
}
