package main

import (
	"fmt"
)

// Graph 表示一个无向图，使用邻接表的形式表示
type Graph struct {
	vertices int     // 顶点数量
	colors   []int   // 每个顶点的颜色
	adjList  [][]int // 邻接表
}

// NewGraph 创建一个新的图
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		colors:   make([]int, vertices),
		adjList:  make([][]int, vertices),
	}
}

// AddEdge 在图中添加一条无向边
func (g *Graph) AddEdge(u, v int) { // 无向图，添加双向边
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

// GreedyColoring 使用贪心算法为图着色
func (g *Graph) GreedyColoring() {
	maxColor := 1 // 当前最大颜色，从1开始编号
	for v := 0; v < g.vertices; v++ {
		if g.colors[v] == 0 { // 如果顶点v尚未着色
			usedColors := make(map[int]bool)        // 用于记录已使用的颜色
			for _, neighbor := range g.adjList[v] { //在邻居中查找
				if g.colors[neighbor] != 0 {
					usedColors[g.colors[neighbor]] = true // 标记邻居的颜色为已使用
				}
			}
			for color := 1; ; color++ {
				if usedColors[color] != true { // 找到第一个未使用的颜色
					g.colors[v] = color
					if color > maxColor {
						maxColor = color // 更新最大颜色
					}
					break
				}
			}
			
		}
	}
}

// PrintColoring 打印图的着色结果
func (g *Graph) PrintColoring() {
	for i, color := range g.colors {
		fmt.Printf("顶点 %d 的颜色: %d\n", i, color)
	}
}

func main() {
	// 创建一个包含5个顶点的图
	graph := NewGraph(5)

	// 添加边
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	// 使用贪心算法为图着色
	graph.GreedyColoring()

	// 打印着色结果
	graph.PrintColoring()
	for i := 0; i < len(graph.adjList); i++ {
		for j := 0; j < len(graph.adjList[i]); j++ {
			fmt.Print(graph.adjList[i][j])
			fmt.Print("\t")
		}
		fmt.Println()
	}
}
