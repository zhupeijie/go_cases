package main

import "fmt"

func main() {

	for i:=0;i<1000;i++{
		x := i*i
		y := mySqrt(x)
		if i != y{
			fmt.Println("=======",i,x,y)
			continue
		}
		//fmt.Println(i,x,y)
	}
	//base := mySqrt(x)
	fmt.Println("end")
	fmt.Println(mySqrt(996012))
}

func mySqrt(x int) int {
	a := 0
	tmp := 0
	base, _ := my1(x, a, tmp)
	return base
}

func my1(x int, a, tmp int) (int, int) {
	//fmt.Println(x,a,tmp)
	for x >= 100 {
		a,tmp = my1(x/100,a,tmp)
		x = x%100
		return calcu(x,a,tmp)
	}
	return calcu(x,a,tmp)
}


func calcu(x ,a,tmp int)(int ,int) {
	for b := 0; b < 10; b++ {
		if (a*10*2+b)*b <= (tmp*100+x) && (a*10*2+b+1)*(b+1) > (tmp*100+x) {
			tmp = tmp*100 + x - (a*10*2+b)*b
			a = a*10 + b
			break
		}
	}
	return a, tmp
}