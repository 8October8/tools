package climap

import (
	"mycode/climap/frame"
	"mycode/climap/plugins"
)

var Cli frame.App

//菜单初始化
func init()  {
	Cli.Name = "tools"
	Cli.Explain = "自写学习golang cli程序框架"
	Cli.Version = "0.1.0"
	Cli.Commands = []frame.Command{
		{Name: "portscan", Usage: "多线程端口开放扫描器", Version: "0.1.0"},
	}
}

func Climap() {
	switch  {
	case frame.InAppCommand(Cli,"portscan"):
		plugins.PortScan()
	default:
		frame.AppHelp(Cli)
	}
}
