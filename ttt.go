package main

import (
	"RevStrPkg"
	"fmt"
	"strings"
)

func main() {
	ii := strings.Split("abcicba", "")
	fmt.Println(ii)
	bb := make([]string, 12)
	copy(bb, ii)
	fmt.Println(bb)
	// len_ii := len(ii)
	// ii = append(ii[:len_ii/2-1], ii[1:], "asdfasdf")
	// fmt.Println(ii)
	// fmt.Printf("%q\n", ii)
	// fmt.Printf("%T\n", ii)
	// len_ii := len(ii)
	// if len_ii%2 == 1 {
	// ii = append(ii[:len_ii/2-1], ii[len_ii/2+1:])
	// dfdfdfdfdfdfffffff// }
	// fmt.Println(ii)
	fmt.Println("====================")
	rmt.Println("====================")
	輸入 := "一二三四"
	反轉後 := RevStrPkg.Aa(輸入)
	fmt.Println(反轉後)
}
