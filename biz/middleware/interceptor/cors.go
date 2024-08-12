package interceptor

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"xzdp/conf"
)

// AllowAllCors 直接放行所有跨域请求并放行所有 OPTIONS 方法
func allowAllCors(ctx context.Context, c *app.RequestContext) {
	method := c.Request.Method()
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	if string(method) == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next(ctx)
	hlog.Debugf("AllowCors:Method:%+v,Path:%+v", string(c.Request.Method()), string(c.Request.Path()))
}

// Cors 按照配置处理跨域请求
func Cors(ctx context.Context, c *app.RequestContext) {
	mode := conf.GetConf().Cors.Mode
	// 放行全部
	if mode == "allow-all" {
		allowAllCors(ctx, c)
		return
	}
	whitelist := checkCors(string(c.GetHeader("origin")))
	// 通过检查, 添加请求头
	if whitelist != nil {
		c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
		c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
		c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
		c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
		if whitelist.AllowCredentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}
	}

	// 严格白名单模式且未通过检查，直接拒绝处理请求
	if whitelist == nil && mode == "strict-whitelist" && !(string(c.Request.Method()) == "GET") {
		c.AbortWithStatus(http.StatusForbidden)
		return
	} else {
		// 非严格白名单模式，无论是否通过检查均放行除了所有 OPTIONS 方法
		if string(c.Request.Method()) == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}

	// 处理请求
	c.Next(ctx)

}

func checkCors(currentOrigin string) *conf.CORSWhitelist {
	for _, whitelist := range conf.GetConf().Cors.WhiteList {
		// 遍历配置中的跨域头，寻找匹配项
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
