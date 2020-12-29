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

type SystemTasksSearch struct {
	dto.Pagination     `search:"-"`
    TypeIndex int64 `form:"typeIndex" search:"type:exact;column:type_index;table:system_tasks" comment:"标签索引"`

    Reward string `form:"reward" search:"type:exact;column:reward;table:system_tasks" comment:"奖励金"`

    IsRelat int64 `form:"isRelat" search:"type:exact;column:is_relat;table:system_tasks" comment:"是否需要关联上级"`

    
}

func (m *SystemTasksSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *SystemTasksSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *SystemTasksSearch) Generate() dto.Index {
	o := *m
	return &o
}

type SystemTasksControl struct {
    
    ID uint `uri:"ID" comment:"id"` // id

    TypeIndex int64 `json:"typeIndex" comment:"标签索引"`
    

    TypeName string `json:"typeName" comment:"标签名"`
    

    Content string `json:"content" comment:"任务内容"`
    

    JumpUrl string `json:"jumpUrl" comment:"跳转地址"`
    

    StaticUrl string `json:"staticUrl" comment:"静态地址"`
    

    TasksUuid string `json:"tasksUuid" comment:""`
    

    UpTkUuid string `json:"upTkUuid" comment:""`
    

    Reward string `json:"reward" comment:"奖励金"`
    

    IsRelat int64 `json:"isRelat" comment:"是否需要关联上级"`
    
}

func (s *SystemTasksControl) Bind(ctx *gin.Context) error {
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

func (s *SystemTasksControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *SystemTasksControl) GenerateM() (common.ActiveRecord, error) {
	return &models.SystemTasks{
	
        Model:     gorm.Model{ID: s.ID},
        TypeIndex:  s.TypeIndex,
        TypeName:  s.TypeName,
        Content:  s.Content,
        JumpUrl:  s.JumpUrl,
        StaticUrl:  s.StaticUrl,
        TasksUuid:  s.TasksUuid,
        UpTkUuid:  s.UpTkUuid,
        Reward:  s.Reward,
        IsRelat:  s.IsRelat,
	}, nil
}

func (s *SystemTasksControl) GetId() interface{} {
	return s.ID
}

type SystemTasksById struct {
	dto.ObjectById
}

func (s *SystemTasksById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *SystemTasksById) GenerateM() (common.ActiveRecord, error) {
	return &models.SystemTasks{}, nil
}
