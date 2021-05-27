package ziwei

//传入五行局(木火土水金) 性别 返回map k地支宫位 v十二长生名称
func GetChangShengMap(wxs, sex string) map[string]string {
	changArr := []string{"长生", "沐浴", "冠带", "临官", "帝旺", "衰", "病", "死", "墓", "绝", "胎", "养"}
	var chmap = make(map[string]string)
	switch wxs {
	case "金":
		zhi := []string{"巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰"}
		switch sex {
		case "男":
			chmap = csMaps(zhi, changArr)
		case "女":
			chmap = csMaps(rZhi(zhi), changArr)
		}
	case "木":
		zhi := []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}
		switch sex {
		case "男":
			chmap = csMaps(zhi, changArr)
		case "女":
			chmap = csMaps(rZhi(zhi), changArr)
		}
	case "火":
		zhi := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
		switch sex {
		case "男":
			chmap = csMaps(zhi, changArr)
		case "女":
			chmap = csMaps(rZhi(zhi), changArr)
		}
	case "水", "土":
		zhi := []string{"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"}
		switch sex {
		case "男":
			chmap = csMaps(zhi, changArr)
		case "女":
			chmap = csMaps(rZhi(zhi), changArr)
		}
	}
	return chmap
}
func csMaps(zhi, changArr []string) map[string]string {
	csmap := make(map[string]string)
	for i := 0; i < len(zhi); i++ {
		for j := 0; j < len(changArr); j++ {
			if i == j {
				csmap[zhi[j]] = changArr[j]
			}
		}
	}
	return csmap
}

//逆排十二辰
func rZhi(zhi []string) []string {
	head := zhi[:1]
	zhi = zhi[1:]
	var arr []string
	for i := len(zhi) - 1; i >= 0; i-- {
		arr = append(arr, zhi[i])
	}
	arr = append(head, arr...)
	return arr
}
