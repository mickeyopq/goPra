package RevStrPkg

func Aa(input string) (output string) {
	n := 0
	//切割成slice，可改成split()試試
	ss1 := make([]rune, len(input))
	for _, r := range input {
		ss1[n] = r
		n++
	}
	//修整slice，長改短
	ss1 = ss1[0:n]
	//反轉字串
	for i := 0; i < n/2; i++ {
		// 前後=後前，close to center
		ss1[i], ss1[n-i-1] = ss1[n-i-1], ss1[i]
	}
	// slice轉成字串
	output = string(ss1)
	return
}
