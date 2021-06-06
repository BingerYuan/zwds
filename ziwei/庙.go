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
	youN := []int{+3, +2, -1, 0, -1, -1, +2, +2, 0, +3, -3, +1, +2, -3, +1, +1, +2, 0, +3, +3}
	xuN := []int{+1, 0, -3, +3, -1, 0, +2, +2, +3, -1, +1, +3, +3, +2, +3, +3, +3, +3, -3, -3}

	var miaoMap = make(map[string]string)

	var haiStarArr []string
	var haiMap = make(map[string]string)

	var ziStarArr []string
	var ziMap = make(map[string]string)

	var chouStarArr []string
	var chouMap = make(map[string]string)

	var yinStarArr []string
	var yinMap = make(map[string]string)

	var maoStarArr []string
	var maoMap = make(map[string]string)

	var chenStarArr []string
	var chenMap = make(map[string]string)

	var siStarArr []string
	var siMap = make(map[string]string)

	var wuStarArr []string
	var wuMap = make(map[string]string)

	var weiStarArr []string              //星数组
	var weiMap = make(map[string]string) //星庙map k星名 v庙．．．．

	var shenStarArr []string
	var shenMap = make(map[string]string)

	var youStarArr []string
	var youMap = make(map[string]string)

	var xuStarArr []string
	var xuMap = make(map[string]string)

	for star, szhi := range maps {
		switch szhi {
		case "亥":
			haiStarArr = append(haiStarArr, star)
			haiMap = findMiaoMap(haiStarArr, starArr, haiN)
		case "子":
			ziStarArr = append(ziStarArr, star)
			ziMap = findMiaoMap(ziStarArr, starArr, ziN)
		case "丑":
			chouStarArr = append(chouStarArr, star)
			chouMap = findMiaoMap(chouStarArr, starArr, chouN)
		case "寅":
			yinStarArr = append(yinStarArr, star)
			yinMap = findMiaoMap(yinStarArr, starArr, yinN)
		case "卯":
			maoStarArr = append(maoStarArr, star)
			maoMap = findMiaoMap(maoStarArr, starArr, maoN)
		case "辰":
			chenStarArr = append(chenStarArr, star)
			chenMap = findMiaoMap(chenStarArr, starArr, chenN)
		case "巳":
			siStarArr = append(siStarArr, star)
			siMap = findMiaoMap(siStarArr, starArr, siN)
		case "午":
			wuStarArr = append(wuStarArr, star)
			wuMap = findMiaoMap(wuStarArr, starArr, wuN)
		case "未":
			weiStarArr = append(weiStarArr, star)
			weiMap = findMiaoMap(weiStarArr, starArr, weiN)
		case "申":
			shenStarArr = append(shenStarArr, star)
			shenMap = findMiaoMap(shenStarArr, siStarArr, shenN)
		case "酉":
			youStarArr = append(youStarArr, star)
			youMap = findMiaoMap(youStarArr, starArr, youN)
		case "戌":
			xuStarArr = append(xuStarArr, star)
			xuMap = findMiaoMap(xuStarArr, siStarArr, xuN)
		}
	}
	miaoMap = mergeMap(haiMap, ziMap, chouMap, yinMap, maoMap, chenMap, siMap, wuMap, weiMap, shenMap, youMap, xuMap)
	return miaoMap
}

// xZhiStarArr:地支宫位对应的星数组 k:星名 v:庙....
func findMiaoMap(xZhiStarArr []string, starArr []string, weiN []int) map[string]string {
	var xZhiNarr []int
	for x := range xZhiStarArr {
		for i := 0; i < len(starArr); i++ {
			if strings.EqualFold(xZhiStarArr[x], starArr[i]) {
				xZhiNarr = append(xZhiNarr, i)
				break
			}
		}
	}
	var xStarMiaoMap = make(map[string]string)
	for wi := 0; wi < len(xZhiNarr); wi++ {
		index := xZhiNarr[wi]
		for wj := 0; wj < len(weiN); wj++ {
			if wj == index {
				sn := starArr[index]
				indexMiaoN := weiN[index]
				s := intToS(indexMiaoN)
				xStarMiaoMap[sn] = s
			}
		}
	}
	return xStarMiaoMap
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
