package my_url

import "testing"

func TestDecode(t *testing.T) {
	urlStr := "https://www.baidu.com/s?ie=utf-8&tn=88093251_75_hao_pg&wd=proto%20golang%20mod"
	urlDecode(urlStr)
}

func TestEncode(t *testing.T) {
	urlStr := "https://www.baidu.com/s?ie=utf-8&tn=88093251_75_hao_pg&wd=proto%20golang%20mod"
	urlEncode(urlStr)
}
