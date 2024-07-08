package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	sum := 0
	dfs(candidates, target, path, &res, 0, sum)
	return res
}
func dfs(candidates []int, target int, path []int, res *[][]int, startIndex int, sum int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		// 剪枝
		if sum > target {
			return
		}
		sum += candidates[i]
		path = append(path, candidates[i])
		// 元素可以重复，所以还是从i开始继续往下树形结构
		dfs(candidates, target, path, res, i, sum)
		path = path[:len(path)-1]
		sum -= candidates[i]
	}
}
func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
