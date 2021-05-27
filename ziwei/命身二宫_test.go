package ziwei

import (
	"fmt"
	"testing"
)

func TestGetMingShenGong(t *testing.T) {
	mgz, hgz := "壬辰", "癸亥"
	//mgz, hgz = "己丑", "己亥"
	mg, sg, msmap := GetMsgMap(mgz, hgz)
	fmt.Printf("命宫:%s 身宫:%s 命盘:%v\n", mg, sg, msmap)
}
