package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
// 	this.Data["Website"] = "beego.me"
// 	this.Data["Email"] = "astaxie@gmail.com"
// 	this.Data["Fqworld"] = "fqworld~~"
    this.Data["幹"]="早班小7有妹"
	this.TplNames = "jndex.tpl"
}
