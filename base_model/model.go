package base_model

import (
	"gitee.com/ShiLiangAPI/goutils/snowflake"
	"gitee.com/ShiLiangAPI/goutils/types"
	"gorm.io/gorm"
	"strconv"
)

type Model struct {
	ID        string           `gorm:"<-:create;type:bigint;primaryKey" json:"id"`
	CreatedAt *types.LocalTime `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt *types.LocalTime `gorm:"comment:'修改时间'" json:"updatedAt"`
}

func (obj *Model) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = strconv.FormatInt(snowflake.GetFlake().GetNextID(), 10)

	return
}
