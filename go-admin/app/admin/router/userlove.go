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
	routerCheckRole = append(routerCheckRole, registerUserLoveRouter)
}

// 需认证的路由代码
func registerUserLoveRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/userlove").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.UserLove{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.UserLoveSearch), func() interface{} {
            list := make([]models.UserLove, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.UserLoveById), nil))
        r.POST("", actions.CreateAction(new(dto.UserLoveControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.UserLoveControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.UserLoveById)))
    }
}
