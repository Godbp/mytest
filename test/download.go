package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// GetImg ...
func GetImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	_ = ioutil.WriteFile("pic.jpg", pix, 0755)

	return

}
