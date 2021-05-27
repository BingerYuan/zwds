package ziwei

import (
	"strings"
)

//命主 身主 传入年干支 命宫地支
func GetMsgZhu(ygz, mg string) (string, string) {
	mgmap := map[string]string{
		"子": "贪狼", "丑": "巨门", "寅": "禄存", "卯": "文曲", "辰": "廉贞", "巳": "武曲",
		"午": "破军", "未": "武曲", "申": "廉贞", "酉": "文曲", "戌": "禄存", "亥": "巨门",
	}
	sgmap := map[string]string{
		"子": "铃星", "丑": "天相", "寅": "天梁", "卯": "天同", "辰": "文昌", "巳": "天机",
		"午": "火星", "未": "天相", "申": "天梁", "酉": "天同", "戌": "文昌", "亥": "天机",
	}
	var mgs, sgs string
	for k, v := range mgmap {
		if strings.EqualFold(mg, k) {
			mgs = v
			break
		}
	}
	for k, v := range sgmap {
		if strings.ContainsAny(ygz, k) {
			sgs = v
			break
		}
	}
	return mgs, sgs
}

//排命盘 返回命宫地支　身宫地支　命盘
func GetMsgMap(mgz, hgz string) (string, string, map[string]string) {
	mg, mgarr, sg := mgsg(mgz, hgz)
	xmap := mingPan(mg, mgarr)
	return mg, sg, xmap
}
func mgsg(mgz, hgz string) (string, []string, string) {
	zarr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//命宫 逆排十二支
	var mgArr []string
	for j := len(zarr) - 1; j >= 0; j-- {
		mgArr = append(mgArr, zarr[j])
	}
	//命宫 出生月支上逆排十二支
	for i := 0; i < len(mgArr); i++ {
		if strings.ContainsAny(mgz, mgArr[i]) {
			head := mgArr[i:]
			end := mgArr[:i]
			mgArr = append(head, end...)
			break
		}
	}
	//身宫 出生月支顺排十二支
	var sgArr []string
	for i := 0; i < len(zarr); i++ {
		if strings.ContainsAny(mgz, zarr[i]) {
			head := zarr[i:]
			end := zarr[:i]
			sgArr = append(head, end...)
		}
	}

	var mgs, sgs string //命宫 身宫
	for i := 0; i < len(zarr); i++ {
		if strings.ContainsAny(hgz, zarr[i]) {
			//逆行找命宫
			for j := 0; j < len(mgArr); j++ {
				if i == j {
					mgs = mgArr[j]
					break
				}
			}
			//顺行找身宫
			for k := 0; k < len(sgArr); k++ {
				if i == k {
					sgs = sgArr[k]
					break
				}
			}
			break
		}
	}
	return mgs, mgArr, sgs //, sgArr
}
func mingPan(mg string, mgArr []string) map[string]string {
	for i := 0; i < len(mgArr); i++ {
		if strings.EqualFold(mg, mgArr[i]) {
			head := mgArr[i:]
			end := mgArr[:i]
			mgArr = append(head, end...)
			break
		}
	}

	var mpMap = make(map[string]string) //k命盘地支 v命盘名称
	mpArr := []string{"命宫", "兄弟宫", "夫妻宫", "子女宫", "财帛宫", "疾厄宫", "迁移宫", "奴仆宫", "事业宫", "田宅宫", "福德宫", "父母宫"}
	for i := 0; i < len(mgArr); i++ {
		for j := 0; j < len(mpArr); j++ {
			if i == j {
				mpMap[mgArr[j]] = mpArr[j]
				break
			}
		}
	}
	return mpMap
}
