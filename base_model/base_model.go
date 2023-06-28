package base_model

import (
	"github.com/ShiLiangAPI/goutils/snowflake"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        int64                 `gorm:"<-:create;primaryKey" json:"id"`
	CreatedAt *LocalTime            `gorm:"comment:'创建时间'" json:"created_at"`
	UpdatedAt *LocalTime            `gorm:"comment:'修改时间'" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"<-:create;index" json:"-"`
}

type Model struct {
	ID        int64      `gorm:"<-:create;primaryKey" json:"id"`
	CreatedAt *LocalTime `json:"created_at"`
	UpdatedAt *LocalTime `json:"updated_at"`
}

func (obj *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = snowflake.GetNextID()

	return
}

func (obj *Model) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = snowflake.GetNextID()

	return
}
