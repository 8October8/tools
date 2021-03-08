package main

import (
	"mycode/climap"
	"mycode/climap/frame"
)

var App frame.App

//菜单初始化
func init()  {
	App.Name = "tools"
	App.Explain = "自写学习golang程序"
	App.Version = "0.1.0"
	App.Commands = []frame.Command{
		{Name: "cli", Usage: "选择程序运行接入方式", Version: "0.1.0"},
		{Name: "web", Usage: "选择程序运行接入方式", Version: "0.1.0"},
		{Name: "gui", Usage: "选择程序运行接入方式", Version: "0.1.0"},
	}
}

func main() {
	switch  {
	case frame.InAppCommand(App,"cli"):
		climap.Climap()
	default:
		frame.AppHelp(App)
	}
}
