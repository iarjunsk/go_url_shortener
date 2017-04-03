//Models are data structures for representing my database concepts.
package models

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"go_shortify_web_app_heroku/controllers"
)


var redisPool *redis.Pool

func Redis_db_init(){
	redisAddr :=  "redis-XXXX.c9.us-east-1-2.ec2.cloud.redislabs.com:XXXX"
	redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisAddr)
			return conn, err
		},
	}
}

func Redis_db_save(long_url string) (string) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	new_short_code := controllers.Hash(long_url) // Hash the long url
	redisConn.Do("SET", new_short_code, long_url) // save to db
	return fmt.Sprint(new_short_code)
}

func Redis_db_get(shortCode string) (string,error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	redirect_url, err := redis.String(redisConn.Do("GET", shortCode))
	return redirect_url,err
}