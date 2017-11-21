package main

import (
	 
	 "myapp/routers"
	 "flag"
	 "github.com/astaxie/beego"
	 "github.com/golang/glog"
	 "myapp/echo/echo_server"
	 "myapp/mysql"
)



func main() {
	
	flag.Parse()
	mysql.Init()
	routers.Init()
	go echo_server.Start()
	beego.Run()
    glog.Flush()
}

