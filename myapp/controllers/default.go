package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"os"
	"fmt"
	"myapp/mysql"
)

type MainController struct {
	beego.Controller
}

type MysqlController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MysqlController) Post() {
    oem := os.Getenv("OEM")
	ver := os.Getenv("VER")
	c.Data["OEM"] = oem
	c.Data["VER"] = ver
	name := c.GetString("username")
	passwd := c.GetString("user_passwd")
	info := c.GetString("user_info")

	fmt.Println(name, passwd, info)
	result := mysql.Insert(mysql.DB,name,passwd,info)
	if(result == true){
		c.TplName="user/signup_success.html"
	} else {
		c.TplName="user/signup_fail.html"
	}
	
}

func (c *MysqlController) Login() {
	oem := os.Getenv("OEM")
	ver := os.Getenv("VER")
	c.Data["OEM"] = oem
	c.Data["VER"] = ver
	name := c.GetString("username")
	passwd := c.GetString("user_passwd")

	user_passwd,user_info,flag := mysql.Query(mysql.DB,name)
	if(flag == false) {
		c.TplName = "user/login_fail.html"
	}else {
		if(passwd == user_passwd){
			c.Data["UserAgent"] = c.Ctx.Request.UserAgent()
			c.Data["IP"] = c.Ctx.Input.IP()
			
			host, err := os.Hostname() 
				if err != nil { 
					fmt.Printf("%s", err) 
				} else { 
					c.Data["HostName"] = host 
				} 

			c.Data["UserName"] = name
			c.Data["Info"] = user_info
		c.TplName = "user/profile.tpl"
		}else {
			c.TplName = "user/login_fail.html"
		}
	}

}


func (main *MysqlController) ShowInfo() {
    main.Data["UserAgent"] = main.Ctx.Request.UserAgent()
	main.Data["IP"] = main.Ctx.Input.IP()
	oem := os.Getenv("OEM")
	ver := os.Getenv("VER")
	main.Data["OEM"] = oem
	main.Data["VER"] = ver
	
	host, err := os.Hostname() 
	    if err != nil { 
	        fmt.Printf("%s", err) 
	    } else { 
	        main.Data["HostName"] = host 
	    } 
	
    main.TplName = "user/profile.tpl"
}
func (main *MainController) SignUp() {
	
	oem := os.Getenv("OEM")
	ver := os.Getenv("VER")
	main.Data["OEM"] = oem
	main.Data["VER"] = ver
	main.TplName = "user/signup.html"
}

func (main *MainController) Login() {

	glog.Info( main.Ctx.Request.UserAgent())
	glog.Info( main.Ctx.Input.IP())
	oem := os.Getenv("OEM")
	ver := os.Getenv("VER")

	main.Data["OEM"] = oem
	main.Data["VER"] = ver
	main.TplName = "user/login.html"
}