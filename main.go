package main

import (
	"github.com/yongcycchen/mall-users/startup"
	"github.com/yongcycchen/mall-users/vars"
)

const APP_NAME = "mall-users"

func main() {
	application := &vars.GRPCApplication{
		Application: &vars.Application{
			LoadConfig: startup.LoadConfig,
			SetupVars:  startup.SetupVars,
			Name:       APP_NAME,
		},
		RegisterGRPCServer: startup.RegisterGRPCServer,
		RegisterGateway:    startup.RegisterGateway,
		RegisterHttpRoute:  startup.RegisterHttpRoute,
	}
	vars.App = application

}
