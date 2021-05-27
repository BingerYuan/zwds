package ziwei

import (
	"fmt"
	"testing"
)

func TestGetWuXing(t *testing.T) {
	ygz := "甲子"
	mgz := "甲寅"
	hgz := "乙丑"
	mg, _, _ := GetMsgMap(mgz, hgz)
	mgzArr := GetYinShouArr(ygz)
	fmt.Printf("命宫:%s 月干支:%s\n", mg, mgzArr)
	n, s := GetWuXing(mg, mgzArr)
	fmt.Println(n, s)
}
