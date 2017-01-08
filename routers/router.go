package routers

import (
	"github.com/astaxie/beego"
	"merit/controllers"
)

func init() {
	//1.首页index
	beego.Router("/", &controllers.MainController{})
	//管理员
	beego.Router("/admin", &controllers.AdminController{}) //admin-get()
	//显示对应侧栏id的右侧界面
	beego.Router("/admin/:id:string", &controllers.AdminController{}, "*:Admin")
	//填充部门表格数据
	beego.Router("/admin/department", &controllers.AdminController{}, "*:Department")
	//根据数字id查询类别或目录分级表
	beego.Router("/admin/department/?:id:string", &controllers.AdminController{}, "*:Department")
	//根据名字查询目录分级表
	beego.Router("/admin/departmenttitle", &controllers.AdminController{}, "*:DepartmentTitle")
	//添加目录类别
	beego.Router("/admin/department/adddepartment", &controllers.AdminController{}, "*:AddDepartment")
	//修改目录类别
	beego.Router("/admin/department/updatedepartment", &controllers.AdminController{}, "*:UpdateDepartment")
	//删除目录类
	beego.Router("/admin/department/deletedepartment", &controllers.AdminController{}, "*:DeleteDepartment")
	//*******价值****
	//填充MERIT表格数据
	beego.Router("/admin/merit", &controllers.AdminController{}, "*:Merit")
	//根据数字id查询
	beego.Router("/admin/merit/?:id:string", &controllers.AdminController{}, "*:Merit")
	//添加
	beego.Router("/admin/merit/addmerit", &controllers.AdminController{}, "*:AddMerit")
	//修改
	beego.Router("/admin/merit/updatemerit", &controllers.AdminController{}, "*:UpdateMerit")
	//删除
	beego.Router("/admin/merit/deletemerit", &controllers.AdminController{}, "*:DeleteMerit")

	//查询所有ip地址段
	beego.Router("/admin/ipsegment", &controllers.AdminController{}, "*:Ipsegment")
	//添加IP地址段
	beego.Router("/admin/ipsegment/addipsegment", &controllers.AdminController{}, "*:AddIpsegment")
	//修改IP地址段
	beego.Router("/admin/ipsegment/updateipsegment", &controllers.AdminController{}, "*:UpdateIpsegment")
	//删除IP地址段
	beego.Router("/admin/ipsegment/deleteipsegment", &controllers.AdminController{}, "*:DeleteIpsegment")
	//编辑成果类型和折标系数表
	beego.Router("/admin/achievcategory", &controllers.Achievement{}, "get:Achievcategory")
	beego.Router("/admin/achievcategory/addachievcategory", &controllers.Achievement{}, "post:AddAchievcategory")
	beego.Router("/admin/achievcategory/updateachievcategory", &controllers.Achievement{}, "post:UpdateAchievcategory")
	beego.Router("/admin/achievcategory/deleteachievcategory", &controllers.Achievement{}, "post:DeleteAchievcategory")

	// beego.Router("/jsoneditor", &controllers.AdminController{}, "get:Jsoneditor")

	//2.2首页进入成果登记
	beego.Router("/achievement", &controllers.Achievement{}, "get:GetAchievement")
	//这个同上面一样
	beego.Router("/getachievement", &controllers.Achievement{}, "get:GetAchievement")
	//个人在线添加成果
	beego.Router("/achievement/addcatalog", &controllers.Achievement{}, "post:AddCatalog")
	//个人在线直接提交成果
	beego.Router("/achievement/sendcatalog", &controllers.Achievement{}, "post:SendCatalog")
	//在线退回成果
	beego.Router("/achievement/downsendcatalog", &controllers.Achievement{}, "post:DownSendCatalog")
	//个人在线修改保存成果
	beego.Router("/achievement/modifycatalog", &controllers.Achievement{}, "post:ModifyCatalog")
	//个人在线删除一条成功
	beego.Router("/achievement/delete", &controllers.Achievement{}, "post:DeleteCatalog")
	//测试某个专业下总成本分布情况
	beego.Router("/achievement/specialty", &controllers.Achievement{}, "get:Specialty")
	//个人当月成果类型组成
	beego.Router("/achievement/echarts", &controllers.Achievement{}, "get:Echarts")
	//beego.Router("/achievement/echarts1", &controllers.Achievement{}, "get:Echarts1")
	//个人当年成果类型组成
	beego.Router("/achievement/echarts2", &controllers.Achievement{}, "get:Echarts2")
	//项目阶段专业，全年成果类型组成
	beego.Router("/achievement/echarts3", &controllers.Achievement{}, "get:Echarts3")

	//点击个人参与的项目，弹出模态框，显示这个项目所有成果
	beego.Router("/achievement/projectachievement", &controllers.Achievement{}, "*:ProjectAchievement")
	//一年来一个项目的每个人的贡献率
	beego.Router("/achievement/projectuserparticipate", &controllers.Achievement{}, "*:ProjectUserParticipate")

	beego.Router("/test", &controllers.AdminController{}, "get:Test")
	beego.Router("/test1", &controllers.AdminController{}, "get:Test1")
	//这个是ue编辑器用的
	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")
	//2.1首页进入价值——根据登陆者的权限，显示对应的侧栏：普通用户直接进入自己的价值侧栏
	beego.Router("/merit", &controllers.MeritController{}, "get:GetMerit")
	// 主页里显示iframe——科室总体情况
	beego.Router("/merit/secofficeshow", &controllers.MeritController{}, "get:Secofficeshow")
	// 填充科室总体情况数据
	beego.Router("/merit/secofficedata", &controllers.MeritController{}, "get:SecofficeData")

	//2.1.1普通用户进入价值，侧栏右边的iframe
	// beego.Router("/merit/myself", &controllers.JsonController{}, "get:Myself")
	//显示所有人价值排名
	// beego.Router("/getperson", &controllers.JsonController{}, "get:GetPerson")
	//进行价值结构编辑
	// beego.Router("/get", &controllers.JsonController{})
	// beego.Router("/json", &controllers.JsonController{}) //这个和上面等价
	// beego.Router("/importjson", &controllers.JsonController{}, "post:ImportJson")

	// beego.Router("/addjson", &controllers.JsonController{}, "post:Addjson")
	// beego.Router("/modifyjson", &controllers.JsonController{}, "get:Modifyjson")          //显示修改页面
	// beego.Router("/modifyjsonpost", &controllers.JsonController{}, "post:ModifyjsonPost") //提交修改
	// beego.Router("/deletejson", &controllers.JsonController{}, "get:Deletejson")

	// beego.Router("/merit/add", &controllers.MeritTopicController{}, "get:MeritAdd")
	// beego.Router("/AddMeritTopic", &controllers.MeritTopicController{}, "post:AddMeritTopic")
	// beego.Router("/view", &controllers.MeritTopicController{}, "get:ViewMeritTopic")
	// beego.Router("/modify", &controllers.MeritTopicController{}, "get:ModifyMeritTopic")
	// beego.Router("/ModifyPost", &controllers.MeritTopicController{}, "post:ModifyPost")
	// beego.Router("/delete", &controllers.MeritTopicController{}, "get:DeleteMeritTopic")

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")

	beego.Router("/regist", &controllers.RegistController{})
	// beego.Router("/registerr", &controllers.RegistController{}, "get:RegistErr")
	beego.Router("/regist/checkuname", &controllers.RegistController{}, "post:CheckUname")
	beego.Router("/regist/getuname", &controllers.RegistController{}, "post:GetUname")
	//get方法用于x-editable的select2方法_作废，select2不必须要动态数据
	beego.Router("/regist/getuname1", &controllers.RegistController{}, "get:GetUname1")

	//成果登记系统
	//成果登记表导入数据库
	beego.Router("/achievement/import_xls_catalog", &controllers.Achievement{}, "post:Import_Xls_Catalog")
	// 主页里显示iframe——科室总体情况
	beego.Router("/achievement/secofficeshow", &controllers.Achievement{}, "get:Secofficeshow")
	// 填充科室总体情况数据
	beego.Router("/achievement/secofficedata", &controllers.Achievement{}, "get:SecofficeData")

	//用户在线登记时，自己发起的成果，还未提交
	beego.Router("/achievement/myself", &controllers.Achievement{}, "get:Myself")
	//自己发起的成果，已经提交
	beego.Router("/achievement/running", &controllers.Achievement{}, "get:Running")
	//别人传来，自己处于设计位置
	beego.Router("/achievement/designd", &controllers.Achievement{}, "get:Designd")
	//别人传来，自己处于设计位置
	beego.Router("/achievement/checked", &controllers.Achievement{}, "get:Checked")
	//别人传来，自己处于设计位置
	beego.Router("/achievement/examined", &controllers.Achievement{}, "get:Examined")
	//查看用户个人时，获取已经完成的数据
	beego.Router("/achievement/completed", &controllers.Achievement{}, "get:Completed")
	//获取自己参与的项目列表
	beego.Router("/achievement/participate", &controllers.Achievement{}, "get:Participate")
	//获取科室的项目列表
	beego.Router("/achievement/secparticipate", &controllers.Achievement{}, "get:SecParticipate")
	//获取科室的当月成果列表
	beego.Router("/achievement/secprojectachievement", &controllers.Achievement{}, "get:SecProjectAchievement")

	//人员管理
	beego.Router("/user/AddUser", &controllers.UserController{}, "*:AddUser")
	beego.Router("/user/UpdateUser", &controllers.UserController{}, "*:UpdateUser")
	beego.Router("/user/deluser", &controllers.UserController{}, "*:DelUser")
	beego.Router("/user/index", &controllers.UserController{}, "*:Index")
	//管理员修改用户资料
	beego.Router("/user/view", &controllers.UserController{}, "get:View")
	beego.Router("/user/view/*", &controllers.UserController{}, "get:View")
	beego.Router("/user/importexcel", &controllers.UserController{}, "post:ImportExcel")

	//用户修改自己密码
	beego.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	//用户登录后查看自己的资料
	beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")

	beego.Router("/role/AddAndEdit", &controllers.RoleController{}, "*:AddAndEdit")
	beego.Router("/role/DelRole", &controllers.RoleController{}, "*:DelRole")
	// beego.Router("/role/AccessToNode", &controllers.RoleController{}, "*:AccessToNode")
	// beego.Router("/role/AddAccess", &controllers.RoleController{}, "*:AddAccess")
	beego.Router("/role/RoleToUserList", &controllers.RoleController{}, "*:RoleToUserList")
	beego.Router("/role/AddRoleToUser", &controllers.RoleController{}, "*:AddRoleToUser")
	beego.Router("/role/Getlist", &controllers.RoleController{}, "*:Getlist")
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/roleerr", &controllers.RoleController{}, "*:Roleerr") //显示权限不够

}
