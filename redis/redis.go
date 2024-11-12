package redis

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/gopering/common-fast/errorx"
	logx "github.com/gopering/common-fast/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	MainDB  *redis.Redis
	MainDB2 *redis2.Client
)

const (
	TypeMain = "main" // 主库
)

// RedisCfg  redis配置
type RedisCfg struct {
	Name string
	redis.RedisConf
	DB int
}

// logError 记录错误日志
func logError(ctx context.Context, operation, key string, err error) {
	logx.ErrorContextf(ctx, "%s Failed, err:%s, key:%s", operation, err, key)
}

// handleRedisError 处理 Redis 错误
func handleRedisError(ctx context.Context, key string, err error) *errorx.CodeErrorResponseContent {
	logError(ctx, "Redis operation", key, err)
	return errorx.ErrCodeInternal.GenError()
}

// setKeyWithExpire 设置键值并设置过期时间
func SetKeyWithExpire(ctx context.Context, key, value string, expire int) *errorx.CodeErrorResponseContent {
	err := MainDB.SetexCtx(ctx, key, value, expire)
	if err != nil {
		return handleRedisError(ctx, key, err)
	}
	return nil
}

// Init 初始化
func Init(list []RedisCfg) {
	for _, item := range list {
		item := item
		switch item.Name {
		case TypeMain:
			logx.InfoContextf(context.Background(), "当前redis db: %d", item.DB)
			fmt.Println("当前redis db: ", item.DB)
			MainDB = item.RedisConf.NewRedis()

			MainDB2 = redis2.NewClient(&redis2.Options{
				Addr:     item.Host,
				Password: item.Pass, // 如果有密码，需要设置密码
				DB:       item.DB,   // 使用默认的数据库
			})
		}
	}
}

func InitRedis() {

	/*	rdb := rdb.NewClient(&rdb.Options{
			Addr:     "r-wz9xxvbfgrutk53a1hpd.redis.rds.aliyuncs.com:6379",
			Password: "Zyxx888888", // 如果有密码，需要设置密码
			DB:       1,            // 使用默认的数据库
		})

		MainDB. = rdb*/

}
