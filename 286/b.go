package main

import (
	"fmt"
	"strings"
)

func replaceNaWithNya(S string) string {
	return strings.ReplaceAll(S, "na", "nya")
}

func main() {
	var N int
	var S string

	// 入力を受け取る
	fmt.Scan(&N)
	fmt.Scan(&S)

	// "na" を "nya" に置き換えて出力する
	result := replaceNaWithNya(S)
	fmt.Println(result)
}
