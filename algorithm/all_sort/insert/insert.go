package insert

func InsertSort1(n int, arr []int) { //第一个元素是哨兵
	for i := 2; i <= n; i++ {
		arr[0] = arr[i]              //把当前要插入的元素赋给哨兵
		for j := i - 1; j > 0; j-- { //j最小为1

			if arr[j] > arr[0] { //要插入的元素大于哨兵
				arr[j+1] = arr[j]
			} else {
				arr[j+1] = arr[0]
				break
			}
			if j == 1 {
				arr[j] = arr[0]
				break
			}
		}

	}
}

func InsertSort2(n int, arr []int) {
	for i := 1; i < n; i++ {
		curInd := i
		for j := i - 1; j >= 0; {
			preInd := j          //i前面的索引
			if arr[j] < arr[i] { //如果要后移
				arr[preInd+1] = arr[preInd]
			} else { //不后移的话增加了直接退出第二层循环
				arr[j] = arr[curInd]
				break
			}
			j--
			if j == -1 {
				arr[0] = arr[curInd]
				break
			}
		}

	}
}
