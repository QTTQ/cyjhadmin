package router

import (
	"github.com/gin-gonic/gin"

	"go-admin/app/admin/middleware"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTestRouter)
}

// 需认证的路由代码
func registerTestRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r := v1.Group("/test").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		model := &models.Tags{}
		type TestControl struct {
			ID     uint  `uri:"ID" comment:""` //
			UserID int64 `json:"userId" comment:"标签索引"`
			Level  int64 `json:"level" comment:"类型名"`
		}
		r.GET("", actions.PermissionAction(), actions.IndexTestAction(model, new(dto.TagsSearch), func() interface{} {
			list := make([]models.Tags, 0)
			return &list
		}))
		r.GET("/:id", actions.PermissionAction(), actions.ViewTestAction(new(dto.TestById), func() interface{} {
			return &TestControl{}
		}))
		r.POST("", actions.CreateAction(new(dto.TagsControl)))
		r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.TagsControl)))
		r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.TagsById)))
	}
}
