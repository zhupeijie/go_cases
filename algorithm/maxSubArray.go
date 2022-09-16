package main

import "fmt"

func main()  {
	m  := []int{-2,1,-3,4,-1,2,1,-5,4}
	max := maxSubArray(m)
	fmt.Println(max)
	m1 := maxSubArrayDynamic(m)
	fmt.Println(m1)
}

//动态规划
//遍历计算以i结尾的最大子序和. 以i=0结尾,最大子序和是nums[0], 以i=1结尾的,最大自诩和是max[num[1],num[0]+num[1]].....以此类推,f[i] = max(f[i-1]+nums[i],nums[i])
//因为f(i),是以第i个字符结尾的
func  maxSubArrayDynamic(nums []int)int {
	max1 := nums[0]
	for i := 1; i < len(nums); i++ {
		 if nums[i]+nums[i-1] > nums[i]{
		 	nums[i] = nums[i]+nums[i-1]
		 }
		 if max1 < nums[i]{
		 	max1 = nums[i]
		 }
	}
	return max1
}




//分治

func maxSubArray(nums []int) int {
	ms := getmax(nums,0,len(nums)-1)
	return ms.mMax
}

func getmax(m []int, l,r int)status {
    if l == r{
    	return status{m[l],m[l],m[l],m[l]}
	}
	n := (l+r)/2
	lM := getmax(m,l,n)
	rM := getmax(m,n+1,r)
	return pushUp(lM,rM)
}

func pushUp(l,r status)status{
	lmax := max(l.lMax,l.iMax+r.lMax)
	rmax := max(r.rMax,r.iMax+l.rMax)
	imax := l.iMax+r.iMax
	mmax := max(max(l.mMax,r.mMax),l.rMax+r.lMax)
	return status{lmax,rmax,imax,mmax}
}
func max(x,y int)int{
	if x > y{
		return x
	}
	return y
}

type status struct {
	//lMax 已最左开头的最大子串和
	//rMax 已最右开头的最大子串和
	//iMax 整串和
	// mMax 最大子串和
	lMax,rMax,iMax,mMax int
}