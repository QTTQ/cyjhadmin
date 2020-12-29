package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type Tags struct {
    gorm.Model
    models.ControlBy
    
    TagIndex int64 `json:"tagIndex" gorm:"type:int(20);comment:标签索引"` // 
    TagName int64 `json:"tagName" gorm:"type:int(20);comment:标签名"` // 
    TypeIndex int64 `json:"typeIndex" gorm:"type:int(2);comment:标签索引"` // 
    TypeName string `json:"typeName" gorm:"type:varchar(200);comment:类型名"` // 
}

func (Tags) TableName() string {
    return "tags"
}

func (e *Tags) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Tags) GetId() interface{} {
	return e.ID
}
