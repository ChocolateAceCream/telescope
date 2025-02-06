package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/lib"
	"github.com/ChocolateAceCream/telescope/backend/router"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/ChocolateAceCream/telescope/backend/utils/dataInitializer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init() (r *gin.Engine, err error) {
	dir, _ := os.Getwd()

	var mode string
	flag.StringVar(&mode, "mode", "debug", "gin mode: release mode, default is debug mode")
	flag.Parse()
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	singleton.Viper = utils.ViperInit(dir)

	singleton.Logger = lib.LoggerInit()

	utils.InitRedis()
	err = utils.InitDB()
	if err != nil {
		return
	}
	err = utils.Migrate()
	if err != nil {
		return
	}
	err = dataInitializer.InitData()
	if err != nil {
		return
	}

	err = utils.InitTranslation()
	if err != nil {
		return
	}

	// TODO: check config to see if this function is enabled. If enabled, check kafka connection
	// err = workers.StartWorkerPool()
	// if err != nil {
	// 	return
	// }

	singleton.AWS = utils.NewAWS(utils.WithS3)

	r = gin.Default()
	router.RouterInit(r)
	return
}

func main() {
	r, err := Init()
	if err != nil {
		return
	}
	defer singleton.Redis.Close()
	defer utils.DB.Close()
	s := &http.Server{
		Addr:        ":4050",
		Handler:     r,
		ReadTimeout: 20 * time.Second, //request timeout
		// WriteTimeout:   20 * time.Second, //response timeout
		MaxHeaderBytes: 1 << 20, //default, 1MB
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			singleton.Logger.Error("server closed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	if singleton.Redis == nil {
		quit <- syscall.SIGINT
	} else {
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	}
	<-quit
	singleton.Logger.Info("Shuting down Server ...")

	// 3 setup withTimeout to preserve connection before close
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		singleton.Logger.Error("Server Shutdown", zap.Error(err))
	}
	singleton.Logger.Info("Server exist successful")
}
