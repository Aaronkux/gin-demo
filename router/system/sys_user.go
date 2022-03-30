package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("register", userApi.Register)
		userRouter.POST("uploadAvatar", userApi.UploadAvatar)
		userRouter.POST("updateSelf", userApi.UpdateSelf)
		userRouter.POST("getUserList", userApi.GetUserList)
		userRouter.POST("getUserById", userApi.GetUserById)
		userRouter.POST("updateUser", userApi.UpdateUser)
		userRouter.POST("deleteUser", userApi.DeleteUser)
	}
}
