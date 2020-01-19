package engine

import (
	"testing"
	"fmt"
)

func TestGenerate(t *testing.T) {

}

func TestDispatch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}, {116.493381, 39.915599}, {116.33068, 39.92224}, {116.304233, 39.986398}}
		Dispatch(originScope)
	})
}

func TestCheckIntersection(t *testing.T) {
	t.Run("", func(t *testing.T) {
		//originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}}
		originScope := [][]float64{{116.449112, 39.992589}, {116.493381, 39.915599}}

		rectangle := [][]float64{{116.488037109375, 39.9298095703125}, {116.47705078125, 39.9298095703125}}

		result := checkPointInRectangle(originScope, rectangle)
		result2 := generatesJs(result)
		fmt.Println(result2)

	})
}

// 检查多边形的顶点是否在矩形内
func TestCheckIntersection2(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// {116.449112, 39.992589}, {116.493381, 39.915599} 这个点是交点
		originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}, {116.493381, 39.915599}, {116.33068, 39.92224}, {116.304233, 39.986398}}

		rectangle := [][]float64{{116.444091796875, 39.995727539062}, {116.455078125, 39.995727539062}, {116.455078125, 39.990234375}, {116.444091796875, 39.990234375}, {116.444091796875, 39.995727539062}}

		result := checkPointInRectangle(originScope, rectangle)
		result2 := generatesJs(result)
		fmt.Println(result2)
	})
}

//  116.400146484375, 39.9957275390625}, {116.4111328125, 39.9957275390625}, {116.4111328125, 39.990234375}, {116.400146484375, 39.990234375},

// maxlng,maxlat(116.4111328125,39.9957275390625)
