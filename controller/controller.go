package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

const keyUserRequestsPrefix = "requests."

type Controller struct {
	r Redis
}

func (c Controller) key(user string) string {
	return fmt.Sprintf("%s%x", keyUserRequestsPrefix, md5.Sum([]byte(user)))
}

func (c Controller) AcceptedRequest(user string, limit int) (int, bool) {
	key := c.key(user)

	if _, err := c.r.Get(key); err == redis.Nil {
		err := c.r.Set(key, "0", time.Second * time.Duration(limit))
		if err != nil {
			log.Println(err)
			return 0, false
		}
	}

	if err := c.r.Inc(key); err != nil {
		log.Println(err)
		return 0, false
	}

	requests, err := c.r.Get(key)
	if err != nil {
		log.Println(err)
		return 0,false
	}
	requestsNum, err := strconv.Atoi(requests)
	if err != nil {
		log.Println(err)
		return 0, false
	}

	if requestsNum > limit {
		return requestsNum, false
	}

	return requestsNum, true

}

var controller = &Controller{}

func Instance() *Controller {
	return controller
}

func SetRedis(redis Redis) {
	controller.r = redis
}
