package engine

import (
	"testing"
	"fmt"
)

func TestGenerate(t *testing.T) {

}

func TestDispatch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}, {116.493381, 39.915599}, {116.33068, 39.92224},{116.304233, 39.986398}}
		//obj := New()
		Dispatch(originScope)
	})
}








func TestCheckIntersection(t *testing.T) {
	t.Run("", func(t *testing.T) {
		//originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}}
		originScope := [][]float64{ {116.449112, 39.992589}, {116.493381, 39.915599}}

		rectangle := [][]float64{{116.466064453125, 39.9627685546875}, {116.47705078125, 39.9627685546875}, {116.47705078125, 39.957275390625}, {116.466064453125, 39.957275390625}}

		//rectangle := [][]float64{ {116.38916015625,39.990234375},{116.400146484375,39.990234375}}
		result := checkPointInRectangle(originScope, rectangle)
		result2 := generatesJs(result)
		fmt.Println(result2)


	})
}



//  116.400146484375, 39.9957275390625}, {116.4111328125, 39.9957275390625}, {116.4111328125, 39.990234375}, {116.400146484375, 39.990234375},

// maxlng,maxlat(116.4111328125,39.9957275390625)