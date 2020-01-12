package engine

import "testing"

func TestGenerate(t *testing.T) {

}



func TestDispatch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		originScope := [][]float64{{116.304233,39.986398}, {116.449112,39.992589}, {116.493381,39.915599}, {116.33068,39.92224}}
		//obj := New()
		Dispatch(originScope)
	})
}
