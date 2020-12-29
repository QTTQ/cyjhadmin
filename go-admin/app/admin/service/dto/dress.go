package dto

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type DressSearch struct {
	dto.Pagination `search:"-"`
	ID             uint `form:"ID" search:"type:exact;column:id;table:dress" comment:"id"`

	Title string `form:"title" search:"type:contains;column:title;table:dress" comment:"装扮名字"`

	SmTitle string `form:"smTitle" search:"type:contains;column:sm_title;table:dress" comment:"装扮小名"`

	TypeIndex int64 `form:"typeIndex" search:"type:exact;column:type_index;table:dress" comment:"装扮类型索引"`
}

func (m *DressSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *DressSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *DressSearch) Generate() dto.Index {
	o := *m
	return &o
}

type DressControl struct {
	ID uint `uri:"ID" comment:"id"` // id

	Title string `json:"title" comment:"装扮名字"`

	SmTitle string `json:"smTitle" comment:"装扮小名"`

	TypeIndex int64 `json:"typeIndex" comment:"装扮类型索引"`

	TypeName string `json:"typeName" comment:"装扮类型"`

	Desc string `json:"desc" comment:"装扮描述"`

	Icon string `json:"icon" comment:"装扮icon"`

	Url string `json:"url" comment:"装扮图片地址"`

	Remarks string `json:"remarks" comment:"备注"`

	Tags string `json:"tags" comment:"标签"`

	Level int64 `json:"level" comment:"装扮等级"`
}

func (s *DressControl) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
		return err
	}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	json.Unmarshal(data, &s)
	// err = ctx.ShouldBind(s)
	// if err != nil {
	// 	log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
	// }
	return err
}

func (s *DressControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *DressControl) GenerateM() (common.ActiveRecord, error) {
	return &models.Dress{

		Model:     gorm.Model{ID: s.ID},
		Title:     s.Title,
		SmTitle:   s.SmTitle,
		TypeIndex: s.TypeIndex,
		TypeName:  s.TypeName,
		Desc:      s.Desc,
		Icon:      s.Icon,
		Url:       s.Url,
		Remarks:   s.Remarks,
		Tags:      s.Tags,
		Level:     s.Level,
	}, nil
}

func (s *DressControl) GetId() interface{} {
	return s.ID
}

type DressById struct {
	dto.ObjectById
}

func (s *DressById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *DressById) GenerateM() (common.ActiveRecord, error) {
	return &models.Dress{}, nil
}
