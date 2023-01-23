package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
)

type CustomerResponse struct {
	User member.MemCustomer `json:"user"`
}

type LoginResponse struct {
	User      member.MemCustomer `json:"user"`
	Token     string             `json:"token"`
	ExpiresAt int64              `json:"expiresAt"`
}
