package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/Aquarian-Age/zwds/ziwei"
)

func main() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("home.html")
		if err != nil {
			fmt.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			return
		}
	} else {
		hand(w, r)
	}
}
func hand(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	t, sex := getymdSex(r)
	body := newBody(t, sex)
	b, _ := json.Marshal(&body)
	err = json.NewEncoder(w).Encode(string(b))
	if err != nil {
		return
	}
}
func getymdSex(r *http.Request) (time.Time, string) {
	ly, err := strconv.Atoi(r.Form["ly"][0])
	if err != nil {
		log.Fatalln("年异常:", err)
	}
	lm, err := strconv.Atoi(r.Form["lm"][0])
	if err != nil {
		log.Fatalln("月异常: ", err)
	}
	ld, err := strconv.Atoi(r.Form["ld"][0])
	if err != nil {
		log.Fatalln("日异常:", err)
	}
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
	*ziwei.ZiWei
	//*ziwei.InfoZw
}

func newBody(t time.Time, sex string) *ZwBody {
	y, m, d, h := t.Year(), int(t.Month()), t.Day(), t.Hour()
	zw := ziwei.NewZiWei(y, m, d, h, sex)
	//zwinfo := zw.NewInfoZw()
	return &ZwBody{
		zw,
		//zwinfo,
	}
}

// func newBody(t time.Time, sex string) *ziwei.ZiWei {
// 	y, m, d, h := t.Year(), int(t.Month()), t.Day(), t.Hour()
// 	return ziwei.NewZiWei(y, m, d, h, sex)
// }
