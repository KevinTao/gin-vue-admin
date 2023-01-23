package feed

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
)

type FedRecord struct {
	global.GVA_MODEL
	Title         string              `json:"title" gorm:"size:128"`
	Content       string              `json:"content"`
	MemCustomerID uint                `json:"memCustomerID" gorm:"comment:关联用户ID"`
	Author        *member.MemCustomer `json:"author" gorm:"foreignkey:MemCustomerID"`
}
