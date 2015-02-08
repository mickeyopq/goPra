package main

import (
	"RevStrPkg"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inin string
	fmt.Println("請輸入字串:")
	fmt.Scanf("%s", &inin)
	ii := strings.Split(inin, "")
	len_ii := len(ii)
	midNum := len_ii / 2
	if len_ii%2 == 1 {
		//刪了midNum那個元素，富雜的寫法
		ii = ii[:midNum+copy(ii[midNum:], ii[midNum+1:])]
		// ii = append(ii[:len_ii/2-1], ii[len_ii/2+1:])
	}
	frontPart_ii := ii[:midNum]
	lastPart_ii := ii[midNum:] //沒事自動加一幹牛
	last轉 := 片字串轉字串(lastPart_ii)
	front轉 := 片字串轉字串(frontPart_ii)

	last轉2 := RevStrPkg.Aa(last轉)
	if last轉2 == front轉 {
		fmt.Println("true啊斯====================")
	} else {
		fmt.Println("FALSE啊斯====================")
	}
}
func 片字串轉字串(revS []string) string {
	b := make([]rune, len(revS))
	for i := 0; i < len(revS); i++ {
		// b[i] = rune(lastPart_ii[i])
		b[i], _, _, _ = strconv.UnquoteChar(revS[i], '"')
	}
	kk := string(b)
	return kk
}
