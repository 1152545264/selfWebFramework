package demo

import (
	"fmt"
	"github.com/1152545264/goWebSelf/framework"
)

type DemoService struct {
	//实现接口
	Service

	//参数
	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	//参数在此展开
	c := params[0].(framework.Container)
	fmt.Println("new demo service")

	//返回实例
	return &DemoService{c: c}, nil
}

// 实现接口
func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}
