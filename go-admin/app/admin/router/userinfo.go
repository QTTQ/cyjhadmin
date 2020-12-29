package router

import (
    "github.com/gin-gonic/gin"

    "go-admin/app/admin/middleware"
    "go-admin/app/admin/models"
    "go-admin/app/admin/service/dto"
    "go-admin/common/actions"
    jwt "go-admin/pkg/jwtauth"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerUserInfoRouter)
}

// 需认证的路由代码
func registerUserInfoRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/userinfo").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.UserInfo{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.UserInfoSearch), func() interface{} {
            list := make([]models.UserInfo, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.UserInfoById), nil))
        r.POST("", actions.CreateAction(new(dto.UserInfoControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.UserInfoControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.UserInfoById)))
    }
}
