package color

import "fmt"

func Forecolor(color, data string) string {
	switch color {
	// 如果字体也为黑色，那么啥也看不见哈哈哈哈
	case "黑色":
		return fmt.Sprintf("\033[30m %s \033[0m", data)
	case "红色":
		return fmt.Sprintf("\033[31m %s \033[0m", data)
	case "绿色":
		return fmt.Sprintf("\033[32m %s \033[0m", data)
	case "黄色":
		return fmt.Sprintf("\033[33m %s \033[0m", data)
	}
	
	return fmt.Sprintf("\033[30m %s \033[0m", data)
}

func VariousColor() {
	// 前景色设置，背景为黑色
	fmt.Println("\033[30m 黑色 \033[0m")
	fmt.Println("\033[31m 红色 \033[0m")
	fmt.Println("\033[32m 绿色 \033[0m")
	fmt.Println("\033[33m 黄色 \033[0m")
	fmt.Println("\033[34m 蓝色 \033[0m")
	fmt.Println("\033[35m 紫色 \033[0m")
	fmt.Println("\033[36m 青色 \033[0m")
	fmt.Println("\033[37m 灰色 \033[0m")
	// 背景色设置，字体为白色
	fmt.Println("\033[40m 黑色 \033[0m")
	fmt.Println("\033[41m 红色 \033[0m")
	fmt.Println("\033[42m 绿色 \033[0m")
	fmt.Println("\033[43m 黄色 \033[0m")
	fmt.Println("\033[44m 蓝色 \033[0m")
	fmt.Println("\033[45m 紫色 \033[0m")
	fmt.Println("\033[46m 青色 \033[0m")
	fmt.Println("\033[47m 灰色 \033[0m")
}
