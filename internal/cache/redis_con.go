package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/settings"
	"sync"
	"time"
)

type Interface interface {
	Ping()
}

type Redis struct{}

var singleton *Redis
var once sync.Once

var redisPool *redis.Pool

func NewRedisConnect() Interface {
	once.Do(func() {
		singleton = &Redis{}
	})
	return singleton
}

func (r *Redis) RedisPoolConnect() *redis.Pool {
	if redisPool != nil {
		return redisPool
	}

	redisPool = &redis.Pool{
		MaxIdle:     100,               // 最大可允許的閒置連線數
		MaxActive:   10000,             // 最大建立的連線數
		IdleTimeout: 300 * time.Second, // 連線過期時間
		Wait:        true,              // 當連線超出限制數量後，是否等待到空閒連線釋放
		Dial: func() (r redis.Conn, err error) {

			r, err = redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", settings.Config.RedisConfig.Host, settings.Config.RedisConfig.Port),
				redis.DialPassword(settings.Config.RedisConfig.Password),
				redis.DialConnectTimeout(5*time.Second), // 建立連線 timeout
				redis.DialReadTimeout(5*time.Second),    // 讀取資料 time out
				redis.DialWriteTimeout(5*time.Second),   // 寫入資料 time out
			)

			if err != nil {
				log.Fatalf(errorcode.RedisConnectError, err.Error())
				return
			}

			return
		},
		TestOnBorrow: func(r redis.Conn, t time.Time) (err error) {
			if time.Since(t) < (5 * time.Second) {
				return
			}

			_, err = r.Do("PING")
			if err != nil {
				log.Fatalf(errorcode.PingRedisError, err.Error())
				return
			}

			return
		},
	}

	return redisPool
}

func (r *Redis) Ping() {
	connPool := r.RedisPoolConnect()
	conn := connPool.Get()

	_, err := conn.Do("PING")
	if err != nil {
		conn.Close()
		log.Fatalf(errorcode.PingRedisError, err.Error())
	}
}
