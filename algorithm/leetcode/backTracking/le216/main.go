package main

import "fmt"

//找出所有相加之和为 n 的 k 个数的组
func combinationSum3(k int, n int) [][]int {
	var res = make([][]int, 0)
	var path = make([]int, 0)
	dfs3(k, n, 0, 1, path, &res)
	return res
}
func dfs3(k, n, sum, startIndex int, path []int, res *[][]int) {
	if sum > n {
		return
	}
	if sum == n && k == len(path) {
		temp := make([]int, k)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := startIndex; i <= 9; i++ {
		//进行剪枝操作，k-len(path)是还需要的数字个数，比如选3个，最多到7开始，9-need+2 = 8就不行了
		// 9-0+2
		if i >= 9-(k-len(path))+2 || sum+i > n {
			break
		}
		path = append(path, i)
		sum += i
		dfs3(k, n, sum, i+1, path, res)
		path = path[:len(path)-1]
		sum -= i
	}
}
func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(1, 1))
}
