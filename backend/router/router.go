package router

import (
	apiV1 "github.com/ChocolateAceCream/telescope/backend/api/v1"
	"github.com/ChocolateAceCream/telescope/backend/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())
	RouteLoader(r)
}

// gin match the route based on First defined first match rule.
func RouteLoader(r *gin.Engine) {
	authApi := apiV1.AuthApi{}
	sseApi := apiV1.SSEApi{}
	jobApi := apiV1.JobApi{}
	awsApi := apiV1.AwsApi{}
	localeApi := apiV1.LocaleApi{}
	userApi := apiV1.UserApi{}
	sketchApi := apiV1.SketchApi{}
	projectApi := apiV1.ProjectApi{}
	v1 := r.Group("/api/v1")
	PublicGroup := v1.Group("/public")
	{

	}
	auth := PublicGroup.Group("/auth")
	// auth.Use(middleware.DefaultLimiter()).Use(middleware.SessionMiddleware())
	{
		auth.POST("/login", authApi.Login)
		auth.GET("/google/callback", authApi.GoogleLogin)
		auth.POST("renew-session", authApi.RefreshToken)
		auth.POST("/send-code", authApi.SendCode)
		auth.POST("/register", authApi.Register)
	}

	locale := PublicGroup.Group("/locale")
	{
		locale.POST("/reload", localeApi.LoadTranslation)
	}

	PrivateGroup := v1.Group("")
	PrivateGroup.Use(middleware.SessionMiddleware())
	user := PrivateGroup.Group("/user")
	{
		user.GET("/info", userApi.GetUserInfo)
		user.POST("/logout", userApi.Logout)
	}
	sse := PrivateGroup.Group("/sse")
	{
		sse.GET("/subscribe", middleware.SSEMiddleware(), sseApi.Subscriber)
	}
	job := PrivateGroup.Group("/service")
	{
		job.POST("/upload", jobApi.Upload)
	}
	aws := PrivateGroup.Group("/aws")
	{
		aws.POST("/generate-presigned-url", awsApi.GeneratePresignedUrl)
		aws.POST("/classify", awsApi.Classify)
		aws.GET("/download", awsApi.Download)
	}
	sketch := PrivateGroup.Group("/sketch")
	{
		sketch.POST("/upload", sketchApi.UploadSketch)
		sketch.GET("/list", sketchApi.GetSketchList)
		sketch.GET("/detail", sketchApi.GetSketchDetail)
		sketch.PUT("/update", sketchApi.UpdateSketch)
		sketch.DELETE("/delete", sketchApi.DeleteSketch)
	}
	project := PrivateGroup.Group("/project")
	{
		project.GET("/list", projectApi.GetProjectList)
		project.GET("/details/:id", projectApi.GetProjectDetails)
	}
}
