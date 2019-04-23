package redisUtils

import (
	"code.byted.org/gopkg/pkg/log"
	"github.com/garyburd/redigo/redis"
)

var redis_conn redis.Conn

func init(){
	var err error
	redis_conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil{
		log.Info("connect to redis error", err)
		return
	}
	log.Info("connect to redis success")

}
