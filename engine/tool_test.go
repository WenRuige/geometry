package engine

import "testing"

func TestGenerateGeohash(t *testing.T) {
	t.Run("", func(t *testing.T) {
		GenerateGeohash([]float64{116.3943515,39.9542065},5)
	})
}
