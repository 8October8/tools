//用于常见cli菜单参数
package frame

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

// App ：app是cli应用程序的主要结构
type App struct {
	// 程序的名称,默认为path.Base(os.Args[0])
	Name string
	// 程序的描述
	Explain string
	// 程序的使用描述
	Usage string
	// 程序的版本
	Version string
	// 程序banner信息
	Banner string
	// 命令列表
	Commands []Command
	// 要解析的标志列表
	Flags []Flag
}

// Command ：command是cli应用程序的独立命令结构
type Command struct {
	// 命令的名称
	Name string
	// 命令的使用描述
	Usage string
	// 命令的版本
	Version string
	// 要解析的标志列表
	Flags []Flag
}

// Flag ：flag是cli应用程序的参数结构
type Flag struct {
	// 标志的名称
	Name string
	// 标志的短名称。通常一个字符(不赞成，使用' Aliases ')
	ShortName string
	// 标志的默认值
	Value string
	// 标志的强制标识，表明该参数是否必须
	Coerce bool
	// 标志的使用描述
	Usage string
}

// GetAppFlag ：获取App(程序)中对应flag(选项)的参数值,输入：父属App类型数据、App(程序)的flag(选项) Name值,输出：参数值字符串
func GetAppFlag(app App,name string) (string,error) {
	for _,flag := range app.Flags {
		if flag.Name == name {
			for i, arg := range os.Args {
				if strings.ToLower(arg) == "--"+strings.ToLower(flag.Name) || arg == "-"+flag.ShortName { //获取参数标志位
					if len(os.Args) == i+1 {
						return flag.Value,nil
					} else {
						if InAppFlag(app,os.Args[i+1])  {
							return flag.Value,nil
						} else {
							return os.Args[i+1],nil
						}
					}
				}
			}
			if flag.Coerce {
				return flag.Value,nil
			} else {
				return "",errors.New("输入中，无此参数选项")
			}
		}
	}
	return "",errors.New("command(命令)无此预设flag(选项),请检查代码中输入command(命令)的flag(选项) Name值")
}

// GetCommandFlag ：获取command(命令)中对应flag(选项)的参数值,输入：父属command类型数据、command(命令)的flag(选项) Name值,输出：参数值字符串或默认值,若无此
func GetCommandFlag(command Command,name string) (string,error){
	for _,flag := range command.Flags {
		if flag.Name == name {
			for i, arg := range os.Args {
				if strings.ToLower(arg) == "--"+strings.ToLower(flag.Name) || arg == "-"+flag.ShortName { //获取参数标志位
					if len(os.Args) == i+1 {
						return flag.Value,nil
					} else {
						if InCommandFlag(command,os.Args[i+1])  {
							return flag.Value,nil
						} else {
							return os.Args[i+1],nil
						}
					}
				}
			}
			if flag.Coerce {
				return flag.Value,nil
			} else {
				return "",errors.New("输入中，无此参数选项")
			}
		}
	}
	return "",errors.New("command(命令)无此预设flag(选项),请检查代码中输入command(命令)的flag(选项) Name值")
}

// GetFlag ：获取对应flag的参数值,输入：为flag类型数据,输出：为参数值或value默认值
func GetFlag(flag Flag) string {
	for i, arg := range os.Args {
		if len(os.Args) == i+1 {
			return flag.Value
		} else {
			if  strings.ToLower(arg) == "--"+strings.ToLower(flag.Name) || arg == "-"+flag.ShortName {
				return os.Args[i+1]
			}
		}
	}
	return flag.Value
}

// GetNArg ：获取除预设程序参数后的其他参数,输入：为App类型数据,输出：为其他参数列表
func NArg(app App) []string {
	var arglist []string  //全部预设命令，参数，参数值

	//获取所有命令参数选项
	for i, arg := range os.Args {
		for _, flag := range app.Flags {
			name := "--"+ flag.Name
			shortname := "-"+ flag.ShortName
			if arg == name  {
				arglist = append(arglist,name)
				arglist = append(arglist,os.Args[i+1])
			}
			if arg == shortname {
				arglist = append(arglist,shortname)
				arglist = append(arglist,os.Args[i+1])
			}

		}
		for _,command := range app.Commands {
			if arg == command.Name {
				arglist = append(arglist,command.Name)
			}
			for _, flag := range command.Flags {
				name := "--"+ flag.Name
				shortname := "-"+ flag.ShortName
				if arg == name  {
					arglist = append(arglist,name)
					arglist = append(arglist,os.Args[i+1])
				}
				if arg == shortname {
					arglist = append(arglist,shortname)
					arglist = append(arglist,os.Args[i+1])
				}
			}
		}
	}
	var NArg []string
	for _,arg := range os.Args {
		for _,y := range arglist {
			if arg != y {
				NArg = append(NArg,arg)
			}
		}
	}
	return NArg
}


