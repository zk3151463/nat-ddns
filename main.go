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
)

func init() {
	flag.BoolVar(&h, "h", false, "获取帮助信息")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本信息及工具目录")
	flag.StringVar(&s, "s", "client", "要使用的功能：client：客户端、server：服务端")
	flag.StringVar(&p, "p", "9000", "服务的端口")

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
		fmt.Println(p)
		client.Setjieping(p)
	}
}
