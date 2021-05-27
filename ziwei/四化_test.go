package ziwei

import "testing"

func TestGetSiHua(t *testing.T) {
	ygz := "甲子"
	mgz := "乙丑"
	hgz := "丙寅"
	star14map := map[string]string{"廉贞星": "丑", "文昌": "申"}
	GetSiHuaMap(ygz, mgz, hgz, star14map)
}
