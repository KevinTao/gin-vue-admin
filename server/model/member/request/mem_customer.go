package request

// Register User register structure
type Register struct {
	Username  string `json:"userName" example:"用户名" binding:"required_without=Phone,gte=4,lte=32"`
	Password  string `json:"passWord" example:"密码" binding:"required_if=Username,gte=8"`
	NickName  string `json:"nickName" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Enable    int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	Phone     string `json:"phone" example:"电话号码" binding:"required_without=Username"`
	Email     string `json:"email" example:"电子邮箱"`
}

// User login structure
type Login struct {
	Username  string `json:"userName" example:"用户名" binding:"required_without=Phone"`
	Phone     string `json:"phone" binding:"required_without=Username"` // 用手机号
	Password  string `json:"password" binding:"required_if=Username"`   // 密码
	SmsCode   string `json:"smsCode" bindging:"required_if=Phone"`
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID        uint   `gorm:"primarykey"`                                                                           // 主键ID
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone     string `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email     string `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Enable    int    `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
}
