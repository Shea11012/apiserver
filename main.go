package main

import (
	"apiserver/config"
	"apiserver/logger"
	"apiserver/model"
	version2 "apiserver/pkg/version"
	"apiserver/router"
	"apiserver/router/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

// @title Apiserver Example API
// @version 1.0
// @description apiserver demo

// @contact.name mxy
// @contact.url http://www.swagger.io/suport
// @contact.email 1872314654@qq.com

// @host localhost:8080
// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	pflag.Parse()

	if *version {
		v := version2.Get()
		marshalled, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 加载日志
	logger.LoadLogger(zap.DebugLevel, "./logs/example.log")

	// 加载数据库
	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{middleware.RequestId()}

	router.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			zap.S().Infof("The router has no response, or it might took too long to start up. %v\n", err)
		}
	}()

	addr := viper.GetString("tls.addr")
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			zap.L().Info("https listening addr", zap.String("addr", addr))
			zap.S().Info(http.ListenAndServeTLS(addr, cert, key, g).Error())
		}()
	}
	zap.S().Infof("listening addr: %s", viper.GetString("addr"))
	zap.S().Infof("addr", http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// api 健康状态自检
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + viper.GetString("addr") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		zap.S().Info("waiting for the router,retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
