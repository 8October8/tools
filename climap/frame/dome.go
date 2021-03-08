package frame

import (
	"fmt"
)

// CliMenuDemo :对climenu中App的dome
func CliMenuDemo()  {
	//对cli应用程序的信息初始设置
	var app App
	app.Name = "程序的名称,默认为path.Base(os.Args[0])"
	app.Explain = "程序的描述"
	app.Usage = "程序的使用描述"
	app.Version = "程序的版本"
	app.Commands = []Command{
		{
			Name: "使用命令",
			Usage: "命令的使用描述",
			Version: "命令的版本",
			Flags: []Flag{
				{
					Name: "标志的名称",
					ShortName: "标志的短名称。通常一个字符",
					Value: "标志的短名称。通常一个字符",
					Usage: "标志的使用描述",
				},
			},
		},
		{
			Name: "使用命令2",
			Usage: "命令的使用描述2",
			Version: "命令的版本2",
			Flags: []Flag{
				{
					Name: "标志的名称2",
					ShortName: "标志的短名称。通常一个字符2",
					Value: "标志的短名称。通常一个字符2",
					Usage: "标志的使用描述2",
				},
			},
		},
	}
	app.Flags = []Flag{
		{
			Name: "标志的名称",
			ShortName: "标志的短名称。通常一个字符",
			Value: "标志的短名称。通常一个字符",
			Usage: "标志的使用描述",
		},
		{
			Name: "标志的名称2",
			ShortName: "标志的短名称。通常一个字符2",
			Value: "标志的短名称。通常一个字符2",
			Usage: "标志的使用描述2",
		},
	}

	//相关数据的调用输出
	fmt.Println("app.Name:\n",app.Name)
	fmt.Println("app.Explain:\n",app.Explain)
	fmt.Println("app.Usage:\n",app.Usage)
	fmt.Println("app.Version:\n",app.Version)
	fmt.Println("app.Commands:\n",app.Commands)
	fmt.Println("app.Commands[0]:\n",app.Commands[0])
	Acommand := app.Commands[0]
	fmt.Println("Acommand.Name:\n", Acommand.Name)
	fmt.Println("Acommand.Usage:\n", Acommand.Usage)
	fmt.Println("Acommand.Version:\n", Acommand.Version)
	fmt.Println("Acommand.Flags:\n", Acommand.Flags)
	fmt.Println("Acommand.Name:\n", Acommand.Flags[0])
	Cflag := Acommand.Flags[0]
	fmt.Println("Cflag.Name:\n", Cflag.Name)
	fmt.Println("Cflag.ShortName:\n", Cflag.ShortName)
	fmt.Println("Cflag.Value:\n", Cflag.Value)
	fmt.Println("Cflag.Usage:\n", Cflag.Usage)

	fmt.Println("app.Flags:\n",app.Flags)
	fmt.Println("app.Flags[0]:\n",app.Flags[0])
	flag := app.Flags[0]
	fmt.Println("app.Flags:\n",flag.Name)
	fmt.Println("flag.ShortName:\n", flag.ShortName)
	fmt.Println("flag.Value:\n", flag.Value)
	fmt.Println("flag.Usage:\n", flag.Usage)

	//帮助菜单
	AppHelp(app)
	CommandHelp(app.Commands[0])
	FlagHelp(app.Flags[0])

	//输出其它非预设参数和选项
	fmt.Println(NArg(app))
}