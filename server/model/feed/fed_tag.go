package feed

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type FedTag struct {
	global.GVA_MODEL
	Name        string `json:"name" gorm:"size:64"`
	Description string `json:"desc" gorm:"size:256;default:-"`
}
