package main

import (
	"github.com/geometry/base"
	"math"
	"fmt"
)

/*
	判断两条线段是否相交
	参考:https://www.cnblogs.com/wuwangchuxin0924/p/6218494.html
	设:以线段p1p2为对角线的矩形为R,设q1q2为对角线的矩形为T,若R,T不想交,则两线段不可能相交
	快速排斥:以两条线段的对角线为矩形,如果不重合的话,那么两条线段一定不可能相交
	1. 线段ab的低点低于cd的高点
*/

func issegmentsintersect(line1 [] base.Point, line2 []base.Point) bool {
	if (math.Min(line1[0].X, line1[1].X) <= math.Max(line2[0].X, line2[1].X) &&
		math.Min(line2[0].Y, line2[1].Y) <= math.Max(line1[0].Y, line1[1].Y) &&
		math.Min(line2[0].X, line2[1].X) <= math.Max(line1[0].X, line1[1].X) &&
		math.Min(line1[0].Y, line1[1].Y) <= math.Max(line2[0].Y, line2[1].Y)) {
		return true
	}
	return false
}

func main() {

	var line1 [] base.Point
	var line2 []base.Point
	point1 := base.Point{
		0, 0,
	}
	point2 := base.Point{
		2, 2,
	}

	line1 = append(line1, point1)
	line1 = append(line1, point2)

	point3 := base.Point{
		0, 1,
	}
	point4 := base.Point{
		1, 0,
	}
	line2 = append(line2, point3)
	line2 = append(line2, point4)

	res := issegmentsintersect(line1, line2)
	fmt.Println(res)

}
