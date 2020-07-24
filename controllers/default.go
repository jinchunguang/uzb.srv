package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	fmt.Println("ospid:" + strconv.Itoa(os.Getpid()))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
