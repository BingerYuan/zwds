package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/starainrt/astro/basic"
	"liangzi.local/qx/pkg/cal"
	"liangzi.local/qx/pkg/zwds"
)

//默认8111 指定端口 -l 9999
func main() {
	flag.Usage = func() {
		fmt.Println("zwds -l port")
		flag.PrintDefaults()
	}
	p := flag.Int("l", 8111, "")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	fmt.Println("Default Server Port:", 8111)
	Main(port)
}
func Main(port string) {
	http.HandleFunc("/zw", ziWeiDouShu)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("listenAndServer:", err)
	}
}

//######################################################
//紫微斗数
//######################################################
func ziWeiDouShu(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("zwds.html")
		err := t.Execute(w, nil)
		if err != nil {
			return
		}
	} else {
		handZwds(w, r)
	}
}
func handZwds(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	t, sex := getZwInfo(r)
	zwBody := NewZwBody(t, sex)
	b := ResultZw(zwBody)
	json.NewEncoder(w).Encode(string(b))
}
func getZwInfo(r *http.Request) (time.Time, string) {
	//年
	ly, err := strconv.Atoi(r.Form["ly"][0])
	if err != nil {
		log.Fatalln("年异常:", err)
	}
	//月
	lm, err := strconv.Atoi(r.Form["lm"][0])
	if err != nil {
		log.Fatalln("月异常: ", err)
	}

	//日
	ld, err := strconv.Atoi(r.Form["ld"][0])
	if err != nil {
		log.Fatalln("日异常:", err)
	}
	//时辰 子时1 丑时2 寅时3...
	lh, err := strconv.Atoi(r.Form["lh"][0])
	if err != nil {
		log.Fatalln("时辰异常:", err)
	}
	//性别
	sex := r.Form["sex"][0]
	t := time.Date(ly, time.Month(lm), ld, lh, 0, 0, 0, time.Local)
	return t, sex
}

type ZwBody struct {
	*zwds.ZW
	Solar string `json:"solar"` //阳历时间
	GZ    string `json:"gz"`    //干支
	Moon  string `json:"moon"`  //阴历
	Info  string `json:"info"`  //命局
}

func NewZwBody(t time.Time, sex string) *ZwBody {
	y, m, d, h := t.Year(), int(t.Month()), t.Day(), t.Hour()
	obj := cal.NewGanZhiInfo(y, m, d, h)
	ygz := obj.YearGZ
	mgz := obj.MonthGZ
	dgz := obj.DayGZ
	hgz := obj.HourGZ
	lm, lday, _, _ := basic.GetLunar(y, m, d) //阴历日数字
	ymdhs := fmt.Sprintf("%d年-%d月-%d日-%d时", y, m, d, h)
	gzs := fmt.Sprintf("%s年 %s月 %s日 %s时", ygz, mgz, dgz, hgz)
	moons := fmt.Sprintf("阴历 %d月 %d日", lm, lday)
	//
	zw := zwds.NewZW(y, m, d, h, sex)
	info := fmt.Sprintf("性别:%s 命宫:%s 身宫:%s 五行:%s", sex, zw.MingGong, zw.ShenGong, zw.WuXingS)

	return &ZwBody{
		zw,
		ymdhs,
		gzs,
		moons,
		info,
	}
}
func ResultZw(zwds *ZwBody) []byte {
	b, err := json.Marshal(zwds)
	if err != nil {
		fmt.Println("zwds:", err)
	}
	return b
}
