package member

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

func (c *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	// customerRouter := Router.Group("customer").Use(middleware.ClientOperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	memberApi := v1.ApiGroupApp.CustomerApiGroup.MemberApi
	{
		customerRouterWithoutRecord.POST("register", memberApi.Register) // 用户注册账号
		customerRouterWithoutRecord.POST("login", memberApi.Login)       //用户登录账号
		//customerRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		//customerRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		//customerRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		//customerRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		//customerRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		//customerRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		//customerRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	{
		//customerRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		//customerRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
	}
}
