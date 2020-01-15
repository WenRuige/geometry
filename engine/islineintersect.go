package engine

import "math"

// 线段是是否相交

func IsLineIntersect(lineA []float64, lineB []float64, lineC []float64, lineD []float64) bool {

	if (math.Min(lineA[0], lineB[0]) <= math.Max(lineC[0], lineD[0]) && //  快速排斥实验；
		math.Min(lineC[0], lineD[0]) <= math.Max(lineA[0], lineB[0]) &&
		math.Min(lineA[1], lineB[1]) <= math.Max(lineC[1], lineD[1]) &&
		math.Min(lineC[1], lineD[1]) <= math.Max(lineA[1], lineB[1]) &&
		straddleExperiment(lineA, lineB, lineC)*straddleExperiment(lineA, lineB, lineD) < 0 && //  跨立实验；
		straddleExperiment(lineC, lineD, lineA)*straddleExperiment(lineC, lineD, lineB) < 0) {
		return true
	}
	return false
}


//double Cross_Prouct(node A, node B, node C) //  计算BA叉乘CA；
//{
//return (B.x-A.x)*(C.y-A.y)-(B.y-A.y)*(C.x-A.x); //向量叉积运算
//}

// 快速排斥实验
func straddleExperiment(lineA []float64, lineB []float64, lineC []float64) float64 {
	return (lineB[0] - lineA[0]) *(lineC[1]-lineA[1]) - (lineB[1]-lineA[1])*(lineC[0]-lineA[0])
}

// 跨立实验
//
//func straddleExperiment(){
//
//}
