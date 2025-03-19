package router

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) SetupRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/users")
	{
		userRouter.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "user router active",
			})
		})
	}

}
