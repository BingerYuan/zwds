// 文昌　文曲　天空　地劫　台辅　封诰
package ziwei

import (
	"strings"
)

//时系诸星　k星名 v地支宫位
func GetHourStarMap(hgz string) map[string]string {
	return findHourZhiStar(hgz)
}
func findHourZhiStar(hgz string) map[string]string {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//	文昌　文曲　天空　地劫　台辅　封诰
	wencArr := []string{"戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥"}
	wenqArr := []string{"辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯"}
	tkArr := []string{"亥", "戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子"}
	djArr := []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}
	taifArr := []string{"午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳"}
	fenggArr := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}

	var wcs, wqs, tks, djs, tfs, fgs string
	for i := 0; i < len(zhiArr); i++ {
		if strings.ContainsAny(hgz, zhiArr[i]) {
			wcs = wencArr[i]
			wqs = wenqArr[i]
			tks = tkArr[i]
			djs = djArr[i]
			tfs = taifArr[i]
			fgs = fenggArr[i]
			break
		}
	}
	//fmt.Printf("文昌:%s　文曲:%s　天空:%s　地劫:%s　台辅:%s　封诰:%s\n", wcs, wqs, tks, djs, tfs, fgs)
	var xmap = make(map[string]string)
	arrs := []string{"文昌", "文曲", "天空", "地劫", "台辅", "封诰"}
	arrz := []string{wcs, wqs, tks, djs, tfs, fgs}
	for k := 0; k < len(arrs); k++ {
		for v := 0; v < len(arrz); v++ {
			if k == v {
				xmap[arrs[v]] = arrz[v]
				break
			}
		}
	}
	return xmap
}

//文昌　文曲支宫位
func findChangQuZhi(hgz string) (string, string) {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//	文昌　文曲　天空　地劫　台辅　封诰
	cArr := []string{"戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥"}
	qArr := []string{"辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯"}

	var wcs, wqs string
	for i := 0; i < len(zhiArr); i++ {
		if strings.ContainsAny(hgz, zhiArr[i]) {
			wcs = cArr[i]
			wqs = qArr[i]
			break
		}
	}
	return wcs, wqs
}
