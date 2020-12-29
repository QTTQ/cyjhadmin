package actions

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/common/dto"
	"go-admin/common/log"
	"go-admin/common/models"
	"go-admin/tools"
	"go-admin/tools/app"
)

// IndexAction 通用查询动作
func IndexTestAction(m models.ActiveRecord, d dto.Index, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := tools.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := tools.GenerateMsgIDFromContext(c)
		list := f()
		object := m.Generate()
		req := d.Generate()
		var count int64

		//查询列表
		err = req.Bind(c)
		if err != nil {
			app.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
			return
		}

		//数据权限检查
		p := GetPermissionFromContext(c)

		err = db.WithContext(c).Model(object).
			Scopes(
				dto.MakeCondition(req.GetNeedSearch()),
				dto.Paginate(req.GetPageSize(), req.GetPageIndex()),
				Permission(object.TableName(), p),
			).
			Find(list).Limit(-1).Offset(-1).
			Count(&count).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("MsgID[%s] Index error: %s", msgID, err)
			app.Error(c, http.StatusInternalServerError, err, "查询失败")
			return
		}
		app.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
		c.Next()
	}
}

// ViewAction 通用详情动作
func ViewTestAction(control dto.Control, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := tools.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := tools.GenerateMsgIDFromContext(c)
		//查看详情
		req := control.Generate()
		err = req.Bind(c)
		if err != nil {
			app.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			app.Error(c, http.StatusInternalServerError, err, "模型生成失败")
			return
		}

		var rsp interface{}
		if f != nil {
			rsp = f()
			fmt.Println(f(), "-----------------f ---")
		} else {
			rsp, _ = req.GenerateM()
		}

		//数据权限检查
		p := GetPermissionFromContext(c)
		// // ------------------
		// dto.MakeCondition(req.GetNeedSearch()),
		// 	dto.Paginate(req.GetPageSize(), req.GetPageIndex()),
		// 	Permission(object.TableName(), p),
		// 	// ------------------
		fmt.Println(object, "----------------")
		err = db.Debug().Model(object).WithContext(c).Scopes(
			Permission(object.TableName(), p),
		).Select("user_love.user_id,dress.level").Where(req.GetId()).Joins("LEFT JOIN dress on dress.level = user_love.user_id").First(rsp).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			app.Error(c, http.StatusNotFound, nil, "查看对象不存在或无权查看")
			return
		}
		if err != nil {
			log.Errorf("MsgID[%s] View error: %s", msgID, err)
			app.Error(c, http.StatusInternalServerError, err, "查看失败")
			return
		}
		app.OK(c, rsp, "查看成功")
		c.Next()
	}
}
