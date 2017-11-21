package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/profile", &controllers.MysqlController{},"get:ShowInfo;post:Login")
	beego.Router("/user/signup", &controllers.MainController{},"get:SignUp")
	beego.Router("/user/signup_result",&controllers.MysqlController{})
	beego.Router("/user/login", &controllers.MainController{},"get:Login")
}
