package http

import (
	"go-learn/internal/app/wweb/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() *http.Server {

	gin.SetMode(viper.GetString("mode"))
	engine := gin.New()
	pprof.Register(engine)
	engine.Use(middleware.AccessLogMiddleware())

	route(engine)

	server := &http.Server{
		Addr:         viper.GetString("addr"),
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
