package my_url

import (
	"fmt"
	"net/url"
)

func urlDecode(urlStr string) {
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println("编码:", escapeUrl)
}

func urlEncode(urlStr string) {
	enEscapeUrl, _ := url.QueryUnescape(urlStr)
	fmt.Println("解码:", enEscapeUrl)
}

// UrlParam url请求参数
type UrlParam struct {
	ID string `json:"id"` // ID
}

// GetUrl 拼接URL
func GetUrl(param *UrlParam) string {
	ahp := ""
	linkURL, err := url.ParseRequestURI(ahp)
	if err != nil {
		return ""
	}
	q := linkURL.Query()
	q.Set("materialId", param.ID)
	linkURL.RawQuery = q.Encode()
	agentHomePage, err := url.QueryUnescape(linkURL.String())
	if err != nil {
		return ""
	}
	return agentHomePage
}
