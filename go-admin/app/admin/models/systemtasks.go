package models

import (
    "gorm.io/gorm"

	"go-admin/common/models"
)

type SystemTasks struct {
    gorm.Model
    models.ControlBy
    
    TypeIndex int64 `json:"typeIndex" gorm:"type:int(2);comment:标签索引"` // 
    TypeName string `json:"typeName" gorm:"type:varchar(40);comment:标签名"` // 
    Content string `json:"content" gorm:"type:longtext;comment:任务内容"` // 
    JumpUrl string `json:"jumpUrl" gorm:"type:varchar(300);comment:跳转地址"` // 
    StaticUrl string `json:"staticUrl" gorm:"type:varchar(300);comment:静态地址"` // 
    TasksUuid string `json:"tasksUuid" gorm:"type:varchar(100);comment:TasksUuid"` // 
    UpTkUuid string `json:"upTkUuid" gorm:"type:varchar(100);comment:UpTkUuid"` // 
    Reward string `json:"reward" gorm:"type:decimal(11,2);comment:奖励金"` // 
    IsRelat int64 `json:"isRelat" gorm:"type:int(2);comment:是否需要关联上级"` // 
}

func (SystemTasks) TableName() string {
    return "system_tasks"
}

func (e *SystemTasks) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SystemTasks) GetId() interface{} {
	return e.ID
}
