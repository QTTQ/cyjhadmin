package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type Dress struct {
    gorm.Model
    models.ControlBy
    
    Title string `json:"title" gorm:"type:varchar(100);comment:装扮名字"` // 
    SmTitle string `json:"smTitle" gorm:"type:varchar(100);comment:装扮小名"` // 
    TypeIndex int64 `json:"typeIndex" gorm:"type:int(2);comment:装扮类型索引"` // 
    TypeName string `json:"typeName" gorm:"type:varchar(200);comment:装扮类型"` // 
    Desc string `json:"desc" gorm:"type:varchar(200);comment:装扮描述"` // 
    Icon string `json:"icon" gorm:"type:varchar(200);comment:装扮icon"` // 
    Url string `json:"url" gorm:"type:varchar(200);comment:装扮图片地址"` // 
    Remarks string `json:"remarks" gorm:"type:varchar(200);comment:备注"` // 
    Tags string `json:"tags" gorm:"type:varchar(50);comment:标签"` // 
    Level int64 `json:"level" gorm:"type:int(20);comment:装扮等级"` // 
}

func (Dress) TableName() string {
    return "dress"
}

func (e *Dress) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Dress) GetId() interface{} {
	return e.ID
}
