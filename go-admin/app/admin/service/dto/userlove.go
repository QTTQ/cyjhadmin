package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type UserLoveSearch struct {
	dto.Pagination     `search:"-"`
    
}

func (m *UserLoveSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *UserLoveSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *UserLoveSearch) Generate() dto.Index {
	o := *m
	return &o
}

type UserLoveControl struct {
    
    ID uint `uri:"ID" comment:"ID"` // ID

    UserId string `json:"userId" comment:"用户id"`
    

    TypeIndex string `json:"typeIndex" comment:"类型标签"`
    

    BlockId string `json:"blockId" comment:"区块id"`
    

    TagIndex string `json:"tagIndex" comment:"标签索引"`
    
}

func (s *UserLoveControl) Bind(ctx *gin.Context) error {
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

func (s *UserLoveControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *UserLoveControl) GenerateM() (common.ActiveRecord, error) {
	return &models.UserLove{
	
        Model:     gorm.Model{ID: s.ID},
        UserId:  s.UserId,
        TypeIndex:  s.TypeIndex,
        BlockId:  s.BlockId,
        TagIndex:  s.TagIndex,
	}, nil
}

func (s *UserLoveControl) GetId() interface{} {
	return s.ID
}

type UserLoveById struct {
	dto.ObjectById
}

func (s *UserLoveById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *UserLoveById) GenerateM() (common.ActiveRecord, error) {
	return &models.UserLove{}, nil
}
