package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	// 定义一个非常小的数字
	const m = 1e-10
	// 假设z初值为1 或者为 x 或者为 x / 2
	z := 1.0
	// 迭代10次
	for i := 1; i <= 10; i++ {
		fmt.Printf("第%d次迭代结果为%f\n", i, z)
		// 当z²-x 比 m 还小就可以退出了
		// 判断退出条件需要+一个绝对值，因为不知道x比1大还是比1小
		if math.Abs(z*z-x) < m {
			fmt.Printf("第%d次退出了\n", i)
			break
		}
		z -= (z*z - x) / (2 * z)
	}
	// 如果10次还没有得到结果则继续
	for math.Abs(z*z-x) > m {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintf("%f is not a postive number", e)
	// return fmt.Sprintf("%v", e) will cause recursive(递归)
	// change float64(e)  can solve this problem
	return fmt.Sprintf("%v is not a postive number", float64(e))
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2.1))
}
