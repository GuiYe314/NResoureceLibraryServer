package net

import (
	"NResoureceLibraryServer/net/netModel"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _netLint *gin.Engine
var _dbLink *gorm.DB

var _user netModel.Users

func Init_Login(dbLink *gorm.DB, netLink *gin.Engine) {
	_netLint = netLink
	_dbLink = dbLink
	register()
	Login()
}

// 注册信息
func register() {
	_netLint.POST("/register", func(c *gin.Context) {
		//创建注册模型
		var registerModel netModel.RegisterModel
		//获取注册信息
		c.BindJSON(&registerModel)
		if len(registerModel.UserName) == 0 {
			SendMassage(c, 410, "用户名不能为空")
			return
		}
		if len(registerModel.Password) == 0 {
			SendMassage(c, 410, "密码不能为空")
			return
		}
		if registerModel.Password != registerModel.ConfirmPassword {
			SendMassage(c, 410, "密码和确认密码不同")
			return
		}

		//查询数据库是否存在用户
		var count int64
		_dbLink.Model(&netModel.Users{}).Where("User_Name = ?", registerModel.UserName).Count(&count)

		if count > 0 {
			SendMassage(c, 410, "用户名已经注册")
			return
		}

		//写入数据库
		_user.User_Name = registerModel.UserName
		_user.Password = registerModel.Password
		_user.Created_At = time.Now().UTC()
		_user.Updated_At = time.Now().UTC()
		_dbLink.AutoMigrate(&netModel.Users{})
		result := _dbLink.Create(&_user)
		if result.Error != nil {
			SendMassage(c, 410, "注册失败:::"+result.Error.Error())
			return
		} else {
			SendMassage(c, 200, "注册成功")
		}

	})
}

func Login() {
	_netLint.POST("/login", func(c *gin.Context) {
		// 创建登陆模块
		var loginModel netModel.LoginModel
		// 获取注册信息
		c.BindJSON(&loginModel)
		if len(loginModel.UserName) == 0 {
			SendMassage(c, 420, "用户名不能为空")
			return
		}
		if len(loginModel.Password) == 0 {
			SendMassage(c, 420, "密码不能为空")
			return
		}
		//查询有没有用户
		var userFind netModel.Users
		_dbLink.Model(&netModel.Users{}).Where("user_name = ?", loginModel.UserName).First(&userFind)
		if len(userFind.User_Name) == 0 {
			SendMassage(c, 421, "未查找到用户:"+loginModel.UserName)
			return
		}
		//判断密码
		if userFind.Password != loginModel.Password {
			SendMassage(c, 422, "密码错误:"+loginModel.UserName)
			return
		}

		GenerateToken(c, int64(userFind.ID), userFind.User_Name)
	})
}
