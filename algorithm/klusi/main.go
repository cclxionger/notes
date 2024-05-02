package main  
  
import (  
	"fmt"  
	"sort"  
)  
  
// Edge 表示图的一条边  
type Edge struct {  
	u, v, weight int  
}  

  
// ByWeight 实现了 sort.Interface，用于对 Edge 切片按权重排序  
type ByWeight []Edge  
  
func (a ByWeight) Len() int           { return len(a) }  
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }  
func (a ByWeight) Less(i, j int) bool { return a[i].weight < a[j].weight }  
  
// findSet 查找元素 x 所属的集合编号  
func findSet(parent []int, x int) int {  
	if parent[x] != x {  
		parent[x] = findSet(parent, parent[x])  
	}  
	return parent[x]  
}  
  
// union 将集合 x 和集合 y 合并  
func union(parent []int, x, y int) {  
	parent[findSet(parent, x)] = findSet(parent, y)  
}  
  
// kruskal 实现了克鲁斯卡尔算法  
func kruskal(n int, edges []Edge) []Edge {  
	// 初始化并查集  
	parent := make([]int, n)  
	for i := 0; i < n; i++ {  
		parent[i] = i  
	}  
  
	// 按照权重对边进行排序  
	sort.Sort(ByWeight(edges))  
  
	// 初始化最小生成树的边  
	mstEdges := []Edge{}  
  
	// 遍历排序后的边  
	for _, edge := range edges {  
		u, v := edge.u, edge.v  
		// 如果边 (u, v) 的两个端点不在同一个集合中，则加入最小生成树  
		if findSet(parent, u) != findSet(parent, v) {  
			mstEdges = append(mstEdges, edge)  
			union(parent, u, v) // 合并集合  
		}  
	}  
  
	return mstEdges  
}  
  
func main() {  
	// 示例图的边和节点数量  
	n := 5 // 节点数量  
	edges := []Edge{  
		{0, 1, 7},  
		{0, 2, 9},  
		{0, 3, 14},  
		{1, 2, 10},  
		{1, 3, 15},  
		{2, 3, 11},  
		{2, 4, 2},  
		{3, 4, 6},  
		{4, 0, 9},  
	}  
  
	// 调用克鲁斯卡尔算法  
	mstEdges := kruskal(n, edges)  
  
	// 打印最小生成树的边  
	fmt.Println("最小生成树的边:")  
	for _, edge := range mstEdges {  
		fmt.Printf("(%d, %d) - %d\n", edge.u, edge.v, edge.weight)  
	}  
}