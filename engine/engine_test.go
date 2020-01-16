package engine

import (
	"testing"
	"fmt"
)

func TestGenerate(t *testing.T) {

}

func TestDispatch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}, {116.493381, 39.915599}, {116.33068, 39.92224}}
		//obj := New()
		Dispatch(originScope)
	})
}

func TestCheckIntersection(t *testing.T) {
	t.Run("", func(t *testing.T) {
		originScope := [][]float64{{116.304233, 39.986398}, {116.449112, 39.992589}}
		rectangle := [][]float64{{116.38916015625, 39.9957275390625}, {116.400146484375, 39.9957275390625}, {116.400146484375, 39.990234375}, {116.38916015625, 39.990234375}, {116.38916015625, 39.9957275390625}}

		//rectangle := [][]float64{ {116.38916015625,39.990234375},{116.400146484375,39.990234375}}
		result := CheckIntersection(originScope, rectangle)
		fmt.Println(result)


	})
}
