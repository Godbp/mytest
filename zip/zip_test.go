package zip

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	err := OssZip("C:\\Users\\admin\\Desktop\\picture", "C:\\Users\\admin\\Desktop\\test.zip")
	if err != nil {
		fmt.Printf("压缩失败 [%v]", err)
	}

}
