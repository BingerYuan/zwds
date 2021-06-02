// 火星 铃星 三台星 八座星 天贵星 恩光星 天才星 天寿星 天伤 天使星
package ziwei

import (
	"strings"
)

/*
火星: 以出生年支和出生时支为条件，从表查出火星所在宫位，并填到空白命盘图内。例：寅年生人在寅午戍栏内查，若子时生则火星在丑，……
铃星: 以出生年支和出生时支为条件，从表查出铃星所在宫位，并填到空白命盘图内。例:子年生人，在申子辰栏内查，若子时生则铃星在戍宫。

三台星: 由左辅星所在宫位上起初一日，沿十二宫顺时针方向数，数至出生日止，即在此宫安三台星。
八座星: 由右弼星所在宫位上起初一日，沿十二宫逆时针方向数，数至出生日止，即在此宫安八座星。

天贵星: 由文曲星所在宫位上起初一日，顺时针方向数，数至出生日宫再退回一宫，即在此宫安天贵星。
恩光星: 由文昌星所在宫位上起初一日,顺时针方向数，数至出生日宫再退回一宫，即在此宫安恩光星。

天才星: 由命宫起子年，沿十二宫顺时针方向数，数至出生年支所在宫止，即在此宫安天才星。
天寿星: 由身宫起子年，沿十二宫顺时针方向数，数至出生年支所在宫止，即在此宫安天寿星。

天伤、天使星: 天伤星固定在奴仆宫，天使星固定在疾厄宫。
*/

//k星名　v地支宫位
func GetOtherMap(ygz, mgz, hgz string, lday int) map[string]string {
	n1, nz1, n2, nz2 := huoLingXing(ygz, hgz)
	n3, nz3, n4, nz4 := sanTaiBaZuo(mgz, lday)
	n5, nz5, n6, nz6 := tianGuienGuang(hgz, lday)
	n7, nz7, n8, nz8 := tianCaiTianShou(ygz, mgz, hgz)
	starArr := []string{n1, n2, n3, n4, n5, n6, n7, n8}
	zhiArr := []string{nz1, nz2, nz3, nz4, nz5, nz6, nz7, nz8}
	var xmap = make(map[string]string)
	for i := 0; i < len(starArr); i++ {
		for j := 0; j < len(zhiArr); j++ {
			xmap[starArr[j]] = zhiArr[j]
		}
	}
	return xmap
}

//天才 天才地支宫位  天寿 天寿地支宫位
func tianCaiTianShou(ygz, mgz, hgz string) (string, string, string, string) {
	mgZhi, _, sgZhi := mgsg(mgz, hgz)
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	mgarr := sortArr(mgZhi, zhiArr) //排序命宫数组
	sgarr := sortArr(sgZhi, zhiArr) //排序身宫数组
	yz := GetZhiS(ygz)          //年支
	var tcs, tss string
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(yz, zhiArr[i]) {
			tcs = mgarr[i]
			tss = sgarr[i]
		}
	}
	return "天才", tcs, "天寿", tss
}

//天贵　恩光
func tianGuienGuang(hgz string, lday int) (string, string, string, string) {
	wcs, wqs := findChangQuZhi(hgz)
	enGuangs, egzhi := enGuan(wcs, lday)
	tianGuis, tgzhi := tianGui(wqs, lday)
	return enGuangs, egzhi, tianGuis, tgzhi
}

//恩光 恩光地支宫位
func enGuan(wcs string, lday int) (string, string) {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(wcs, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	zhiArr = append(zhiArr, zhiArr...)
	zhiArr = append(zhiArr, zhiArr...)
	zhiArr = append(zhiArr, zhiArr...)

	var egs string
	for j := 0; j < len(zhiArr); j++ {
		if j == lday-2 {
			egs = zhiArr[j]
			break
		}
	}
	return "恩光", egs
}

//天贵 天贵地支宫位
func tianGui(wqs string, lday int) (string, string) {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(wqs, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	zhiArr = append(zhiArr, zhiArr...)
	zhiArr = append(zhiArr, zhiArr...)
	zhiArr = append(zhiArr, zhiArr...)

	var tgs string
	for j := 0; j < len(zhiArr); j++ {
		if j == lday-2 {
			tgs = zhiArr[j]
			break
		}
	}
	return "天贵", tgs
}

//三台名称　地支宫位　八座名称　地支宫位
func sanTaiBaZuo(mgz string, lday int) (string, string, string, string) {
	zfs, ybs := findFuBiZhi(mgz)
	san, sanZhi := sanTai(zfs, lday)
	ba, baZhi := baZuo(ybs, lday)
	return san, sanZhi, ba, baZhi
}

//八座
func baZuo(ybs string, lday int) (string, string) {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var rArr []string
	for i := len(zhiArr) - 1; i >= 0; i-- {
		rArr = append(rArr, zhiArr[i])
	}

	for j := 0; j < len(rArr); j++ {
		if strings.EqualFold(ybs, rArr[j]) {
			head := rArr[:j]
			end := rArr[j:]
			rArr = append(end, head...)
			break
		}
	}

	rArr = append(rArr, rArr...)
	rArr = append(rArr, rArr...)
	var bzs string
	for k := 0; k < len(rArr); k++ {
		if k == lday -1 {
			bzs = rArr[k]
			break
		}
	}
	return "八座", bzs
}

//三台星　传左辅地支宫位 阴历日数字　返回星名　地支宫位
func sanTai(zfs string, lday int) (string, string) {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//找到左辅位置
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zfs, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	zhiArr = append(zhiArr, zhiArr...)
	zhiArr = append(zhiArr, zhiArr...)
	var sts string
	for j := 0; j < len(zhiArr); j++ {
		if j == lday -1 {
			sts = zhiArr[j]

			break
		}
	}
	return "三台", sts
}

// 火星　火星地支宫位 铃星　铃星地支宫位
func huoLingXing(ygz, hgz string) (string, string, string, string) {
	arr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"} //地支
	//火星
	ywx := []string{"丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子"} //寅午戌
	szc := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"} //申子辰
	syc := []string{"卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅"} //巳酉丑
	hmw := []string{"酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申"} //亥卯未
	//铃星
	ywxl := []string{"卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅"} //寅午戌
	szcl := []string{"戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉"} //申子辰
	sycl := []string{"戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉"} //巳酉丑
	hmwl := []string{"戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉"} //亥卯未

	yz := GetZhiS(ygz)
	hz := GetZhiS(hgz)
	var s, sl string
	switch yz {
	case "寅", "午", "戌":
		s = findS(hz, arr, ywx)
		sl = findS(hz, arr, ywxl)
	case "申", "子", "辰":
		s = findS(hz, arr, szc)
		sl = findS(hz, arr, szcl)
	case "巳", "酉", "丑":
		s = findS(hz, arr, syc)
		sl = findS(hz, arr, sycl)
	case "亥", "卯", "未":
		s = findS(hz, arr, hmw)
		sl = findS(hz, arr, hmwl)
	}
	return "火星", s, "铃星", sl
}
func findS(zhi string, arr, xyz []string) string {
	var s string
	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(arr[i], zhi) {
			s = xyz[i]
			break
		}
	}
	return s
}
