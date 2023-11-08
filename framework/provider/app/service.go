package app

import (
	"errors"
	"github.com/1152545264/goWebSelf/framework"
	"github.com/1152545264/goWebSelf/framework/util"
	"path/filepath"
)

type HadeApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础路径
	appId      string              //表示当前这个app的唯唯一id，可用于分布式锁等

	configMap map[string]string //配置加载
}

// AppID 表示这个app的唯一ID
func (app HadeApp) AppID() string {
	return app.appId
}

// Version 实现版本
func (app HadeApp) Version() string {
	return "0.0.3"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (app HadeApp) BaseFolder() string {
	if app.BaseFolder() != "" {
		return app.baseFolder
	}
	// 如果参数也没有，则默认使用当前路径
	return util.GetExecDirectory()
}

// ConfigFolder 表示配置文件地址
func (app HadeApp) ConfigFolder() string {
	if val, ok := app.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (app HadeApp) LogFolder() string {
	if val, ok := app.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "log")
}

func (app HadeApp) HttpFolder() string {
	if val, ok := app.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "http")
}

func (app HadeApp) ConsoleFolder() string {
	if val, ok := app.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "console")
}

func (app HadeApp) StorageFolder() string {
	if val, ok := app.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (app HadeApp) ProviderFolder() string {
	if val, ok := app.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app", "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (app HadeApp) MiddlewareFolder() string {
	if val, ok := app.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(app.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (app HadeApp) CommandFolder() string {
	if val, ok := app.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(app.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (app HadeApp) RuntimeFolder() string {
	if val, ok := app.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (app HadeApp) TestFolder() string {
	if val, ok := app.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "test")
}

// DeployFolder 定义测试需要的信息
func (app HadeApp) DeployFolder() string {
	if val, ok := app.configMap["deploy_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "deploy")
}

// NewHadeApp 初始化HadeApp
func NewHadeApp(params ...any) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	//有两个参数 容器和baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return HadeApp{baseFolder: baseFolder, container: container}, nil
}
