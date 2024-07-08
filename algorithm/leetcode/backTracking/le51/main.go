package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	// 原本棋盘是二维数组，可以把符合条件的结果添加到res三维切片里面
	// 由于leetcode要返回二位切片，那么需要把每一次符合条件的二维数组结果转变成一维数组
	chessboard := make([][]string, n)
	// 初始化chessboard棋盘
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	dfs(n, &res, 0, chessboard)
	return res
}
func dfs(n int, res *[][]string, row int, chessboard [][]string) {
	if row == n {
		temp := make([]string, n)
		// 自定义copy，把二维数组copy到一维数组里面
		for rowIndex, rowVal := range chessboard {
			temp[rowIndex] = strings.Join(rowVal, "")
		}
		*res = append(*res, temp)
		return
	}

	for j := 0; j < n; j++ {
		if isVaild(chessboard) {
			chessboard[row][j] = "Q"
			dfs(n, res, row+1, chessboard)
			chessboard[row][j] = "."
		}
	}
}
func isVaild(chessboard [][]string) bool {
	return true
}

func main() {
	fmt.Println(solveNQueens(2))
	
	// [[Q. ..][.Q ..][.. Q.][...Q]]
}
