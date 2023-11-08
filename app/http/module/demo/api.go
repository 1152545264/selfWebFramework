package demo

import (
	demoService "github.com/1152545264/goWebSelf/app/provider/demo"
	"github.com/1152545264/goWebSelf/framework/contract"
	"github.com/1152545264/goWebSelf/framework/gin"
	"net/http"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/deo_post", api.DemoPost)
	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

func (api *DemoApi) Demo(c *gin.Context) {
	appService := c.MustMake(contract.AppKey).(contract.App)
	baseFolder := appService.BaseFolder()
	c.JSONP(http.StatusOK, baseFolder)
}
func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(http.StatusOK, usersDTO)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(http.StatusOK, nil)
}
