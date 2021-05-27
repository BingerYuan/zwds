package ziwei

import (
	"fmt"
	"testing"
)

func TestGetYinShouArr(t *testing.T) {
	ygz := "甲子"
	ygz = "戊子"
	ygz = "丁卯" //{"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥", "甲子", "乙丑"}
	arr := GetYinShouArr(ygz)
	fmt.Println(arr)
}
