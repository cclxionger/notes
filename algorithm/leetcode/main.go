package main

import (
	"fmt"
	monotonestack "towpointer/monotoneStack"
)

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	// result := make([]int,len(nums),len(nums))
	// result = towpointer.SortedSquares(nums)
	// fmt.Println(result)
	// arr2 := []int {2,3,1,2,4,3}
	// n := towpointer.MinSubArrayLen(7,arr2)
	// fmt.Println(n)
	// n := 2
	// result := dynamic.ClimbStairs(n)
	// fmt.Println(result)
	// a := dynamic.UniquePaths2(1, 3)
	// fmt.Println(a)
	// obstacleGrid := [][]int{{0, 0}, {1,1},{0,0,}}
	// a := dynamic.UniquePathsWithObstacles2(obstacleGrid)
	// fmt.Println(a)
	temperatures := []int{78, 74, 75, 71, 69, 72, 76}
	a := monotonestack.DailyTemperatures2(temperatures)
	fmt.Println(a)

}
