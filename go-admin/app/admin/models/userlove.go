package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type UserLove struct {
    gorm.Model
    models.ControlBy
    
    UserId string `json:"userId" gorm:"type:int(20);comment:用户id"` // 
    TypeIndex string `json:"typeIndex" gorm:"type:int(2);comment:类型标签"` // 
    BlockId string `json:"blockId" gorm:"type:int(20);comment:区块id"` // 
    TagIndex string `json:"tagIndex" gorm:"type:int(20);comment:标签索引"` // 
}

func (UserLove) TableName() string {
    return "user_love"
}

func (e *UserLove) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserLove) GetId() interface{} {
	return e.ID
}
