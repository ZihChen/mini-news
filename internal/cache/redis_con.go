package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"mini-news/app/global/settings"
	"sync"
	"time"
)

type Interface interface {
	Ping()
	CheckHashKeyExist(hashKey string) (exist bool, goErr errorcode.Error)
	Set(key string, value interface{}, expire int) (goErr errorcode.Error)
	HashSet(hashKey string, key, value interface{}, expire int) (goErr errorcode.Error)
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

func (r *Redis) redisPoolConnect() *redis.Pool {
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
	connPool := r.redisPoolConnect()
	conn := connPool.Get()

	defer func() {
		conn.Close()
	}()

	_, err := conn.Do("PING")
	if err != nil {
		conn.Close()
		log.Fatalf(errorcode.PingRedisError, err.Error())
	}
}

func (r *Redis) CheckHashKeyExist(hashKey string) (exist bool, goErr errorcode.Error) {
	RedisPool := r.redisPoolConnect()
	conn := RedisPool.Get()

	defer func() {
		conn.Close()
	}()

	result, _ := conn.Do("EXISTS", hashKey)
	exist, err := redis.Bool(result, nil)
	if err != nil {
		helper.ErrorHandle(errorcode.ErrorRedis, errorcode.CheckHashKeyExistError, err.Error())
	}

	return
}

func (r *Redis) Set(key string, value interface{}, expire int) (goErr errorcode.Error) {
	RedisPool := r.redisPoolConnect()
	conn := RedisPool.Get()

	defer func() {
		conn.Close()
	}()

	if _, err := conn.Do("SET", key, value, "EX", expire); err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorRedis, errorcode.RedisSetError, err.Error())
		return
	}

	return
}

func (r *Redis) HashSet(hashKey string, key, value interface{}, expire int) (goErr errorcode.Error) {
	RedisPool := r.redisPoolConnect()
	conn := RedisPool.Get()

	defer func() {
		conn.Close()
	}()

	exist, _ := r.CheckHashKeyExist(hashKey)

	if _, err := conn.Do("HSET", hashKey, key, value); err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorRedis, errorcode.RedisHashSetError, err.Error())
		return
	}

	if exist == false {
		if _, err := conn.Do("EXPIRE", hashKey, expire); err != nil {
			goErr = helper.ErrorHandle(errorcode.ErrorRedis, errorcode.RedisSetExpireError, err.Error())
		}
	}

	return
}
