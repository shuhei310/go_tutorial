package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// case文
	switch time.Saturday {
	// 実行日が土曜日
	case today + 0:
		fmt.Println("Today.")
	// 実行日が金曜日
	case today + 1:
		fmt.Println("Tomorrow.")
	// 実行日が木曜日
	case today + 2:
		fmt.Println("In two days.")
	// それ以外
	default:
		fmt.Println("Too far away.")
	}
}