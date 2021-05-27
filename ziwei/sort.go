package ziwei

import "strings"

//顺序排地支
//传入年分对应的地支(寅午戍年生人在"辰"...) 原地支数组 返回排序后的地支数组
func sortArr(zhi string, zhiArr []string) []string {
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zhi, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	return zhiArr
}

//逆序排地支
func rSortArr(zhi string, zhiArr []string) []string {
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zhi, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	//
	head := zhiArr[:1]
	end := zhiArr[1:]
	var rArr []string
	for i := len(end) - 1; i >= 0; i-- {
		rArr = append(rArr, end[i])
	}
	rArr = append(head, rArr...)
	return rArr
}

//取天干 传干支 返回干
func getGanS(gz string) string {
	garr := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	var g string
	for i := 0; i < len(garr); i++ {
		if strings.ContainsAny(gz, garr[i]) {
			g = garr[i]
			break
		}
	}
	return g
}
