package ziwei

import "testing"

func TestGetOtherMap(t *testing.T) {
	ygz := "寅午戌"
	ygz = "申子辰"
	mgz := "庚午"
	hgz := "己丑"
	lday := 29
	GetOtherMap(ygz, mgz, hgz, lday)
}
