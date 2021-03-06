package handler

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var cache map[string]string
var client http.Client

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	url := r.FormValue("q")
	url = strings.Replace(url, ":/h", "://h", 1)
	url = strings.Replace(url, "https://", "http://", 1)
	if !strings.HasPrefix(url, "http://h5.cyol.com/special/daxuexi/") {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "错误输入")
		return
	}
	fmt.Println(url)

	temp := strings.Split(url, "/")
	temp = temp[:len(temp)-1]
	path := strings.Join(temp, "/")

	var title string
	if v, ok := cache[path]; ok {
		title = v
	} else {
		res, err := client.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, "服务异常：请求失败 "+err.Error())
			return
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, "服务异常：请求错误 "+strconv.Itoa(res.StatusCode))
			return
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, "服务异常：内容解析错误")
			return
		}
		title = doc.Find("title").Text()
		if strings.Contains(title, "“青年大学习”") {
			cache[path] = title
		}
	}

	path += "/images/end.jpg"
	_, _ = fmt.Fprintf(w, `<html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1, user-scalable=no"><title>%s</title></head><body style="margin:0"><div style="width:100vw;height:100vh;background-image: url(%s);background-size: 100%% 100%%;"></div></body></html>`, title, path)
}

func init() {
	dialer := &net.Dialer{
		Timeout:   6 * time.Second,
		KeepAlive: 3 * time.Second,
	}
	transport := http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, network, "218.1.70.80:80")
		}}
	client = http.Client{Transport: &transport}
	cache = make(map[string]string)
}
