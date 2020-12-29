package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type UserInfoSearch struct {
	dto.Pagination `search:"-"`
	ID             uint `form:"ID" search:"type:exact;column:id;table:user_info" comment:"id"`

	NickName string `form:"nickName" search:"type:contains;column:nick_name;table:user_info" comment:"昵称"`

	RealName string `form:"realName" search:"type:contains;column:real_name;table:user_info" comment:"真实姓名"`

	PhoneNum string `form:"phoneNum" search:"type:exact;column:phone_num;table:user_info" comment:"电话号"`
}

func (m *UserInfoSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *UserInfoSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *UserInfoSearch) Generate() dto.Index {
	o := *m
	return &o
}

type UserInfoControl struct {
	ID uint `uri:"ID" comment:"id"` // id

	NickName string `json:"nickName" comment:"昵称"`

	BirthDate time.Time `json:"birthDate" comment:"生日"`

	Sex string `json:"sex" comment:"性别"`

	Autograph string `json:"autograph" comment:""`

	RealName string `json:"realName" comment:"真实姓名"`

	JhNum string `json:"jhNum" comment:"江湖号"`

	UnionNum string `json:"unionNum" comment:""`

	PhoneNum string `json:"phoneNum" comment:"电话号"`

	Password string `json:"password" comment:""`

	Actor string `json:"actor" comment:"头像"`

	Email string `json:"email" comment:""`

	QqOpenId string `json:"qqOpenId" comment:""`

	WxOpenId string `json:"wxOpenId" comment:""`

	DistId string `json:"distId" comment:""`

	DeviceThemeId string `json:"deviceThemeId" comment:""`

	DomainName string `json:"domainName" comment:""`

	StaffNum string `json:"staffNum" comment:""`

	Province string `json:"province" comment:""`

	City string `json:"city" comment:""`

	Category string `json:"category" comment:""`

	PayPassword string `json:"payPassword" comment:""`

	IsfollowWx string `json:"isfollowWx" comment:""`

	FirstRecharge string `json:"firstRecharge" comment:""`
}

func (s *UserInfoControl) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
		return err
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
	}
	return err
}

func (s *UserInfoControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *UserInfoControl) GenerateM() (common.ActiveRecord, error) {
	return &models.UserInfo{

		Model:         gorm.Model{ID: s.ID},
		NickName:      s.NickName,
		BirthDate:     s.BirthDate,
		Sex:           s.Sex,
		Autograph:     s.Autograph,
		RealName:      s.RealName,
		JhNum:         s.JhNum,
		UnionNum:      s.UnionNum,
		PhoneNum:      s.PhoneNum,
		Password:      s.Password,
		Actor:         s.Actor,
		Email:         s.Email,
		QqOpenId:      s.QqOpenId,
		WxOpenId:      s.WxOpenId,
		DistId:        s.DistId,
		DeviceThemeId: s.DeviceThemeId,
		DomainName:    s.DomainName,
		StaffNum:      s.StaffNum,
		Province:      s.Province,
		City:          s.City,
		Category:      s.Category,
		PayPassword:   s.PayPassword,
		IsfollowWx:    s.IsfollowWx,
		FirstRecharge: s.FirstRecharge,
	}, nil
}

func (s *UserInfoControl) GetId() interface{} {
	return s.ID
}

type UserInfoById struct {
	dto.ObjectById
}

func (s *UserInfoById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *UserInfoById) GenerateM() (common.ActiveRecord, error) {
	return &models.UserInfo{}, nil
}
