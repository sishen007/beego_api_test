package controllers

import (
	"beegoApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

// Main Api Controller 这是模块注释
type MainController struct {
	beego.Controller
}

// 应用注释
// @Title API所表达的含义
// @Description API详细的描述
// @param 参数名 参数类型1(formData 表示是 post 请求的数据; query 表示带在 url 之后的参数; path 表示请求路径上得参数; body 表示是一个 raw 数据请求; header 表示带在 header 信息中得参数) 参数字段类型 是否必须 注释
// @Param name query string false nameFiled
// @Success 200 {string} 成功返回给客户端的信息(三个参数:status_code {object/string/int等} 返回的对象或者字符串信息)
// @Failure 403 失败返回的信息(两个参数:status_code 错误信息)
// @Failure 400 no enough input
// @router /hello [get]
func (m *MainController) Hello() {
	data := make(map[string]string)
	data["hello"] = "Hello World! this is my one test api controlle"
	data["appname"] = beego.BConfig.AppName
	data["appname2"] = beego.AppConfig.String("appname")
	data["httpport"] = beego.AppConfig.String("httpport")
	data["dev.httpport"] = beego.AppConfig.String("dev::httpport")
	data["prod.httpport"] = beego.AppConfig.String("prod::httpport")
	data["GOPATH"] = beego.AppConfig.String("gopathtest")
	data["GOPATHNO"] = beego.AppConfig.String("gopathno")
	m.Data["json"] = data
	m.ServeJSON()
	m.StopRun()
}

// @Title 获取用户
// @Description 获取用户
// @Param id query int true Id
// @Success 200 {object} models.User
// @Failure 400 no enough input
// @router / [get]
func (this *MainController) Get() {
	var (
		err   error
		idStr string
		id    int
		user  *models.User
	)
	// 获取参数
	idStr = this.Input().Get("id")
	id, err = strconv.Atoi(idStr)
	if user, err = models.Get(id); err != nil {
		logs.Info("获取失败", err)
	} else {
		logs.Info("获取成功", user)
	}
	this.Data["json"] = user
}

// @Title 新增用户
// @Description 新增用户
// @Param name formData string false nameFiled
// @router / [post]
func (this *MainController) Post() {
	var (
		err  error
		name string
	)
	// 获取参数+附带默认值
	name = this.GetString("name", "sishen007")
	if err = models.Add(name); err != nil {
		logs.Info("插入失败", err)
	}
	this.Data["json"] = "插入成功"
}

// @Title 更新用户
// @Description 更新用户
// @Param id formData int false IdFiled
// @Param name formData string false nameFiled
// @Success 200 {object} models.User
// @Failure 400 no enough input
// @router / [put]
func (this *MainController) Put() {
	var (
		err   error
		id    int
		name  string
		idStr string
		user  *models.User
	)
	// 获取参数+附带默认值
	// 获取参数
	idStr = this.Input().Get("id")
	name = this.GetString("name", "sishen007")
	id, err = strconv.Atoi(idStr)
	if user, err = models.Modify(id, name); err != nil {
		logs.Info("更新失败", err)
	}
	this.Data["json"] = user
}

// @Title API所表达的含义
// @Description API详细的描述
// @Param id formData int false IdFiled
// @Success 200 {string} 删除成功
// @Failure 400 no enough input
// @router / [delete]
func (this *MainController) Delete() {
	var (
		err   error
		id    int
		idStr string
	)
	// 获取参数+附带默认值
	// 获取参数
	idStr = this.Input().Get("id")
	id, err = strconv.Atoi(idStr)
	if err = models.Del(id); err != nil {
		logs.Info("更新失败", err)
	}
	this.Data["json"] = "删除成功"
}

func (m *MainController) GetFunc() {
	m.Data["json"] = "Get func .... "
	//m.Finish()
	m.ServeJSON()
	m.StopRun()
}
func (m *MainController) Finish() {
	//m.Data["json"] = "Finish..."
	m.ServeJSON()
}

func (m *MainController) PostFunc() {
	m.Data["json"] = "post func .... "
	m.ServeJSON()
}
