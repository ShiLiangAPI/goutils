package base_model

import (
	"github.com/ShiLiangAPI/goutils/snowflake"
	"github.com/ShiLiangAPI/goutils/types"
	"gorm.io/gorm"
	"strconv"
)

type Model struct {
	ID        string           `gorm:"<-:create;type:bigint;primaryKey" json:"id"` // id
	CreatedAt *types.LocalTime `gorm:"comment:'创建时间'" json:"created_at"`           // 创建时间
	UpdatedAt *types.LocalTime `gorm:"comment:'修改时间'" json:"updated_at"`           // 修改时间
}

func (obj *Model) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = strconv.FormatInt(snowflake.GetFlake().GetNextID(), 10)

	return
}
