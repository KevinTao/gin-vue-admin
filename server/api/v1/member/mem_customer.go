package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"

	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
	memberRes "github.com/flipped-aurora/gin-vue-admin/server/model/member/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type MemberApi struct {
}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
func (m *MemberApi) Register(c *gin.Context) {
	var r memberReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		e := err.(validator.ValidationErrors)[0]
		global.GVA_LOG.Error(e.Namespace())
		// fmt.Println(e.Namespace())
		// fmt.Println(e.Field())
		// fmt.Println(e.StructNamespace())
		// fmt.Println(e.StructField())
		// fmt.Println(e.Tag())
		// fmt.Println(e.ActualTag())
		// fmt.Println(e.Kind())
		// fmt.Println(e.Type())
		// fmt.Println(e.Value())
		// fmt.Println(e.Param())
		response.FailWithMessage(err.Error(), c)
		return
	}
	//err = utils.Verify(r, utils.RegisterVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	//var authorities []system.SysAuthority
	//for _, v := range r.AuthorityIds {
	//	authorities = append(authorities, system.SysAuthority{
	//		AuthorityId: v,
	//	})
	//}
	customer := &member.MemCustomer{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	customerReturn, err := memberService.Register(*customer)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(memberRes.CustomerResponse{User: customerReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(memberRes.CustomerResponse{User: customerReturn}, "注册成功", c)
}

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      memnberReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=memberRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (m *MemberApi) Login(c *gin.Context) {
	var l memberReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// err = utils.Verify(l, utils.LoginVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }

	//无需验证码
	u := &member.MemCustomer{Phone: l.Phone, Password: l.Password}
	customer, err := memberService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if customer.Enable != 1 {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	m.TokenNext(c, *customer)

	//验证码
	// if store.Verify(l.CaptchaId, l.Captcha, true) {
	// 	u := &member.MemCustomer{Phone: l.Phone, Password: l.Password}
	// 	customer, err := memberService.Login(u)
	// 	if err != nil {
	// 		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
	// 		response.FailWithMessage("用户名不存在或者密码错误", c)
	// 		return
	// 	}
	// 	if customer.Enable != 1 {
	// 		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
	// 		response.FailWithMessage("用户被禁止登录", c)
	// 		return
	// 	}
	// 	m.TokenNext(c, *customer)
	// 	return
	// }
	// response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (b *MemberApi) TokenNext(c *gin.Context, customer member.MemCustomer) {
	j := utils.NewJWT() // 唯一签名
	claims := j.CreateMemberClaims(memberReq.BaseClaims{
		UUID:     customer.UUID,
		ID:       customer.ID,
		Phone:    customer.Phone,
		Username: customer.Username,
	})
	token, err := j.CreateMemberToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(memberRes.LoginResponse{
			User:      customer,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	//多点登陆处理

	// if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
	// 	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	// 		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	// 		response.FailWithMessage("设置登录状态失败", c)
	// 		return
	// 	}
	// 	response.OkWithDetailed(systemRes.LoginResponse{
	// 		User:      user,
	// 		Token:     token,
	// 		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}, "登录成功", c)
	// } else if err != nil {
	// 	global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	// 	response.FailWithMessage("设置登录状态失败", c)
	// } else {
	// 	var blackJWT system.JwtBlacklist
	// 	blackJWT.Jwt = jwtStr
	// 	if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
	// 		response.FailWithMessage("jwt作废失败", c)
	// 		return
	// 	}
	// 	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	// 		response.FailWithMessage("设置登录状态失败", c)
	// 		return
	// 	}
	// 	response.OkWithDetailed(systemRes.LoginResponse{
	// 		User:      user,
	// 		Token:     token,
	// 		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}, "登录成功", c)
	// }
}
