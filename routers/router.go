// @APIVersion 1.0.0
// @Title beegoApi
// @Description 这是一个beegoApi ApiDemo
// @Contact 646240358@qq.com
// @TermsOfServiceUrl http://beegoApi.test
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beegoApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 访问连接地址/simple  get方式方法该路由会映射到GetFunc函数,post提交方式映射到PostFunc函数中,并且不需要@router映射
	beego.Router("/simple", &controllers.MainController{}, "post:PostFunc;get:GetFunc")
	// 路由可以多层嵌套
	// 但是自动化文档只支持二级解析,一级版本号,二级应用模块
	nsV1 := beego.NewNamespace("/v1",
		//beego.NSNamespace("/object",
		//	beego.NSInclude(
		//		&controllers.ObjectController{},
		//	),
		//),
		//beego.NSNamespace("/user",
		//	beego.NSInclude(
		//		&controllers.UserController{},
		//	),
		//),
		beego.NSNamespace("/main",
			//beego.NSNamespace("/echo",
			beego.NSInclude(
				&controllers.MainController{},
			),
			//),
		),
	)
	//nsv2 := beego.NewNamespace("/v2",
	//	beego.NSNamespace("/main",
	//		beego.NSInclude(
	//			&controllers.MainController{},
	//		),
	//	),
	//)
	beego.AddNamespace(nsV1 /*, nsv2*/)
}
