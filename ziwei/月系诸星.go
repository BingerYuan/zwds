// 左辅　右弼　阴煞　天刑　天姚　天月　天巫
package ziwei

import (
	"strings"
)

//月系主星map k星名  v地支宫位
func GetMonthStarMap(mgz string) map[string]string {
	return findStarZhi(mgz)
}
func findStarZhi(mgz string) map[string]string {
	mzArr := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}      //月支顺序
	zfArr := []string{"辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯"}      //左辅星寅月起辰 顺行
	ybArr := []string{"戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥"}      //右弼星
	ysArr := []string{"寅", "子", "戌", "申", "午", "辰", "寅", "子", "戌", "申", "午", "辰"}      //阴煞
	txArr := []string{"酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申"}      //天刑
	tyArr := []string{"丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子"}      //天姚
	tianYueArr := []string{"戌", "巳", "辰", "寅", "未", "卯", "亥", "未", "寅", "午", "戌", "寅"} //天月
	twArr := []string{"巳", "申", "亥", "寅", "巳", "申", "亥", "寅", "巳", "申", "亥", "寅"}      //天巫

	var zfs string    //左辅
	var ybs string    //右弼
	var yss string    //阴煞
	var txs string    //天刑
	var tys string    //天姚
	var tianys string //天月
	var tws string    //天巫
	for i := 0; i < len(mzArr); i++ {
		if strings.ContainsAny(mgz, mzArr[i]) {
			zfs = zfArr[i]
			ybs = ybArr[i]
			yss = ysArr[i]
			txs = txArr[i]
			tys = tyArr[i]
			tianys = tianYueArr[i]
			tws = twArr[i]
			break
		}
	}
	arr := []string{"左辅", "右弼", "阴煞", "天刑", "天姚", "天月", "天巫"}
	zhiarr := []string{zfs, ybs, yss, txs, tys, tianys, tws}
	var xmap = make(map[string]string)
	for k := 0; k < len(arr); k++ {
		for v := 0; v < len(zhiarr); v++ {
			if k == v {
				xmap[arr[v]] = zhiarr[v]
				break
			}
		}
	}
	return xmap
}

//左辅右弼地支宫位
func findFuBiZhi(mgz string) (string, string) {
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
	return zfs, ybs
}
