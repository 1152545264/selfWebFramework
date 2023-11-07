package main

import (
	"fmt"
	"github.com/1152545264/goWebSelf/framework/gin"
	"github.com/1152545264/goWebSelf/framework/provider/demo"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

// SubjectListController 对应路由 /subject/list/all
func SubjectListController(c *gin.Context) {
	// h获取demo服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)
	//调用服务实例的方法
	foo := demoService.GetFoo()

	//输出结果
	c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))

}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
