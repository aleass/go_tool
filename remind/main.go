package main

import (
	"fmt"
	"github.com/go-toast/toast"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var user = "gua"
var count int64

func webServer() {
	var t1 = template.New("html")
	t1, err := t1.Parse(`<script type="text/javascript">window.open('','_self');window.close();</script>`)
	if err != nil {
		println(err)
		os.Exit(2)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			println(err)
			return
		}
		_ = t1.Execute(w, nil)
		count, err = Drinking()
		if err != nil {
			println(err)
		}
		println(user+"已喝水: "+strconv.Itoa(int(count)), time.Now().Format("2006/01/02 15:04:05"))
	})
	err = http.ListenAndServe(":1111", nil)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
}

func timer(t chan struct{}) {
	for {
		n := time.Now().Hour()
		if n < 9 || n > 20 {
			goto sleep
		}
		t <- struct{}{}
	sleep:
		time.Sleep(time.Minute * 15)
	}
}

func drinkInit() {
	drink, err := GetDrinkInfo(time.Now().Format("2006-01-02"))
	if err != nil {
		println(err.Error())
		return
	}
	count = drink.DrinkCount
}

func main() {
	println("enter默认gua，或 请输入用户名：")
	_, _ = fmt.Scanln(&user)
	println("当前用户: " + user)
	var c = make(chan os.Signal)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	var t = make(chan struct{})
	go webServer()
	go timer(t)
	path, _ := os.Getwd()
	path += "\\logo.png"
	drinkInit()
	for {
		select {
		case <-c:
			println("喝水量:" + strconv.Itoa(int(count)))
			_ = Db.Close()
			os.Exit(1)
		case <-t:
			notification := toast.Notification{
				AppID:   "Microsoft.Windows.Shell.RunDialog",
				Title:   "吨吨吨",
				Message: "喝水啦！！！当前喝水次数是：" + strconv.Itoa(int(count)),
				Icon:    path, // 文件必须存在
				Actions: []toast.Action{
					{"protocol", "吨吨吨", "http://127.0.0.1:1111/"},
				},
			}
			err := notification.Push()
			if err != nil {
				println(err.Error())
				return
			}
		}
	}
}
