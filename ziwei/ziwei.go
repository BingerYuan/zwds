package ziwei

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/starainrt/astro/basic"
)

type ZiWei struct {
	Solar string `json:"solar"` //阳历信息
	Gzs   string `json:"gzs"`   //干支信息
	Moon  string `json:"moon"`  //阴历信息

	YinShou  []string          `json:"yin_shou"`  //寅首(月干支)　0亥月　1子月
	MingGong string            `json:"ming_gong"` //命宫
	ShenGong string            `json:"shen_gong"` //身宫
	MingPan  map[string]string `json:"ming_pan"`  //命盘 k星名　v地支
	MingZhu  string            `json:"ming_zhu"`  //命主
	ShenZhu  string            `json:"shen_zhu"`  //身主
	WuXing   string            `json:"wu_xing"`   //五行局

	Star14      map[string]string `json:"star14"`        //14主星
	ChangSheng  map[string]string `json:"chang_sheng"`   //十二长生
	MonthStar   map[string]string `json:"month_star"`    //月系诸星
	YearZhiStar map[string]string `json:"year_zhi_star"` //年支系诸星
	YearGanStar map[string]string `json:"year_gan_star"` //年干系诸星
	HourStar    map[string]string `json:"hour_star"`     //时系诸星
	SiHua       map[string]string `json:"si_hua"`        //四化

	DaXian   map[string]string `json:"da_xian"`   //大限
	XiaoXian map[string][]int  `json:"xiao_xian"` //小限

	OtherStar map[string]string `json:"other_star"` //其他星

	Miao map[string]string `json:"miao"` //庙...
}

func NewZiWei(y, m, d, h int, sex string) *ZiWei {
	obj := gz.NewGanZhi(y,m,d,h)
	ygz:=obj.YGZ
	mgz:=obj.MGZ
	dgz:=obj.DGZ
	hgz:=obj.HGZ
	//
	lm, lday, _, _ := basic.GetLunar(y, m, d) //阴历日数字
	solar := fmt.Sprintf("%d年-%d月-%d日-%d时", y, m, d, h)
	gzs := fmt.Sprintf("%s年 %s月 %s日 %s时", ygz, mgz, dgz, hgz)
	moon := fmt.Sprintf("阴历 %d月 %d日", lm, lday)
	//
	yinShouArr := GetYinShouArr(ygz)
	mg, sg, mpMap := GetMsgMap(mgz, hgz)   //命宫　身宫　命盘十二辰
	mingZhu, shenZhu := GetMsgZhu(ygz, mg) //命主　身主
	wxn, wxs := GetWuXing(mg, yinShouArr)
	wx := fmt.Sprintf("五行:%s%d局", wxs, wxn) //五行局
	//14主星
	zwzhi := GetZiWei(wxn, lday)
	star14map := GetStar14Map(zwzhi) //14主星
	//十二长生
	changShengMap := GetChangShengMap(wxs, sex)
	//月系诸星
	monthStarmap := GetMonthStarMap(mgz)
	//年支系诸星
	yearZhiStarmap := GetYearZhiStarMap(ygz)
	//年干系诸星
	yearGanStarmap := GetYearGanStarMap(ygz)
	//时系诸星
	hourStarmap := GetHourStarMap(hgz)
	//四化
	siHuamap := GetSiHuaMap(ygz, mgz, hgz, star14map)
	//大小限
	daXianmap := GetDaXianMap(ygz, mgz, hgz, sex)
	xiaoXianmap := GetXiaoXianMap(ygz, sex)
	//其他星
	otherStarmap := GetOtherMap(ygz, mgz, hgz, lday)
	//庙
	miaoMap := GetMiaoMap(ygz, hgz, star14map)

	zw := &ZiWei{
		Solar:    solar,
		Gzs:      gzs,
		Moon:     moon,
		YinShou:  yinShouArr,
		MingGong: mg,
		ShenGong: sg,
		MingPan:  mpMap,
		MingZhu:  mingZhu,
		ShenZhu:  shenZhu,
		WuXing:   wx,

		Star14:      star14map,
		ChangSheng:  changShengMap,
		MonthStar:   monthStarmap,
		YearZhiStar: yearZhiStarmap,
		YearGanStar: yearGanStarmap,
		HourStar:    hourStarmap,
		SiHua:       siHuamap,

		DaXian:   daXianmap,
		XiaoXian: xiaoXianmap,

		OtherStar: otherStarmap,
		Miao:      miaoMap,
	}
	return zw
}
