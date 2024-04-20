package main
/*
author:常成琳
*/

import (
	"fmt"
	"sort"
)

var resultIndex []int //声明全局变量resultIndex 切片
// Item 结构体表示每个物品的信息
type Item struct {
	Index   int     // 物品的编号
	Weight  int     // 物品的重量
	Value   int     // 物品的价值
	UnitVal float64 // 物品的单位价值（价值/重量）
}

// UnitValue 类型为Item切片，用于按单位价值排序
type UnitValue []Item

// 用sort方法必须要实现这三个接口
// Len 方法返回切片中元素的数量
func (b UnitValue) Len() int {
	return len(b)
}

// Swap 方法交换切片中两个元素的位置
func (b UnitValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less 方法比较两个元素的单位价值，用于排序
func (b UnitValue) Less(i, j int) bool { //如果后面的元素单位价值大，那就交换
	return b[i].UnitVal < b[j].UnitVal
}

// calculateUnitValue 函数计算物品的单位价值
func calculateUnitValue(item Item) float64 {
	return float64(item.Value) / float64(item.Weight)
}

// selectItems 函数选择总重量不超过给定值的物品，返回最大价值
func selectItems(items []Item, totalWeight int) float64 {

	// 按单位价值降序排序物品
	sort.Sort(sort.Reverse(UnitValue(items)))

	maxValue := 0.0 // 最大价值
	for _, item := range items {
		if totalWeight >= item.Weight {
			// 如果当前物品可以完全装入背包
			totalWeight -= item.Weight
			maxValue += (float64)(item.Value)
			resultIndex = append(resultIndex, item.Index)
			// fmt.Println(item.Index)
		} else {
			// 如果当前物品不能完全装入背包，只取部分
			fractionalValue := float64(totalWeight) * calculateUnitValue(item)
			maxValue += fractionalValue
			break
		}
	}
	return maxValue
}

func main() {
	// 初始化物品列表
	items := []Item{
		{Index: 0, Weight: 2, Value: 10},
		{Index: 1, Weight: 3, Value: 5},
		{Index: 2, Weight: 5, Value: 15},
		{Index: 3, Weight: 7, Value: 7},
		{Index: 4, Weight: 1, Value: 6},
		{Index: 5, Weight: 4, Value: 18},
		{Index: 6, Weight: 1, Value: 3},
	}

	// 设置每个物品的单位价值
	for i := range items {
		items[i].UnitVal = calculateUnitValue(items[i])
	}
	// 选择总重量不超过15的物品，并计算最大价值
	maxValue := selectItems(items, 15)
	fmt.Printf("最大价值是: %f\n", maxValue)
	fmt.Printf("选取的物品的索引有%d\n", resultIndex)
}
