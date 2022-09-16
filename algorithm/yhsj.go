package main

import "fmt"

func main(){
	//res := generate(5)
	res := getRow(2)
	fmt.Println("here is a",res)
}
func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		_item := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i{
				_item = append(_item,1)
				continue
			}
			_item = append(_item,res[i-1][j-1]+res[i-1][j])
		}
		res = append(res,_item)
	}
	return res
}


func getRow(rowIndex int) []int {
	res := make([]int,rowIndex+1)
	for i:= 0;i <= rowIndex;i++ {
		res[0] = 1
		res[i] = 1
		for j:= i-1;j>0;j-- {
			res[j] = res[j]+ res[j-1]
		}
	}
	return res
}

