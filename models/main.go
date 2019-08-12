package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct
type User struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:"size(100)"`
}

// 初始化数据库连接
func init() {
	// set default database
	//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
	//	beego.AppConfig.String("mysqluser"),
	//	beego.AppConfig.String("mysqlpasswd"),
	//	beego.AppConfig.String("mysqladdr"),
	//	beego.AppConfig.String("mysqlport"),
	//	beego.AppConfig.String("mysqldatabase"),
	//)
	//// 注册数据库驱动,默认已注册
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//
	//orm.RegisterDataBase(
	//	"default",
	//	beego.AppConfig.String("sql_driver"),
	//	dataSource,
	//	30)
	//
	//// register model
	//// orm.RegisterModelWithPrefix("prefix_", new(User)) // 使用表前缀
	//orm.RegisterModel(new(User),new())
	//// create table
	//orm.RunSyncdb("default", false, true)
}

// 增
func Add(userName string) (err error) {
	var (
		o    orm.Ormer
		id   int64
		user *User
	)
	o = orm.NewOrm()
	user = new(User)
	user.Name = userName
	if id, err = o.Insert(user); err != nil {
		logs.Info("插入失败.")
	} else {
		logs.Info("插入数据成功:", id)
	}
	return
}

// 删
func Del(userId int) (err error) {
	var (
		o    orm.Ormer
		num  int64
		user *User
	)
	o = orm.NewOrm()
	user = new(User)
	user.Id = userId
	if num, err = o.Delete(user); err != nil {
		logs.Info("删除失败", err)
	} else {
		logs.Info("删除成功", num)
	}
	return
}

// 改
func Modify(userId int, userName string) (user *User, err error) {
	var (
		o   orm.Ormer
		num int64
	)
	o = orm.NewOrm()
	user = new(User)
	user.Id = userId
	if err = o.Read(user); err != nil {
		logs.Info("读取失败")
	}
	user.Name = userName
	if num, err = o.Update(user); err != nil {
		logs.Info("更新失败", err)
	} else {
		logs.Info("更新成功", num)
	}
	return
}

// 查
func Get(userId int) (user *User, err error) {
	var (
		o orm.Ormer
	)
	o = orm.NewOrm()
	user = new(User)
	user.Id = userId
	if err = o.Read(user); err != nil {
		logs.Info("查询失败", err)
	} else {
		logs.Info("查询成功")
	}
	return
}
