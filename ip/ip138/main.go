package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"time"
)

const (
	logHead  = "IP138:"
	ip138Url = "https://%d.ip138.com/"
)

func GetPublicIP() string {
	url := fmt.Sprintf(ip138Url, time.Now().Year())
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		PrintErr(err)
		return ""
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		PrintErr(err)
		return ""
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		PrintErr(err)
		return ""
	}

	ip := ""
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		ip = strings.TrimPrefix(s.Text(), "您的IP地址是：")
	})
	return ip
}

func PrintErr(err error) {
	fmt.Printf(logHead+"err:%s\n", err)
}

func main() {
	ip := GetPublicIP()
	fmt.Printf("your public ip is: %s\n", ip)
}
