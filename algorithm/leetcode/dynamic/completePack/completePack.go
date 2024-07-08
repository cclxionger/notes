package dynamic

type Items struct {
	Weight int
	Val    int
}

func SelectDouble2(items []Items, volume int) int { //二维dp
	n := len(items)
	//初始化dp，i代表第i个物品，j代表容量为j的时候，dp代表第i个物品，容量为j的时候选取的最大价值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, volume+1)
	}
	//初始化第一行
	for j := 1; j <= volume; j++ {
		if j >= items[0].Weight {
			dp[0][j] = items[0].Val + dp[0][j-items[0].Weight]
		}
	}
	//dp[i][j]依赖于当前行的前面值和dp[i-1][j] ，不依赖左上值
	for i := 1; i < n; i++ {
		curWegith := items[i].Weight //第i个物品的重量
		curVal := items[i].Val       //第i个物品的价值
		for j := 1; j <= volume; j++ {
			if j >= curWegith {
				dp[i][j] = max(dp[i-1][j], curVal+dp[i][j-curWegith])
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}

	}
	// fmt.Println(dp)
	return dp[n-1][volume]

}
