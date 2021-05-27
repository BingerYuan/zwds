package ziwei

import (
	"fmt"
	"strings"
)

//#######################################################
//起大限
//#######################################################
//传入年干支 月干支 时干支 性别 返回map k地支 v大限年龄
func GetDaXianMap(ygz, mgz, hgz, sex string) map[string]string {
	return daXian(ygz, mgz, hgz, sex)
}

/*
 阳干 甲 丙 戊 庚 壬
 阳干出生的男人 从命宫起运 顺行 每十年安一宫
 阳干出生的女人 从命宫起运 逆行 每十年安一宫

 阴干 乙 丁 己 辛 癸
阴干出生的男人 从命宫起运 逆行 每十年安一宫
阴干出生的女人 从命宫起运 顺行 每十年安一宫
*/
/*
年龄(虚岁)以命宫五行局数为标准。
例如，水二局第一大限为命宫，从2岁起至11岁止；
第二大限阳男阴女为父母宫，阴男阳女为兄弟宫，从12岁起至21岁止；第三大限类推。
*/
//起大限 传入年干支 月干支 时干支 性别 返回map k地支 v大限年龄
func daXian(ygz, mgz, hgz, sex string) map[string]string {
	mg, _, _ := mgsg(mgz, hgz)
	mgzArr := GetYinShouArr(ygz)
	wxn, _ := wuXing(mg, mgzArr)
	arr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	//找命宫地支
	var mGArr []string //命宫
	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(mg, arr[i]) {
			head := arr[:i]
			end := arr[i:]
			mGArr = append(end, head...) //命宫排首位
			break
		}
	}
	/*
		 step 周期数字固定为9
		 wxn  五行局数
		十年周期数字 n=0第一个十年 n=1第二个十年
		大限起始年龄= wxn+n*10
		大限终止年龄=(wxn+step)+n*10
		十年n周期数字 n=1第一个十年 n=2第二个十年
	*/
	step := 9
	var min, max int
	var ageArr []string //大限数组 0第一个十年
	//j==0第一个十年大限 min:十年大限的起始年龄 max:十年大限的终止年龄
	for j := 0; j < len(arr); j++ {
		min = wxn + j*10
		max = (wxn + step) + j*10
		age := fmt.Sprintf("%d-%d岁", min, max)
		ageArr = append(ageArr, age)
	}

	gs := GetGanS(ygz)
	var xmap = make(map[string]string)
	switch gs {
	case "甲", "丙", "戊", "庚", "壬": //阳干
		switch sex {
		case "男":
			xmap = yangNanYiNv(mGArr, ageArr)
		case "女":
			xmap = yangNvYinNan(mGArr, ageArr)
		}
	case "乙", "丁", "己", "辛", "癸": //阴干
		switch sex {
		case "男":
			xmap = yangNvYinNan(mGArr, ageArr)
		case "女":
			xmap = yangNanYiNv(mGArr, ageArr)
		}
	}
	return xmap
}

// 顺行十二辰 阳干出生的男人 阴干出生的女人
//传入 命宫地支排在首位的十二辰数组 大限年龄数组 返回map k地支 v大限年龄
func yangNanYiNv(mGArr, ageArr []string) map[string]string {
	var dxMap = make(map[string]string)
	for i := 0; i < len(mGArr); i++ {
		for j := 0; j < len(ageArr); j++ {
			if i == j {
				dxMap[mGArr[j]] = ageArr[j]
				break
			}
		}
	}
	return dxMap
}

// 逆行十二辰 阳干出生的女人 阴干出生的男人
//传入 命宫地支排在首位的十二辰数组 大限年龄数组 返回map k地支 v大限年龄
func yangNvYinNan(mGArr, ageArr []string) map[string]string {
	head := mGArr[:1]
	mGArr = mGArr[1:]

	//逆行十二辰
	var arr []string
	for x := len(mGArr) - 1; x >= 0; x-- {
		arr = append(arr, mGArr[x])
	}
	arr = append(head, arr...)

	var dxMap = make(map[string]string)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(ageArr); j++ {
			if i == j {
				dxMap[arr[j]] = ageArr[j]
				break
			}
		}
	}
	return dxMap
}

