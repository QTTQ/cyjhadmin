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

type TestSearch struct {
	dto.Pagination `search:"-"`
	CreatedAt      time.Time `form:"createdAt" search:"type:exact;column:created_at;table:tags" comment:"创建时间"`

	TagIndex int64 `form:"tagIndex" search:"type:exact;column:tag_index;table:tags" comment:"标签索引"`

	TagName int64 `form:"tagName" search:"type:contains;column:tag_name;table:tags" comment:"标签名"`

	TypeIndex int64 `form:"typeIndex" search:"type:exact;column:type_index;table:tags" comment:"标签索引"`

	TypeName string `form:"typeName" search:"type:contains;column:type_name;table:tags" comment:"类型名"`
}

func (m *TestSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *TestSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *TestSearch) Generate() dto.Index {
	o := *m
	return &o
}

type TestControl struct {
	ID uint `uri:"ID" comment:""` //

	UserID int64 `json:"userId" comment:"标签索引"`
	Level  int64 `json:"level" comment:"类型名"`
}

func (s *TestControl) Bind(ctx *gin.Context) error {
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

func (s *TestControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *TestControl) GenerateM() (common.ActiveRecord, error) {
	return &models.Test{

		Model:  gorm.Model{ID: s.ID},
		UserId: s.UserID,
		Level:  s.Level,
	}, nil
}

func (s *TestControl) GetId() interface{} {
	return s.ID
}

type TestById struct {
	dto.ObjectById
}

func (s *TestById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *TestById) GenerateM() (common.ActiveRecord, error) {
	return &models.Test{}, nil
}
