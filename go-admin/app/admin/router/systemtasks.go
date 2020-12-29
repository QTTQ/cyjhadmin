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
	routerCheckRole = append(routerCheckRole, registerSystemTasksRouter)
}

// 需认证的路由代码
func registerSystemTasksRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/systemtasks").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.SystemTasks{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.SystemTasksSearch), func() interface{} {
            list := make([]models.SystemTasks, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.SystemTasksById), nil))
        r.POST("", actions.CreateAction(new(dto.SystemTasksControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.SystemTasksControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.SystemTasksById)))
    }
}
