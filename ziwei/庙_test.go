package ziwei

import (
	"fmt"
	"testing"
)

func TestGetMiaoMap(t *testing.T) {
	//2021年-5月-26日-10时
	star14map := map[string]string{
		"七杀星": "午", "天同星": "亥", "天府星": "子", "天机星": "卯", "天梁星": "巳", "天相星": "辰", "太阳星": "丑", "太阴星": "丑", "巨门星": "卯", "廉贞星": "申", "武曲星": "子", "破军星": "戌", "紫微星": "辰", "贪狼星": "寅",
	}
	ygz := "甲子"
	hgz := "甲子"
	mmap := GetMiaoMap(ygz, hgz, star14map)
	fmt.Printf("庙map:%v\n", mmap)
	/*
		七杀星:平 天同星:庙 天府星:地 天机星:平 天梁星:陷 天相星:地 太阳星:旺
		太阴星:陷 巨门星:平 廉贞星:陷 擎羊星:利 文昌:庙 文曲:庙 武曲星:平
		火星:地 破军星:平 紫微星:旺 贪狼星:陷 铃星:地 陀罗星:陷
	*/
}
