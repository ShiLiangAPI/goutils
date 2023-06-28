package base_model

import (
	"gitee.com/ShiLiangAPI/goutils/snowflake"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"strconv"
)

type BaseModel struct {
	ID        string                `gorm:"<-:create;primaryKey" json:"id"`
	CreatedAt *LocalTime            `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt *LocalTime            `gorm:"comment:'修改时间'" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"<-:create;index" json:"-"`
}

type Model struct {
	ID        string     `gorm:"<-:create;primaryKey" json:"id"`
	CreatedAt *LocalTime `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt *LocalTime `gorm:"comment:'修改时间'" json:"updatedAt"`
}

func (obj *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = strconv.FormatInt(snowflake.GetNextID(1, 1), 10)

	return
}

func (obj *Model) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = strconv.FormatInt(snowflake.GetNextID(1, 1), 10)

	return
}
