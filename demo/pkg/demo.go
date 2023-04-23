package pkg

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type XRedis redis.Client

var xr *XRedis

func NewRedisManager() *XRedis {
	once.Do(func() {
		xr = new(XRedis)
		xr = Init()
	})
	return xr
}

var once sync.Once

// Init init redis manager
func Init() *XRedis {
	// init redis client
	xr1 := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:12121",
		Username:    "",
		Password:    "aaaaa+",
		DB:          0,
		ReadTimeout: 2 * time.Second, // default read and write timeout is 2s
	})
	//xr3 := XRedis(xr1)

	xr2 := (*XRedis)(xr1)

	_, err := xr2.Ping(context.Background()).Result()

	if err != nil {
		panic(fmt.Sprintf("ping redis [%s] failed, error:%s", "10.90.73", err.Error()))
	}
	return xr2
}

type Config struct {
	Name     string // instance name
	Addr     string // host:port address.
	Username string // username
	Password string // password
	DB       int    // selected db
	PoolSize int    // connection pool size, must > 3
}
