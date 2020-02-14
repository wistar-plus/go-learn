package middleware

import (
	"bytes"
	"go-learn/pkg/logger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AccessLogMiddleware AccessLogMiddleware
func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var url, param string
		ip := c.ClientIP()
		method := c.Request.Method
		proto := c.Request.Proto
		ua := c.GetHeader("User-Agent")
		form := c.Request.PostForm.Encode()

		urls := strings.Split(c.Request.RequestURI, "?")
		switch len(urls) {
		case 2:
			param = urls[1]
			fallthrough
		case 1:
			url = urls[0]
		}

		if method == http.MethodPost || method == http.MethodPut {
			body, err := ioutil.ReadAll(c.Request.Body)
			c.Request.Body.Close()
			if err == nil {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				form = string(body)
			}
		}
		//TODO: uid ,uname
		c.Next()
		//TODO: response code,msg

		logger.AccessLog(
			zap.Time("t", start),
			zap.String("ip", ip),
			zap.String("method", method),
			zap.String("url", url),
			zap.String("param", param),
			zap.String("proto", proto),
			zap.String("ua", ua),
			//	zap.String("uid", uid),
			//	zap.String("uname", uname),
			zap.String("form", form),
			//	zap.String("res_code", resCode),
			//	zap.String("res_msg", resMsg),
			zap.Int64("cost_time", time.Since(start).Nanoseconds()/1e6),
		)
	}
}
