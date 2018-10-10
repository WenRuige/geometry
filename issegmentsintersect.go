package main

import "github.com/geometry/base"

/*
	判断两条线段是否相交
	参考:https://www.cnblogs.com/wuwangchuxin0924/p/6218494.html
	设:以线段p1p2为对角线的矩形为R,设q1q2为对角线的矩形为T,若R,T不想交,则两线段不可能相交
	快速排斥:以两条线段的对角线为矩形,如果不重合的话,那么两条线段一定不可能相交
	1. 线段ab的低点低于cd的高点
*/

func issegmentsintersect(point base.Point) {

}

func main() {
	issegmentsintersect()

}
