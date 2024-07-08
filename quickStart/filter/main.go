package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		inStr := in.Text()
		if inStr == "" {
			fmt.Println("exit")
			break
		}
		str := strings.ToUpper(inStr)
		fmt.Println(str)
	}
	if err := in.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
