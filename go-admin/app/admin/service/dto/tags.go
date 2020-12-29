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

type TagsSearch struct {
	dto.Pagination `search:"-"`
	CreatedAt      time.Time `form:"createdAt" search:"type:exact;column:created_at;table:tags" comment:"创建时间"`

	TagIndex int64 `form:"tagIndex" search:"type:exact;column:tag_index;table:tags" comment:"标签索引"`

	TagName int64 `form:"tagName" search:"type:contains;column:tag_name;table:tags" comment:"标签名"`

	TypeIndex int64 `form:"typeIndex" search:"type:exact;column:type_index;table:tags" comment:"标签索引"`

	TypeName string `form:"typeName" search:"type:contains;column:type_name;table:tags" comment:"类型名"`
}

func (m *TagsSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *TagsSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *TagsSearch) Generate() dto.Index {
	o := *m
	return &o
}

type TagsControl struct {
	ID uint `uri:"ID" comment:""` //

	TagIndex int64 `json:"tagIndex" comment:"标签索引"`

	TagName int64 `json:"tagName" comment:"标签名"`

	TypeIndex int64 `json:"typeIndex" comment:"标签索引"`

	TypeName string `json:"typeName" comment:"类型名"`
}

func (s *TagsControl) Bind(ctx *gin.Context) error {
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

func (s *TagsControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *TagsControl) GenerateM() (common.ActiveRecord, error) {
	return &models.Tags{

		Model:     gorm.Model{ID: s.ID},
		TagIndex:  s.TagIndex,
		TagName:   s.TagName,
		TypeIndex: s.TypeIndex,
		TypeName:  s.TypeName,
	}, nil
}

func (s *TagsControl) GetId() interface{} {
	return s.ID
}

type TagsById struct {
	dto.ObjectById
}

func (s *TagsById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *TagsById) GenerateM() (common.ActiveRecord, error) {
	return &models.Tags{}, nil
}
