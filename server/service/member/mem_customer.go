package member

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CustomerService struct {
}

//@author: [calvin]
//@function: Register
//@description: 会员注册
//@param: c member.MemCustomer
//@return: customer member.MemCustomer, err error

func (customerService *CustomerService) Register(c member.MemCustomer) (customer member.MemCustomer, err error) {
	if !errors.Is(global.GVA_DB.Where("phone = ?", c.Phone).First(&customer).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return customer, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	c.Password = utils.BcryptHash(c.Password)
	c.UUID = uuid.NewV4()
	if len(c.Username) == 0 {
		c.Username = c.Phone
	}
	err = global.GVA_DB.Create(&c).Error
	return c, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (customerService *CustomerService) Login(c *member.MemCustomer) (customer *member.MemCustomer, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user member.MemCustomer
	if len(c.Phone) > 0 {
		err = global.GVA_DB.Where("phone = ?", c.Phone).First(&user).Error
		if err == nil {
			//短信验证码验证
			// if ok := utils.BcryptCheck(c.Password, user.Password); !ok {
			// 	return nil, errors.New("密码错误")
			// }
			// MenuServiceApp.UserAuthorityDefaultRouter(&user)
		}
	} else if len(c.Username) > 0 {
		err = global.GVA_DB.Where("username = ?", c.Username).First(&user).Error
		if err == nil {
			if ok := utils.BcryptCheck(c.Password, user.Password); !ok {
				return nil, errors.New("密码错误")
			}
		}
	} else {
		return nil, errors.New("缺少用户名或手机号码")
	}
	return &user, err
}
