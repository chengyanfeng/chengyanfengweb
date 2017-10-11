package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	."chengyanfengweb/models"
)
type ChengController struct {
	beego.Controller
}

func (this *ChengController)Getstudent(){
	student:=Student{}
	fmt.Print("testok")
	DB.Find(&student)
	fmt.Println(student)
	this.TplName="index.html"
}
func (this *ChengController) Get(){
	this.Data["32423"]="women"
}