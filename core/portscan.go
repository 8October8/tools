package core

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// PortScanCode ：多线程端口开放扫描器,核心功能函数,输入：IP列表、端口列表、线程数、协议、延时数,输出：开放IP端口列表、关闭IP端口列表
func PortScanCode(hosts []string,ports []string,threads int,protocol string,timeout int)(proton []string,portoff []string){

	thread := make(chan int, threads)  	//定义缓存通道,设置线程数
	var wg sync.WaitGroup				// 声明一个等待组

	for _,host := range hosts {
		for _,port := range ports {
			address := host+":"+port
			fmt.Printf("当前扫描目标：%-21s  \r",address)
			wg.Add(1)		//等待组加1
			thread <- 1			//缓存通道写入加一,若缓存已满,则堵塞程序,等待缓存空位
			go func(address string) {
				conn, err := net.DialTimeout(protocol, address, time.Millisecond*time.Duration(timeout))
				if err == nil {
					proton = append(proton, address)
					_ = conn.Close()
				} else {
					portoff = append(portoff, address)
				}
				<- thread		//缓存通道读取加一,释放一位缓存
				defer wg.Done()		//等待组减1
			}(address)
		}
	}
	wg.Wait()		//等待等待组置零,协程运行完毕
	return
}