package main

import (
	"fmt"
	"math"
)

// 多边形面积

/*

	平面状态下求多边形面积
	地理坐标系求多边形面积:= s * 1e-5
*/

func polygonarea(arr [][]float64) float64 {
	sum := float64(0)
	arr = append(arr, arr[0])
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			sum = sum + arr[i][0]*arr[i+1][1] - arr[i+1][0]*arr[i][1]
		}
	}
	return math.Abs(sum / 2)
}

func main() {
	//p := [][]float64{{-3, -2}, {-1, 4}, {6, 1}, {3, 10}, {-4, 9}}
	p := [][]float64{{0, 0}, {0, 2}, {2, 2}, {2, 0}}
	res := polygonarea(p)
	fmt.Println(res)
}
