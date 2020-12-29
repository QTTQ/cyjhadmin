package models

import (
	"time"

	"gorm.io/gorm"

	"go-admin/common/models"
)

type UserInfo struct {
	gorm.Model
	models.ControlBy

	NickName      string    `json:"nickName" gorm:"type:varchar(100);comment:昵称"`             //
	BirthDate     time.Time `json:"birthDate" gorm:"type:datetime;comment:生日"`                //
	Sex           string    `json:"sex" gorm:"type:int(2);comment:性别"`                        //
	Autograph     string    `json:"autograph" gorm:"type:varchar(100);comment:Autograph"`     //
	RealName      string    `json:"realName" gorm:"type:varchar(64);comment:真实姓名"`            //
	JhNum         string    `json:"jhNum" gorm:"type:varchar(100);comment:江湖号"`               //
	UnionNum      string    `json:"unionNum" gorm:"type:varchar(100);comment:UnionNum"`       //
	PhoneNum      string    `json:"phoneNum" gorm:"type:varchar(100);comment:电话号"`            //
	Password      string    `json:"password" gorm:"type:varchar(100);comment:Password"`       //
	Actor         string    `json:"actor" gorm:"type:varchar(200);comment:头像"`                //
	Email         string    `json:"email" gorm:"type:varchar(40);comment:Email"`              //
	QqOpenId      string    `json:"qqOpenId" gorm:"type:varchar(60);comment:QqOpenId"`        //
	WxOpenId      string    `json:"wxOpenId" gorm:"type:varchar(60);comment:WxOpenId"`        //
	DistId        string    `json:"distId" gorm:"type:int(20);comment:DistId"`                //
	DeviceThemeId string    `json:"deviceThemeId" gorm:"type:int(20);comment:DeviceThemeId"`  //
	DomainName    string    `json:"domainName" gorm:"type:varchar(60);comment:DomainName"`    //
	StaffNum      string    `json:"staffNum" gorm:"type:varchar(100);comment:StaffNum"`       //
	Province      string    `json:"province" gorm:"type:varchar(20);comment:Province"`        //
	City          string    `json:"city" gorm:"type:varchar(20);comment:City"`                //
	Category      string    `json:"category" gorm:"type:int(2);comment:Category"`             //
	PayPassword   string    `json:"payPassword" gorm:"type:varchar(100);comment:PayPassword"` //
	IsfollowWx    string    `json:"isfollowWx" gorm:"type:int(2);comment:IsfollowWx"`         //
	FirstRecharge string    `json:"firstRecharge" gorm:"type:int(2);comment:FirstRecharge"`   //
}

func (UserInfo) TableName() string {
	return "user_info"
}

func (e *UserInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserInfo) GetId() interface{} {
	return e.ID
}
