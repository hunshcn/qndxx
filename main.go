package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	r.URL.Path = strings.Replace(r.URL.Path[1:], ":/", "://", 1)
	if !strings.HasPrefix(r.URL.Path, "http") {
		_, _ = io.WriteString(w, "错误输入")
		return
	}
	// 请求html页面
	res, err := http.Get(r.URL.Path)
	if err != nil {
		// 错误处理
		_, _ = io.WriteString(w, "服务异常：请求失败")
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		_, _ = io.WriteString(w, "服务异常：请求错误")
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		_, _ = io.WriteString(w, "服务异常：内容解析错误")
		return
	}
	title := doc.Find("title").Text()
	temp := strings.Split(r.URL.Path, "/")
	temp = temp[:len(temp)-1]
	path := strings.Join(temp, "/") + "/images/end.jpg"
	_, _ = fmt.Fprintf(w, `<html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1, user-scalable=no"><title>%s</title></head><body style="margin:0"><div style="width:100vw;height:100vh;background-image: url(%s);background-size: 100%% 100%%;"></div></body></html>`, title, path)
}

func main() {

	http.HandleFunc("/", index)

	_ = http.ListenAndServe(":8090", nil)
}
