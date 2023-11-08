package http

import (
	"github.com/1152545264/goWebSelf/app/http/module/demo"
	"github.com/1152545264/goWebSelf/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
