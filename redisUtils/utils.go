package redisUtils

import (
	"code.byted.org/gopkg/pkg/log"
	"github.com/garyburd/redigo/redis"
)

func SetKeyValue(key string, value string)  {
	_, err := redis_conn.Do("SET", key, value)
	if err != nil{
		log.Info("redis set failed:",err)
	}
}

func GetKeyValue(key string) string {
	_, err:= redis.String(redis_conn.Do("get", "string"))
	if err != nil{
		log.Info("redis get failed", err)
	}
}

func AddListValue(listKey string, value string) {
	_, err := redis_conn.Do("LPUSH", listKey, value)
	if err != nil {
		log.Info("redis list value add failed", err)
	}
}

func GetListLength(listKey string) int{
	
}
