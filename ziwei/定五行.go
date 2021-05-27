package ziwei

import (
	"github.com/Aquarian-Age/xa/pkg/gz"
	"strings"
)

//定五行 传年干支 命宫地支　月干支 返回命宫五行数字 五行属性
func GetWuXing(mg string, mgzArr []string) (int, string) {
	return wuXing(mg, mgzArr)
}
func wuXing(mg string, mgzArr []string) (int, string) {
	//水二局、木三局、金四局、土五局、火六局
	wxMap := map[int]string{2: "水", 3: "木", 4: "金", 5: "土", 6: "火"}
	//顺布宫干到命宫地支位
	var wxN int    //纳因数字
	var wxS string //五行属性
	for j := 0; j < len(mgzArr); j++ {
		if strings.ContainsAny(mg, mgzArr[j]) {
			nys := gz.GetNaYin(mgzArr[j]) //纳因
			for k, v := range wxMap {
				if strings.ContainsAny(nys, v) {
					wxN = k
					wxS = v
					break
				}
			}
			break
		}
	}
	return wxN, wxS
}
