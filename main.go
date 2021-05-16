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
		},
	}
}
