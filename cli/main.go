package main

import (
	"fmt"

	"github.com/starainrt/astro/basic"
	"liangzi.local/qx/pkg/cal"
	"liangzi.local/qx/pkg/cmd"
	"liangzi.local/qx/pkg/zwds"
)

func main() {
	info := cmd.GetInPut()
	y := info.Y
	m := info.M
	d := info.D
	h := info.H
	sex := info.Sex
	fmt.Println(y, m, d, h, sex)
	obj := cal.NewGanZhiInfo(y, m, d, h)
	ygz := obj.YearGZ
	mgz := obj.MonthGZ
	hgz := obj.HourGZ
	lm, lday, _, _ := basic.GetLunar(y, m, d)
	fmt.Printf("%s年 %s月 阴历%d月%d日 %s时\n", ygz, mgz, lm, lday, hgz)
	zw := zwds.NewZW(y, m, d, h, sex)
	fmt.Printf("性别:%s 命宫:%s 身宫:%s 五行局:%s\n\n", sex, zw.MingGong, zw.ShenGong, zw.WuXingS)
	fmt.Printf("亥宫:%v\n", zw.Hai)
	fmt.Printf("子宫:%v\n", zw.Zi)
	fmt.Printf("丑宫:%v\n", zw.Chou)
	fmt.Printf("寅宫:%v\n", zw.Yin)
	fmt.Printf("卯宫:%v\n", zw.Mao)
	fmt.Printf("辰宫:%v\n", zw.Mao)
	fmt.Printf("巳宫:%v\n", zw.Si)
	fmt.Printf("午宫:%v\n", zw.Wu)
	fmt.Printf("未宫:%v\n", zw.Wei)
	fmt.Printf("申宫:%v\n", zw.Shen)
	fmt.Printf("酉宫:%v\n", zw.You)
	fmt.Printf("戌宫:%v\n", zw.Xu)
}
