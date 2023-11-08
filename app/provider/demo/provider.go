package demo

import "github.com/1152545264/goWebSelf/framework"

type DemoProvider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *DemoProvider) Name() string {
	return DemoKey
}

func (sp *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

func (sp *DemoProvider) IsDefer() bool {
	return false
}

func (sp *DemoProvider) Params(c framework.Container) []any {
	return []any{sp.c}
}

func (sp *DemoProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}
