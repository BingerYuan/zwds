package ziwei

import (
	"fmt"
	"testing"
)

func TestGetMstar(t *testing.T) {
	mgz := "庚寅"
	mgz = "己丑"
	mgz = "己亥"
	s := GetMonthStarMap(mgz)
	fmt.Printf("\n%s月　左辅星地支:%s\n", mgz, s)
}
