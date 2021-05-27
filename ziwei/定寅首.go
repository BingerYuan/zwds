package ziwei

import (
	"strings"
)

//传入年干支 返回月干支　0亥　１子
func GetYinShouArr(ygz string) []string {
	garr := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zarr := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
	gs := getGanS(ygz)
	xmap := map[string]string{
		"甲": "丙", "己": "丙",
		"乙": "戊", "庚": "戊",
		"丙": "庚", "辛": "庚",
		"丁": "壬", "壬": "壬",
		"戊": "甲", "癸": "甲",
	}
	var xarr []string
	for k, v := range xmap {
		if strings.EqualFold(gs, k) {
			arr := sortArr(v, garr)
			arr = append(arr, arr[:2]...)
			for i := 0; i < len(arr); i++ {
				for j := 0; i < len(zarr); j++ {
					if i == j {
						xarr = append(xarr, arr[j]+zarr[j])
						break
					}
				}
			}
			break
		}
	}
	end := xarr[:9]
	head := xarr[9:]
	xarr = append(head, end...)
	return xarr
}
