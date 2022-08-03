package http

import (
	"context"
	"fmt"
	"git.dustess.com/mk-base/gin-ext/constant"
	"git.dustess.com/mk-base/gin-ext/middleware"
	interfaceAuth "git.dustess.com/mk-base/gin-ext/middleware/interface-auth"
	"git.dustess.com/mk-biz/mk-plan-center/infrastructure/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ginSwaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"os"
	"time"
)

var httpServer *http.Server

func router() http.Handler {
	r := gin.New()

	// 跨域中间件
	r.Use(middleware.CORS())

	// 请求日志
	r.Use(middleware.Logger())
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
	}
	// validator 信息翻译
	uni := middleware.NewZHUNI()
	binding.Validator = middleware.NewV9Validator(uni)
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
	}
	// validator 错误处理
	r.Use(middleware.NewErrorHandler(uni).HandleErrors)
	// 注册链路追踪
	r.Use(middleware.WithTrace())
	// 错误恢复
	r.Use(gin.Recovery())
	// 监控
	prometheus := middleware.NewPrometheus("mk_plan_center")
	prometheus.Use(r)

	r.Use(interfaceAuth.InterfaceAuth(config.Get().App))
	err := initRouter(r)
	if err != nil {
		fmt.Errorf("init router %v", err)
		return r
	}

	// 错误恢复
	r.Use(gin.Recovery())

	// release 模式下不提供接口文档访问
	if string(constant.ReleaseMode) != os.Getenv("GIN_MODE") {
		r.GET(v1prefix+"/swagger/*any", ginSwagger.WrapHandler(ginSwaggerFiles.Handler))
	}

	return r
}

// Serve 启动服务
// @title 账户服务(mk-plan-center)
// @version 1.0
// @description 示例服务
// @host mk-dev.dustess.com
// @BasePath /mk_plan_center
func Serve() error {
	conf := *config.Get()
	fmt.Sprintf("正在启动http服务，监听端口%v", conf.Server.Port)
	httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port),
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
