package models

import (
	"go-admin/common/models"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	models.ControlBy

	UserId int64 `json:"userId" gorm:"type:int(20);comment:用户id"` //
	Level  int64 `json:"level" gorm:"type:int(20);comment:装扮等级"`  //
}

func (Test) TableName() string {
	return "user_love"
}

func (e *Test) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Test) GetId() interface{} {
	return e.ID
}
