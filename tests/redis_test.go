package test

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestRedis(t *testing.T) {
	c, err := redis.Dial("tcp", "192.168.1.123:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	//密码授权
	c.Do("AUTH", "123456")
	c.Do("SET", "a", 134)
	a, err := redis.Int(c.Do("GET", "a"))

	fmt.Println(a)

	defer c.Close()
}
