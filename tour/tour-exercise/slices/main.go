package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 定义图片列为dy
	pic := make([][]uint8, dy)
	// 遍历每一行，把图片每一行都初始化
	for i := range pic {
		// 图片宽度为dx
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			// 使用 (x+y)/2 作为灰度值
			pic[i][j] = uint8((i + j) / 2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
	// 稍后创建文件把图片展示出来
}
