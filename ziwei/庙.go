package ziwei

import "strings"

//庙3　旺2　地1　利0　平-1　不得地-2　陷-3
func GetMiaoMap(ygz, hgz string, star14map map[string]string) map[string]string {
	n1, nz1, n2, nz2 := huoLingXing(ygz, hgz)
	n3, nz3, n4, nz4 := qingYangTuoLuo(ygz)
	wcs, wqs := findChangQuZhi(hgz)
	wcn, wqn := "文昌", "文曲"

	starNameArr := []string{n1, n2, n3, n4, wcn, wqn}
	zhiArr := []string{nz1, nz2, nz3, nz4, wcs, wqs}
	xmap := make(map[string]string) //k star v地支宫位
	for i := 0; i < len(starNameArr); i++ {
		for j := 0; j < len(zhiArr); j++ {
			if i == j {
				xmap[starNameArr[j]] = zhiArr[j]
				break
			}
		}
	}
	maps := mergeMap(xmap, star14map)

	//
	starArr := []string{
		"紫微星", "天机星", "太阳星", "武曲星", "天同星", "廉贞星",
		"天府星", "太阴星", "贪狼星", "巨门星", "天相星", "天梁星", "七杀星", "破军星",
		"火星", "铃星", "擎羊星", "陀罗星", "文昌", "文曲",
	}
	haiN := []int{+2, -1, -3, -1, +3, -3, +1, +3, -3, +2, +1, -3, -1, -1, 0, 0, 0, -3, 0, +2}   //宫支亥
	ziN := []int{-1, +3, -3, +2, +2, -1, +3, +3, +2, +2, +3, +3, +2, +3, -3, -3, +2, 0, +1, +1} //宫支子
	chouN := []int{+3, -3, -2, +3, -2, 0, +3, +3, +3, -2, +3, +2, +3, +2, +1, +1, +3, +3, +3, +3}
	yinN := []int{+3, +1, +2, +1, 0, +3, +3, -2, -1, +3, +3, +3, +3, +1, +3, +3, 0, -3, -3, -1}
	maoN := []int{+2, +2, +3, 0, +3, -1, +1, -3, +1, +3, -3, +2, +2, -3, 0, 0, -3, 0, +2, +2}
	chenN := []int{+1, 0, +2, +3, -1, 0, +2, -3, +3, -1, -1, +3, +1, +2, -3, -3, +3, +3, +1, +1}
	siN := []int{+2, -1, +2, -1, +3, -3, +1, -3, -3, -1, +1, -3, -1, -1, +1, +1, 0, -3, +3, +3}
	wuN := []int{+3, +3, +3, +2, -3, -1, +2, -3, +2, +2, +1, +3, +2, +3, +3, +3, -3, 0, -3, -3}
	weiN := []int{+3, -3, +1, +3, -2, 0, +3, -1, +3, -1, +1, +2, +3, +2, 0, 0, +3, +3, 0, +2}
	shenN := []int{+2, +1, 0, +1, +2, +3, +1, 0, -1, +3, +3, -3, +3, +1, -3, -3, 0, -3, +1, +1}
	youN := []int{+2, +2, -1, 0, -1, -1, +2, +2, 0, +3, -3, +1, +2, -3, +1, +1, +2, 0, +3, +3}
	xuN := []int{+1, 0, -3, +3, -1, 0, +2, +2, +3, -1, +1, +3, +3, +2, +3, +3, +3, +3, -3, -3}

	var miaoMap = make(map[string]string)
	for star, szhi := range maps {
		for i := 0; i < len(starArr); i++ {
			if strings.EqualFold(star, starArr[i]) {
				switch szhi {
				case "亥":
					miaoMap = miaoWangXian(starArr, haiN)
				case "子":
					miaoMap = miaoWangXian(starArr, ziN)
				case "丑":
					miaoMap = miaoWangXian(starArr, chouN)
				case "寅":
					miaoMap = miaoWangXian(starArr, yinN)
				case "卯":
					miaoMap = miaoWangXian(starArr, maoN)
				case "辰":
					miaoMap = miaoWangXian(starArr, chenN)
				case "巳":
					miaoMap = miaoWangXian(starArr, siN)
				case "午":
					miaoMap = miaoWangXian(starArr, wuN)
				case "未":
					miaoMap = miaoWangXian(starArr, weiN)
				case "申":
					miaoMap = miaoWangXian(starArr, shenN)
				case "酉":
					miaoMap = miaoWangXian(starArr, youN)
				case "戌":
					miaoMap = miaoWangXian(starArr, xuN)
				}
				break
			}
		}
	}
	return miaoMap
}

//传入星数组 地支数组
func miaoWangXian(starArr []string, zhiN []int) map[string]string {
	xmap := make(map[string]string)
	for i := 0; i < len(starArr); i++ {
		for j := 0; j < len(zhiN); j++ {
			if i == j {
				s := intToS(zhiN[j])
				xmap[starArr[j]] = s //zhiN[j]
				break
			}
		}
	}
	return xmap
}

//庙+3　旺+2　地+1　利 0　平-1　不得地-2　陷-3）
func intToS(zhiN int) string {
	var s string
	switch zhiN {
	case 3:
		s = "庙"
	case 2:
		s = "旺"
	case 1:
		s = "地"
	case 0:
		s = "利"
	case -1:
		s = "平"
	case -2:
		s = "不得地"
	case -3:
		s = "陷"
	}
	return s
}
