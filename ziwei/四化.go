package ziwei

import (
	"fmt"
	"strings"
)

//传入四化星数组　四化地支宫位数组  14星map 返回k星名(四化) v地支宫位
func GetSiHuaMap(ygz, mgz, hgz string, star14map map[string]string) map[string]string {
	shmap := siHuaZhi(ygz)
	fbmap := fuBi(mgz)
	cqmap := changQu(hgz)
	maps := mergeMap(star14map, fbmap, cqmap)

	var xmap = make(map[string]string) //k星名(四化)　v地支宫位
	for sk, sv := range shmap {
		for k, v := range maps {
			if strings.EqualFold(sk, k) {
				s := fmt.Sprintf("%s(%s)", k, sv)
				xmap[s] = v
				break
			}
		}
	}
	return xmap
}

//合并多个map
func mergeMap(maps ...map[string]string) map[string]string {
	xmap := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			xmap[k] = v
		}
	}
	return xmap
}

//月系:左辅　右弼
func fuBi(mgz string) map[string]string {
	mzArr := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"} //月支顺序
	zfArr := []string{"辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯"} //左辅星寅月起辰 顺行
	ybArr := []string{"戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥"} //右弼星
	var zfs, ybs string
	for i := 0; i < len(mzArr); i++ {
		if strings.ContainsAny(mgz, mzArr[i]) {
			zfs = zfArr[i]
			ybs = ybArr[i]
			break
		}
	}
	var xmap = make(map[string]string)
	xmap["左辅"] = zfs
	xmap["右弼"] = ybs
	return xmap
}

//时系:文昌　文曲
func changQu(hgz string) map[string]string {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	wencArr := []string{"戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥"}
	wenqArr := []string{"辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯"}
	var wcs, wqs string
	for i := 0; i < len(zhiArr); i++ {
		if strings.ContainsAny(hgz, zhiArr[i]) {
			wcs = wencArr[i]
			wqs = wenqArr[i]
			break
		}
	}
	var xmap = make(map[string]string)
	xmap["文昌"] = wcs
	xmap["文曲"] = wqs
	//fmt.Printf("时干支:%s %v\n", hgz, xmap)
	return xmap
}

//十四正星＋文昌 文曲 右弼 左辅
//传年干支　返回四化星数组　四化星地支宫位数组
func siHuaZhi(ygz string) map[string]string {
	garr := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

	//0甲年....9癸年
	huaLu := []string{"廉贞星", "天机星", "天同星", "太阴星", "贪狼星", "武曲星", "太阳星", "巨门星", "天梁星", "破军星"}
	huaQuan := []string{"破军星", "天梁星", "天机星", "天同星", "太阴星", "贪狼星", "武曲星", "太阳星", "紫微星", "巨门星"}
	huaKe := []string{"武曲星", "紫微星", "文昌", "天机星", "右弼", "天梁星", "太阴星", "文曲", "左辅", "太阴星"}
	huaJi := []string{"太阳星", "太阴星", "廉贞星", "巨门星", "天机星", "文曲", "天同星", "文昌", "武曲星", "贪狼星"}

	var hls, hqs, hks, hjs string
	for i := 0; i < len(garr); i++ {
		if strings.ContainsAny(ygz, garr[i]) {
			hls = huaLu[i]
			hqs = huaQuan[i]
			hks = huaKe[i]
			hjs = huaJi[i]
			break
		}
	}
	//fmt.Printf("%s年 化禄星:%s 化权星:%s 化科星:%s 化忌星:%s \n", ygz, hls, hqs, hks, hjs)
	arrs := []string{"化禄", "化权", "化科", "化忌"}
	arrz := []string{hls, hqs, hks, hjs}
	var xmap = make(map[string]string) //k星名 v四化名称
	for k := 0; k < len(arrs); k++ {
		for v := 0; v < len(arrz); v++ {
			if k == v {
				xmap[arrz[v]] = arrs[v]
				break
			}
		}
	}
	return xmap
}
