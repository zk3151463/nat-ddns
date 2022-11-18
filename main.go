package main

import (
	"flag"
	"fmt"
	"nat-ddns/client"
)

// 实际中应该用更好的变量名
var (
	h bool

	v, V bool
	// t, T bool

	s string
	p string
	u string
	x int64
	y int64
)

func init() {
	flag.BoolVar(&h, "h", false, "获取帮助信息")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本信息及工具目录")
	flag.StringVar(&s, "s", "client", "要使用的功能：client：客户端、server：服务端")
	flag.StringVar(&p, "p", "./1.png", "路径")
	flag.StringVar(&u, "u", "./1p.png", "路径")
	flag.Int64Var(&x, "x", 1080, "宽")
	flag.Int64Var(&y, "y", 3000, "高")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	// flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	// flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	// flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	// flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	//flag.Usage = usage
}

func main() {
	flag.Parse()

	if len(p) >= 0 {
		if len(u) >= 0 {
			fmt.Println("正在拍照")
			client.Setjieping(p, u, x, y)
			fmt.Println("拍照结束")
		} else {
			fmt.Println("请输入截图链接")
			return
		}

	}
}
