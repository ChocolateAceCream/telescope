package router

import (
	apiV1 "github.com/ChocolateAceCream/telescope/backend/api/v1"
	"github.com/ChocolateAceCream/telescope/backend/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	// r.Use(middleware.CORSMiddleware())
	RouteLoader(r)
}

// gin match the route based on First defined first match rule.
func RouteLoader(r *gin.Engine) {
	authApi := apiV1.AuthApi{}
	sseApi := apiV1.SSEApi{}
	jobApi := apiV1.JobApi{}
	awsApi := apiV1.AwsApi{}
	localeApi := apiV1.LocaleApi{}
	v1 := r.Group("/api/v1")
	v1.Use(middleware.SessionMiddleware())
	PublicGroup := v1.Group("/public")
	{

	}
	auth := PublicGroup.Group("/auth")
	// auth.Use(middleware.DefaultLimiter()).Use(middleware.SessionMiddleware())
	{
		auth.POST("/login", authApi.Login)
	}

	locale := PublicGroup.Group("/locale")
	{
		locale.POST("/reload", localeApi.LoadTranslation)
	}

	PrivateGroup := v1.Group("")
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
		aws.POST("/upload", awsApi.Upload)
		aws.GET("/download", awsApi.Download)
	}
}
