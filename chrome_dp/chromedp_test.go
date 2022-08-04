package chrome_dp

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"testing"
)

func TestChrome(t *testing.T) {
	url := "https://www.baidu.com"
	var buf []byte
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	if ac := chromedp.Run(ctx, fullScreenshot(url, 100, &buf)); ac != nil {
		t.Errorf("RunSnapShop --- 截屏失败err = [%v]", ac)
	}
	err := ioutil.WriteFile("picture.png", buf, 0777)
	if err != nil {
		t.Errorf("写文件报错---[%s]", err.Error())
	}
}
