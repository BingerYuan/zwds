package ziwei

import (
	"fmt"
	"testing"
)

func TestGetChangShengMap(t *testing.T) {
	wxs := "木"
	sex := "男" //map[丑:冠带 亥:长生 午:死 卯:帝旺 子:沐浴 寅:临官 巳:病 戌:养 未:墓 申:绝 辰:衰 酉:胎]
	sex = "女"  //map[丑:胎 亥:长生 午:衰 卯:墓 子:养 寅:绝 巳:病 戌:沐浴 未:帝旺 申:临官 辰:死 酉:冠带]
	csmap := GetChangShengMap(wxs, sex)
	fmt.Println(csmap)
}
