package global

import (
	"eva/biz/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"

	"eva/biz/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	EVA_DB      *gorm.DB
	EVA_DBList  map[string]*gorm.DB
	EVA_REDIS   *redis.Client
	EVA_MongoDB *mongo.Client
	EVA_CONFIG  config.Server
	EVA_VP      *viper.Viper
	// EVA_LOG    *oplogging.Logger
	EVA_LOG                 *zap.Logger
	EVA_Timer               timer.Timer = timer.NewTimerTask()
	EVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return EVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := EVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
