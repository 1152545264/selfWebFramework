package http

import (
	"github.com/1152545264/goWebSelf/framework/gin"
)

func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Routes(r)
	return r, nil
}
