package sopdelay

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var httpServer *http.Server

func router() http.Handler {
	r := gin.New()
	return r
}

// Serve 启动服务
// @title 账户服务(mk-plan-center)
// @version 1.0
// @description 示例服务
// @host mk-dev.dustess.com
// @BasePath /mk_plan_center
func Serve() error {

	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", "127.0.0.1", "5000"),
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return httpServer.ListenAndServe()
}

// Shutdown 关闭服务
func Shutdown() error {
	fmt.Println("正在关闭http服务")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		return err
	}
	fmt.Println("http服务成功关闭")
	return nil
}
