package routers

import (
	"Hello/admin"
	"Hello/controllers"
	"Hello/controllers/admin2"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/admin",&admin.Usercontroller{})
	beego.Router("/admin2",&admin2.Usercontroller2{})
	beego.Router("/regular/?:id",&controllers.RegularController{})
	beego.Router("/file",&controllers.FormController{})
	beego.Router("/test_input", &controllers.SessionController{})
	beego.Router("/test_login", &controllers.Session2Controller{})
    beego.Router("/login",&controllers.L{})
}
