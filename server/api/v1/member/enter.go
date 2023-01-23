package member

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	MemberApi
	BaseApi
}

var (
	memberService = service.ServiceGroupApp.MemberServiceGroup.CustomerService
)