/*
参数判断系列
*/
// InAppCommand ：判断是否参数是否存在预设command(命令),输入：判断的父属App类型数据、判断的command Name值,输出：bool值
func InAppCommand(app App,name string) bool {
	for _,command := range app.Commands {
		if command.Name == name {
			for _, arg := range os.Args {
				if strings.ToLower(arg) == strings.ToLower(name) {
					return true
				}
			}
			return false
		}
	}
	fmt.Println("App(程序)无此预设command(命令),请检查代码中输入的command Name值")
	os.Exit(1)
	return false
}

// InCommand ：判断是否存在预设command,输入：为Command类型数据,输出：为bool值
func InCommand(command Command,name string) bool {
	if strings.ToLower(name) == strings.ToLower(command.Name)  {
		return true
	}
	return false
}

// InAppFlag ：判断App(程序)中是否存在对应flag(选项)的参数设置,输入：父属App类型数据、对应name值,输出：bool值
func InAppFlag(app App,name string) bool {
	for _,flag := range app.Flags {
		if  InFlag(flag,name) {
			return true
		}
	}
	return false
}

// InCommandFlag ：判断command(命令)中是否存在对应flag(选项),输入：父属command类型数据、command(命令)的flag(选项) Name值,输出：bool值
func InCommandFlag(command Command,name string) bool{
	for _,flag := range command.Flags {
		if  InFlag(flag,name) {
			return true
		}
	}
	return false
}

// InFlag ：判断是否存在预设flag,输入：为flag类型数据,输出：为bool值
func InFlag(flag Flag,name string) bool {
	if  strings.ToLower(name) == "--"+strings.ToLower(flag.Name) ||  name == "-"+flag.ShortName {
		return true
	}
	return false
}


/*
帮助菜单系列
*/
// AppHelp ：cli应用程序帮助菜单,输入为App类型数据
func AppHelp(app App) {
	if app.Banner != "" {
		fmt.Println(app.Banner)
	}
	if app.Name == "" {  //判断cli应用程序是否设置标题，为空时设程序名为标题
		app.Name = path.Base(os.Args[0])
	}
	fmt.Println("程序名（Name）：\n    ",app.Name)
	fmt.Println("描述（Explain）：\n    ",app.Explain)
	if app.Usage == "" {  //判断cli应用程序是否设置使用方法，为空时设默认使用方法
		app.Usage = app.Name+" [options] command [command options] "
	}
	fmt.Println("使用（Usage）：\n    ",app.Usage)
	fmt.Println("版本（Version）：\n    ",app.Version)
	if len(app.Flags) != 0 {
		fmt.Println("选项（Options）：")
		for _,flag := range app.Flags {
			FlagHelp(flag)
		}
	}
	if len(app.Commands) != 0 {
		fmt.Println("命令（Commands）：")
		for _,command := range app.Commands {
			fmt.Printf("     %-10s：%s\n",command.Name,command.Usage)
			if len(command.Flags) != 0 {
				fmt.Println("命令选项（Command Options）：")
				for _,flag := range command.Flags {
					FlagHelp(flag)
				}
			}
		}
	}
	os.Exit(0)
}

// CommandHelp ：命令帮助菜单,输入为Command类型数据
func CommandHelp(command Command) {
	if command.Name != "" {
		fmt.Println("命令（Commands）：")
		fmt.Printf("     %-10s：%s\n\n",command.Name,command.Usage)
		if len(command.Flags) != 0 {
			fmt.Println("命令选项（Command Options）：")
			for _,flag := range command.Flags {
				FlagHelp(flag)
			}
		}
		os.Exit(0)
	}
}

// FlagHelp ：参数帮助菜单,输入为flag类型数据
func FlagHelp(flag Flag)  {
	if flag.Name != "" {
		var name string
		if flag.ShortName == "" {
			name = "--"+flag.Name
		} else {
			name = "-"+flag.ShortName+","+"--"+flag.Name
		}
		fmt.Printf("     %-15s  ：%s（default:“%s”）\n",name,flag.Usage,flag.Value)
	}
}
