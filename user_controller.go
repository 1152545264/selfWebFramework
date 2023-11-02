package main

import (
	"github.com/1152545264/goWebSelf/framework"
)

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}
