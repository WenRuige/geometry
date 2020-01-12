package engine

import (
	"strings"
	"fmt"
)

// 生成js字符串
func GenerateJSString(geographicEngine GeographicEngine) string {
	str := "["
	for _, v := range geographicEngine.GridList {
		str += generatesJs(v.Scope)
		str += ","
	}
	str += strings.TrimSuffix(str, ",")
	str += "]"
	return str

}
func generatesJs(result [][]float64) string {
	str := "["
	for _, v := range result {
		if v == nil {
			continue
		}
		str += fmt.Sprintf("[%v,%v],", v[0], v[1])
	}
	str = strings.TrimSuffix(str, ",")
	str += "]"
	return str

}