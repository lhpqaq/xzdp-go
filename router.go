// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "xzdp/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.StaticFS("/imgs", &app.FS{Root: "upload/", GenerateIndexPages: false})
}
