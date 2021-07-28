package main

func main() {

}

func descNumString(n int) string {
	numsStr := "1"
	if n == 1 {
		return numsStr
	}

	for i := 2; i <= n; i++ {
		numsStr = descStr(numsStr)
	}
	return numsStr
}

func descStr(numStr string) string {
	
	return numStr
}
