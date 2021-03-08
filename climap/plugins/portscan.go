package plugins

import (
	"fmt"
	"mycode/climap/frame"
	"mycode/common"
	"mycode/core"
	"os"
	"strconv"
)

// clidataset ：cli程序数据输入函数
func clidataset() frame.Command {
	var portscan frame.Command
	portscan.Name = "portscan"
	portscan.Usage = "多线程端口开放扫描器"
	portscan.Version = "0.1.0"
	portscan.Flags = []frame.Flag{
		{Name: "host", ShortName: "H", Value: "127.0.0.1", Usage: "检测的IP：（192.168.1.1）（192.168.1.1,192.168.1.2）（192.168.1.1/24）（192.168.1.1-192.168.1.254）（192.168.1.1-254）"},
		{Name: "port", ShortName: "P", Value: "21,22,80,135,443,445,3306,3389,5000,8080,10000", Usage: "检测的端口：（80）（80,81）（1-65535）(-65535)"},
		{Name: "thread", ShortName: "T", Value: "200",Coerce: true,Usage: "最大线程数"},
		{Name: "protocol", ShortName: "p", Value: "tcp",Coerce: true, Usage: "检测协议,tcp 或者 udp"},
		{Name: "timeout", ShortName: "t", Value: "1500",Coerce: true, Usage: "检测超时时间,毫秒"},
		{Name: "verbose", ShortName: "v", Value: "1",Coerce: true, Usage: "输出详细信息等级,1-6级"},
	}
	return portscan
}

// clihelp ：cli程序帮助菜单
func clihelp(portscan frame.Command) {
	if len(os.Args) == 3 {
		frame.CommandHelp(portscan)
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		if common.InList("help",os.Args,"on") || common.InList("--help",os.Args,"on") || common.InList("-help",os.Args,"on") {
			frame.CommandHelp(portscan)
			os.Exit(0)
		}
	}
}

// output ：数据输出函数
func output(proton []string,portoff []string,verbose int)  {
	if len(proton) != 0 || len(portoff) != 0 {
		fmt.Println()
		switch verbose {
		case 2:
			for _,address := range portoff {
				fmt.Printf("%-21s  关闭\n",address)
			}
			for _,address := range proton {
				fmt.Printf("%-21s  开放\n",address)
			}
		default:
			for _,address := range proton {
				fmt.Printf("%-21s  开放\n",address)
			}
		}
	} else {
		fmt.Println("无扫描数据结果")
	}
}

// PortScan ：多线程端口开放扫描器,主体函数
func PortScan() {
	portscan := clidataset() //cli程序数据输入
	clihelp(portscan)        //cli程序帮助菜单

	//参数处理
	host,err := frame.GetCommandFlag(portscan,"host")
	var hosts []string
	if err == nil {
		hosts = common.IPlist(host)
	}
	port,err := frame.GetCommandFlag(portscan,"port")
	var ports []string
	if err == nil {
		ports = common.PortList(port)
	}
	thread,err := frame.GetCommandFlag(portscan,"thread")
	var threads int
	if err == nil {
		threads,_ = strconv.Atoi(thread)
	}
	protocol,err := frame.GetCommandFlag(portscan,"protocol")
	var protocols string
	if err == nil {
		protocols = protocol
	}
	timeout,err := frame.GetCommandFlag(portscan,"timeout")
	var timeouts int
	if err == nil {
		timeouts,_ = strconv.Atoi(timeout)
	}
	verbose,err := frame.GetCommandFlag(portscan,"verbose")
	var verboses int
	if err == nil {
		verboses,_ = strconv.Atoi(verbose)
	}

	proton,portoff := core.PortScanCode(hosts,ports,threads,protocols,timeouts) //功能函数

	output(proton,portoff,verboses) //数据输出
}


