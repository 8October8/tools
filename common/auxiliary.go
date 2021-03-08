//部分辅助功能函数
package common

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// InList ：判断字符串切片是否包含某个值,第三个参数判断是否进行大小写敏感
func InList(value string, list []string, onoff string) bool {
	onoff = strings.ToLower(onoff)
	if onoff == "on" {
		value = strings.ToLower(value)
		for _, v := range list {
			v = strings.ToLower(v)
			if v == value {
				return true
			}
		}
	} else {
		for _, v := range list {
			if v == value {
				return true
			}
		}
	}
	return false
}

// IPCheck ：对IP正确性的校验
func IPCheck(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

//IPlist ：对常见ip参数的列表分解，检测的IP参数：（192.168.1.1）（192.168.1.1，192.168.1.2）（192.168.1.1/24）（192.168.1.1-192.168.1.254）（192.168.1.1-254）
func IPlist(ips string) []string {
	var list []string  //初始化列表切片
	switch {

	//通过搜索是否存在“,”符号，判断是否为分隔符“,”的列表字符串，是则分割为切片数组
	case strings.ContainsAny(ips, ","):
		list = strings.Split(ips, ",")

	//通过搜索是否存在“/”符号，判断是否为“/”划分的子网，是则通过net进行切片数组转换
	case strings.ContainsAny(ips, "/"):
		ip, ipNet, err := net.ParseCIDR(ips)
		if err != nil {
			fmt.Println("IP参数存在错误：", ips)
			os.Exit(0)
		}
		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); {
			list = append(list, ip.String())
			for j := len(ip) - 1; j >= 0; j-- {
				ip[j]++
				if ip[j] > 0 {
					break
				}
			}
		}

	//通过搜索是否存在“-”符号，判断是否为分隔符“-”的连续列表字符串，是则转换为切片数组
	case strings.ContainsAny(ips, "-"):
		s := strings.LastIndex(ips, "-") //定位“-”符号位置

		if !IPCheck(ips[:s]) { //对IP参数进行校验
			fmt.Println("IP参数存在错误：", ips)
			os.Exit(0)
		}

		var LeftIP, RightIP []string  //定义区部变量
		var L, R int

		LeftIP = strings.Split(ips[:s], ".")  //分解左边IP参数，提取最后一位数据
		L, err := strconv.Atoi(LeftIP[3])

		//判断右方数据为IP还是数字，为IP时分解IP参数，提取最后一位数据
		if IPCheck(ips[s+1:]) {
			//参数校验
			RightIP = strings.Split(ips[s+1:], ".")
			if  LeftIP[0] != RightIP[0] || LeftIP[1] != RightIP[1] || LeftIP[2] != RightIP[2] {  //左右IP参数前三位对比，防止出现跨c段参数
				fmt.Println("IP参数存在错误：", ips)
				os.Exit(0)
			}
			R, err = strconv.Atoi(RightIP[3])
		} else {
			R, err = strconv.Atoi(ips[s+1:])
		}

		if R > 255 || err != nil {
			fmt.Println("IP参数存在错误：", ips)
			os.Exit(0)
		}

		//判断大小，如果为左边大则交换顺序
		if L > R {
			L, R = R, L
		}

		//拼接完整IP并追究至切片
		for ; L <= R; L++ {
			var ip strings.Builder
			ip.WriteString(LeftIP[0])
			ip.WriteString(".")
			ip.WriteString(LeftIP[1])
			ip.WriteString(".")
			ip.WriteString(LeftIP[2])
			ip.WriteString(".")
			ip.WriteString(strconv.Itoa(L))
			list = append(list, ip.String())
		}

	//无上述分隔符时，默认为单IP参数
	default:
		list = append(list, ips)
	}

	//对列表中IP进行校验
	for _, ip := range list {
		if !IPCheck(ip) {
			fmt.Println("IP列表存在错误数据：", ip)
			os.Exit(0)
		}
	}
	return list
}

//PortList ：对常见端口参数的列分解，检测的口参数：（80）（80，81）（1-65535）(-65535)
func PortList(posts string) []string {
	var list []string //定义返回列表
	switch {

	//通过搜索是否存在“,符号，判断是否为分隔符“,”的列表字符串是则分割为切片数组
	case strings.ContainsAny(posts, ","):
		list = strings.Split(posts, ",")

	//通过搜索是否存在“-”符号，判断是否为分隔符“-”的连续列表字符串，是则转换为切片数组
	case strings.ContainsAny(posts, "-"):

		s := strings.Index(posts, "-")    //位“-”符号位置
		L, err := strconv.Atoi(posts[:s]) //切割左右
		R, err := strconv.Atoi(posts[s+1:])
		if err != nil {
			fmt.Println("ports参数存在错误：", err)
			os.Exit(0)
		}
		if L > R { //判断大小，如果为左边大则交换顺序
			L, R = R, L
		}
		for ; L <= R; L++ { //追加成为切片数据
			list = append(list, strconv.Itoa(L))
		}

	//无上述分隔符时，默认为单post参数
	default:
		list = append(list, posts)
	}

	//列表中port进行校验
	for _, port := range list {
		_, err := strconv.ParseFloat(port, 64)
		L, err := strconv.Atoi(port)
		if err != nil || L < 0 || L > 65535 {
			fmt.Println("post列表存在错误数据：", port)
			os.Exit(0)
		}
	}
	return list
}
