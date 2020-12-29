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
	routerCheckRole = append(routerCheckRole, registerDressRouter)
}

// 需认证的路由代码
func registerDressRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/dress").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.Dress{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.DressSearch), func() interface{} {
            list := make([]models.Dress, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.DressById), nil))
        r.POST("", actions.CreateAction(new(dto.DressControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.DressControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.DressById)))
    }
}
