package ziwei

import (
	"fmt"
	"strings"
)

type InfoZw struct {
	SiHuaY string `json:"si_hua_y"` //年干四化
	SiHuaM string `json:"si_hua_m"` //月干四化
	SiHuaD string `json:"si_hua_d"` //日干化忌
}

func (zw *ZiWei) NewInfoZw() *InfoZw {

	siHuaYgz := getSiHuaXgz(zw.Ygz)
	siHuaMgz := getSiHuaXgz(zw.Mgz)
	siHuaDgz := getSiHuaXgz(zw.Dgz)

	izw := &InfoZw{
		SiHuaY: siHuaYgz,
		SiHuaM: siHuaMgz,
		SiHuaD: siHuaDgz,
	}
	return izw
}

//年月干四化
func getSiHuaXgz(mgz string) string {
	hl, hls := GetHuaLu(mgz)
	huaLus := fmt.Sprintf("%s:%s ", hls, hl)
	hq, hqs := GetHuaQuan(mgz)
	huaQuans := fmt.Sprintf("%s:%s ", hqs, hq)
	hk, hks := GetHuaKe(mgz)
	huaKes := fmt.Sprintf("%s:%s ", hks, hk)
	hj, hjs := GetHuaJi(mgz)
	huaJis := fmt.Sprintf("%s:%s ", hjs, hj)
	s := fmt.Sprintf("%s %s %s %s %s", mgz, huaLus, huaQuans, huaKes, huaJis)
	return s
}

//化禄星
func GetHuaLu(ygz string) (string, string) {
	hls := huaLuF(ygz)
	return "化禄", hls
}

//化禄 廉贞星 天机星	天同星	太阴星	贪狼星	武曲星	太阳星	巨门星	天梁星	破军星
func huaLuF(ygz string) string {
	gmap := map[string]string{
		"甲": "廉贞星", "乙": "天机星",
		"丙": "天同星", "丁": "太阴星",
		"戊": "贪狼星", "己": "武曲星",
		"庚": "太阳星", "辛": "巨门星",
		"壬": "天梁星", "癸": "破军星",
	}
	var hls string
	for k, v := range gmap {
		if strings.ContainsAny(ygz, k) {
			hls = v
			break
		}
	}
	return hls
}

//化权星
func GetHuaQuan(ygz string) (string, string) {
	hqs := huaQuanF(ygz)
	return "化权", hqs
}

//化权星	破军星	天梁星	天机星	天同星	太阴星	贪狼星	武曲星	太阳星	紫微星	巨门星
func huaQuanF(ygz string) string {
	gmap := map[string]string{
		"甲": "破军星", "乙": "天梁星",
		"丙": "天机星", "丁": "天同星",
		"戊": "太阴星", "己": "贪狼星",
		"庚": "武曲星", "辛": "太阳星",
		"壬": "紫微星", "癸": "巨门星",
	}
	var hqs string
	for k, v := range gmap {
		if strings.ContainsAny(ygz, k) {
			hqs = v
			break
		}
	}
	return hqs
}

//化科星
func GetHuaKe(ygz string) (string, string) {
	hks := huaKeF(ygz)
	return "化科", hks
}

//化科星	武曲星	紫微星	文昌星	天机星	右弼星	天梁星	太阴星	文曲星	左辅星	太阴星
func huaKeF(ygz string) string {
	gmap := map[string]string{
		"甲": "武曲星", "乙": "紫微星",
		"丙": "文昌星", "丁": "天机星",
		"戊": "右弼星", "己": "天梁星",
		"庚": "太阴星", "辛": "文曲星",
		"壬": "左辅星", "癸": "太阴星",
	}
	var hks string
	for k, v := range gmap {
		if strings.ContainsAny(ygz, k) {
			hks = v
			break
		}
	}
	return hks
}

//化忌星
func GetHuaJi(ygz string) (string, string) {
	hjs := huaJiF(ygz)
	return "化忌", hjs
}

//化忌星	太阳星	太阴星	廉贞星	巨门星	天机星	文曲星	天同星	文昌星	武曲星	贪狼星
func huaJiF(ygz string) string {
	gmap := map[string]string{
		"甲": "太阳星", "乙": "太阴星",
		"丙": "廉贞星", "丁": "巨门星",
		"戊": "天机星", "己": "文曲星",
		"庚": "天同星", "辛": "文昌星",
		"壬": "武曲星", "癸": "贪狼星",
	}
	var hjs string
	for k, v := range gmap {
		if strings.ContainsAny(ygz, k) {
			hjs = v
			break
		}
	}
	return hjs
}
