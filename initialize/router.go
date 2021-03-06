package initialize

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/middleware"
	"gandi.icu/demo/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// Router.StaticFS(global.AM_CONFIG.Local.Path, http.Dir(global.AM_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.AM_LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.AM_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.AM_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System
	// exampleRouter := router.RouterGroupApp.Example
	// autocodeRouter := router.RouterGroupApp.Autocode
	PublicGroup := Router.Group("")
	PrivateGroup := Router.Group("")
	// {
	// 	// 健康监测
	// 	PublicGroup.GET("/health", func(c *gin.Context) {
	// 		c.JSON(200, "ok")
	// 	})
	// }
	// {
	systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	// 	systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	// }
	PrivateGroup.Use(middleware.JWTAuth())
	// .Use(middleware.CasbinHandler())
	// {
	// 	systemRouter.InitJwtRouter(PrivateGroup)                 // jwt相关路由
	systemRouter.InitApiRouter(PrivateGroup)       // 注册功能api路由
	systemRouter.InitUserRouter(PrivateGroup)      // 注册用户路由
	systemRouter.InitAuthorityRouter(PrivateGroup) // 注册用户路由
	systemRouter.InitMenuRouter(PrivateGroup)      // 注册menu路由
	systemRouter.InitCasbinRouter(PrivateGroup)    // 权限相关路由
	systemRouter.InitBranchRouter(PrivateGroup)
	systemRouter.InitSaleRouter(PrivateGroup)
	systemRouter.InitFileRouter(PublicGroup)
	systemRouter.InitReferralRouter(PrivateGroup)
	systemRouter.InitClientRouter(PrivateGroup)
	systemRouter.InitViewRouter(PrivateGroup)
	// 	systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
	// 	systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由

	// 	exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
	// 	exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
	// 	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	// }

	// InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.AM_LOG.Info("router register success")
	return Router
}
