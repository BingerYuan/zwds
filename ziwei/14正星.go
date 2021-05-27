package ziwei

import (
	"strings"
)

/* "天机星", "太阳星", "武曲星", "天同星", "廉贞星",
"天府星", "太阴星", "贪狼星", "巨门星", "天相星", "天梁星", "七杀星", "破军星",
"紫微星"*/
//十四正星 k星名　v地支宫位
func GetStar14Map(zwzhi string) map[string]string {
	stararr := []string{"天机星", "太阳星", "武曲星", "天同星", "廉贞星",
		"天府星", "太阴星", "贪狼星", "巨门星", "天相星", "天梁星", "七杀星", "破军星",
		"紫微星"}
	zhiarr := GetStar13Arr(zwzhi)
	zhiarr = append(zhiarr, zwzhi)  //紫微宫地支放到末尾
	smap := make(map[string]string) //k星名 v宫位
	for i := 0; i < len(stararr); i++ {
		for j := 0; j < len(zhiarr); j++ {
			if i == j {
				smap[stararr[j]] = zhiarr[j]
				break
			}
		}
	}
	return smap
}

//"天机星", "太阳星", "武曲星", "天同星", "廉贞星","天府星", "太阴星", "贪狼星", "巨门星", "天相星", "天梁星", "七杀星", "破军星"
//返回 13星数组(这里没有紫微星的地支宫位) 0天机 5天府
func GetStar13Arr(zwzhi string) []string {
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	tjarr := star5("亥", zhiArr)
	taiYangArr := star5("酉", zhiArr)
	wuQuArr := star5("申", zhiArr)
	tianTongArr := star5("未", zhiArr)
	lianZhenArr := star5("辰", zhiArr)
	//
	tianFuArr := star8("辰", zhiArr)
	taiYinArr := star8("巳", zhiArr)
	tanLangArr := star8("午", zhiArr)
	juMenArr := star8("未", zhiArr)
	tianXiangArr := star8("申", zhiArr)
	tianLiangArr := star8("酉", zhiArr)
	qiShaArr := star8("戌", zhiArr)
	poJunArr := star8("寅", zhiArr)
	return star13(zwzhi, zhiArr,
		tjarr, taiYangArr, wuQuArr, tianTongArr, lianZhenArr,
		tianFuArr, taiYinArr, tanLangArr, juMenArr, tianXiangArr, tianLiangArr, qiShaArr, poJunArr)
}
func star13(zwzhi string, zhiArr []string, arr ...[]string) []string {
	var n int
	for x := 0; x < len(zhiArr); x++ {
		if strings.EqualFold(zwzhi, zhiArr[x]) {
			n = x
			break
		}
	}
	////
	var arr13 []string
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if n == j {
				arr13 = append(arr13, arr[i][j])
				break
			}
		}
	}
	return arr13
}

//紫微在子 天机在亥 传入紫微星的地支宫位  地支数组 返回天机星在十二辰数组 0紫微在子时的天机宫位 1紫微在丑时的天机宫位
//"天机星", "太阳星", "武曲星", "天同星", "廉贞星",
func star5(zwzhi string, zhiArr []string) []string {
	return sortArr(zwzhi, zhiArr)
}

//"天府星", "太阴星", "贪狼星", "巨门星", "天相星", "天梁星", "七杀星", "破军星"
func star8(zwzhi string, zhiArr []string) []string {
	return rSortArr(zwzhi, zhiArr)
}
