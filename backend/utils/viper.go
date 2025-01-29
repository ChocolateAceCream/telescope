package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/workers"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TODO:
// 优先级: 命令行 > 环境变量 > 默认值

func ViperInit(path string) *viper.Viper {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("config") // name of config file (without extension)
	v.AddConfigPath(".")      // path to look for the config file in
	v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	err := v.ReadInConfig()
	if err != nil {
		// global.LOGGER.Error(fmt.Sprintf("fatal error config file: %s", err))
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	v.WatchConfig() // watch config change, hot reload

	v.OnConfigChange(func(e fsnotify.Event) {
		singleton.Logger.Info(fmt.Sprintf("config file changed: %s", e.Name))
		handleResizerWorkersChange(context.TODO(), v)
	})
	if err = v.UnmarshalKey(gin.Mode(), &singleton.Config); err != nil {
		fmt.Println(err)
	}

	// // root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	// global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	// global.BlackCache = local_cache.NewCache(
	// 	local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
	//

	//test
	// fmt.Println("redis port from config ", v.Get("redis.addr"))
	return v
}

func handleResizerWorkersChange(ctx context.Context, v *viper.Viper) {
	// Add a small delay to ensure config file is fully updated
	time.Sleep(200 * time.Millisecond)
	// Get the current environment (mode)
	currentEnv := gin.Mode()
	oldCount, ok := workers.GetWorkerCountByWorkerPoolType("resizer")
	if !ok {
		singleton.Logger.Error("Failed to get resizer worker count")
		return
	}
	fmt.Println("oldCount", oldCount)
	// Print the environment
	newCount := v.GetInt(currentEnv + "." + "workers.resizer.count")
	fmt.Println("newCount", newCount)

	//When Viper detects a config change, the file might not be fully written to disk yet (depending on how the editor writes to the file), leading to a temporary incomplete state of the config file. This can cause the GetInt() function to return a default value (in this case, 0) during the first trigger of OnConfigChange. To avoid this, we add a small delay before reading the new value.
	if newCount != oldCount && newCount > 0 && oldCount > 0 {
		// Update the configuration
		fmt.Println("update the count")
		singleton.Config.Workers.Resizer.Count = newCount
		workers.RestartWorkerPoolByPoolName("resizer")
		// Update the worker pool
	}
}
