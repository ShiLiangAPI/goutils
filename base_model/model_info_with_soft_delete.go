package base_model

import (
	"github.com/ShiLiangAPI/goutils/snowflake"
	"github.com/ShiLiangAPI/goutils/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"strconv"
)

type BaseModelInfo struct {
	ID        string                `gorm:"<-:create;type:bigint;primaryKey" json:"id"`                            // id
	CreatedAt *types.LocalTime      `gorm:"comment:'创建时间'" json:"created_at"`                                      // 创建时间
	UpdatedAt *types.LocalTime      `gorm:"comment:'修改时间'" json:"updated_at"`                                      // 修改时间
	DeletedAt soft_delete.DeletedAt `gorm:"<-:create;softDelete:milli;index;uniqueIndex:udx_1" json:"-"`           // 删除时间
	Name      string                `gorm:"comment:'名称';not null" json:"name"`                                     // 名称
	Code      string                `gorm:"comment:'编码';type:varchar(255);not null;uniqueIndex:udx_1" json:"code"` // 编码
	IsEnabled bool                  `gorm:"comment:'是否启用';default:true" json:"is_enabled"`                         // 是否启用
	Sort      int                   `gorm:"comment:'排序'" json:"sort"`                                              // 排序
}

func (obj *BaseModelInfo) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = strconv.FormatInt(snowflake.GetFlake().GetNextID(), 10)
	if obj.Code == "" {
		obj.Code = strconv.Itoa(int(uuid.New().ID()))
	}

	return
}
