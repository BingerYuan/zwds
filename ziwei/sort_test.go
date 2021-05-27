package ziwei

import (
	"fmt"
	"testing"
)

func TestSortZhiArr(t *testing.T) {
	zarr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	zhi := zarr[11]
	arr := sortArr(zhi, zarr)
	fmt.Println(arr)
	rarr := rSortArr(zhi, zarr)
	fmt.Println(rarr)
}

func TestGetGans(t *testing.T) {
	ygz := "甲子"
	ygz = "乙丑"
	ygz = "癸亥"
	gs := getGanS(ygz)
	fmt.Println(ygz, gs)
}
