package dynamic

type Items struct {
	Weight int
	Val    int
}

func SelectDouble(items []Items, volume int) int { //二维dp
	n := len(items)
	//初始化dp，i代表第i个物品，j代表容量为j的时候，dp代表第i个物品，容量为j的时候选取的最大价值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, volume+1)
	}
	//初始化第一行
	for j := 1; j <= volume; j++ {
		if j >= items[0].Weight {
			dp[0][j] = items[0].Val
		}
	}
	//dp[i][j]取决与上一行左上角和dp[i-1][j]的值
	//所以是i和j是可以相反的
	for i := 1; i < n; i++ {
		curWegith := items[i].Weight //第i个物品的重量
		curVal := items[i].Val       //第i个物品的价值
		//j可以正序遍历，也可以倒叙遍历
		for j := 0; j <= volume; j++ {
			// for j := volume; j >= 0; j-- {
			if j >= curWegith {
				dp[i][j] = max(curVal+dp[i-1][j-curWegith], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	// fmt.Println(dp)

	return dp[n-1][volume]

}

func SelectSingle(items []Items, volume int) int {
	n := len(items)
	dp := make([]int, volume+1)
	for i := 0; i < n; i++ {
		curWegith := items[i].Weight //第i个物品的重量
		curVal := items[i].Val       //第i个物品的价值
		//为什么j倒叙遍历呢
		/*
			如果正向遍历，dp[j-curWegith]受到前面第i个物品dp【j】的影响
			如果倒叙遍历，dp[j-curWegith]收到的是第i-1个物品dp【j】的影响，这个和二维数组是一样的
		*/
		//那为什么二维数组j可以正序遍历呢(当然也可以倒叙)
		/*
			因为二维数组j受到的是第i-1个物品dp[i-1][j-curWeight]的影响
			与第i行前面元素无关
		*/
		for j := volume; j >= 0; j-- {
			if j >= curWegith {
				dp[j] = max(dp[j], curVal+dp[j-curWegith])
			}
		}
		// fmt.Println(dp)

	}
	return dp[volume]
}