//#########################################################
//起小限
//##########################################################

//传入年干支 性别 返回map k地支宫位 v宫位对应的年龄数组
func GetXiaoXianMap(ygz, sex string) map[string][]int {
	return xiaoXian(ygz, sex)
}

/*
起小限岁数(虚岁)：
小限宫起点宫所在宫位如下：在生年支三合库支的对宫，
即寅午戍年生人在辰，申子辰年生人在戍，
亥卯未年生人在丑，巳酉丑年生人在未。
不论阳阴，男顺女逆，从小限宫起一岁，一年行一宫。
例如：寅年生人，小限宫在辰，即从辰宫宫起1岁，男命顺下巳宫安2岁，午宫安3岁，……；
女命从辰宫起1岁，逆数下至卯宫安2岁，寅宫安3岁，亥宫安4岁，……。
*/
//传入年干支 性别 返回map k地支宫位 v改宫位对应的年龄数组
func xiaoXian(ygz, sex string) map[string][]int {
	zs := GetZhiS(ygz) //年支
	zhiArr := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var xmap = make(map[string][]int)
	switch zs {
	case "寅", "午", "戌": //辰
		rzhiArr := sortArr("辰", zhiArr)
		switch sex {
		case "男": //顺
			xmap = shunXing(rzhiArr)
		case "女": //逆
			xmap = niXing(rzhiArr)
		}
	case "申", "子", "辰": //戌
		rzhiArr := sortArr("戌", zhiArr)
		switch sex {
		case "男":
			xmap = shunXing(rzhiArr)
		case "女":
			xmap = niXing(rzhiArr)
		}
	case "亥", "卯", "未": //丑
		rzhiArr := sortArr("丑", zhiArr)
		switch sex {
		case "男":
			xmap = shunXing(rzhiArr)
		case "女":
			xmap = niXing(rzhiArr)
		}
	case "巳", "酉", "丑": //未
		rzhiArr := sortArr("未", zhiArr)
		switch sex {
		case "男":
			xmap = shunXing(rzhiArr)
		case "女":
			xmap = niXing(rzhiArr)
		}
	}
	return xmap
}

//男顺行十二辰 传入排序的地支宫位数组 返回宫位对应的小限年龄数组
func shunXing(rzhiArr []string) map[string][]int {
	var arr []string
	for j := 1; j <= 120; j++ {
		for i := 0; i < len(rzhiArr); i++ {
			if j%12 == 0 {
				s := fmt.Sprintf("%s:%d", rzhiArr[11], j)
				arr = append(arr, s)
				break
			}
			if i == j%12-1 && j%12 != 0 {
				s := fmt.Sprintf("%s:%d", rzhiArr[i], j)
				arr = append(arr, s)
				break
			}
		}
	}

	var xmap = make(map[string][]int) //k地支 v小限年龄数字
	var arrn []int
	for i := 0; i < len(rzhiArr); i++ {
		z := rzhiArr[i]
		for j := range arr {
			if strings.HasPrefix(arr[j], z) {
				arrn = append(arrn, j+1)
			}
		}
	}

	xmap[rzhiArr[0]] = arrn[:10]
	xmap[rzhiArr[1]] = arrn[10:20]
	xmap[rzhiArr[2]] = arrn[20:30]
	xmap[rzhiArr[3]] = arrn[30:40]
	xmap[rzhiArr[4]] = arrn[40:50]
	xmap[rzhiArr[5]] = arrn[50:60]
	xmap[rzhiArr[6]] = arrn[60:70]
	xmap[rzhiArr[7]] = arrn[70:80]
	xmap[rzhiArr[8]] = arrn[80:90]
	xmap[rzhiArr[9]] = arrn[90:100]
	xmap[rzhiArr[10]] = arrn[100:110]
	xmap[rzhiArr[11]] = arrn[110:120]
	return xmap
}

//女逆行十二辰
func niXing(rzhiArr []string) map[string][]int {
	//逆排
	head := rzhiArr[:1]
	end := rzhiArr[1:]
	var rArr []string
	for i := len(end) - 1; i >= 0; i-- {
		rArr = append(rArr, end[i])
	}
	rArr = append(head, rArr...)
	return shunXing(rArr)
}
