package api

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang/controller"
	"log"
	"net/http"
	"strconv"
	"time"
)

const defaultLimit = 10

const cookieLimiterName = "user-limiter"

func limiter(c *gin.Context) {

	user, err := c.Request.Cookie(cookieLimiterName)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotAcceptable)
		c.Abort()
		return
	}

	requests, accepted := controller.Instance().AcceptedRequest(user.Value, defaultLimit)
	if accepted == false {
		c.Status(http.StatusTooManyRequests)
		c.Abort()
	}

	c.Header("X-RateLimit-Limit", strconv.Itoa(defaultLimit))
	c.Header("X-RateLimit-Remaining", strconv.Itoa(10-requests))

}

func router(publicPath string) http.Handler {

	router := gin.Default()
	router.Use(func(context *gin.Context) {
		_, err := context.Request.Cookie(cookieLimiterName)
		if err == http.ErrNoCookie {
			context.SetCookie(cookieLimiterName, fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))), 0, "", "", false, false)
		}
	})
	router.Use(static.Serve("/", static.LocalFile(publicPath, true)))

	api := router.Group("/api")
	api.Use(limiter)

	api.GET("/ping", handlerPing)

	return router
}

func handlerPing(c *gin.Context) {
	c.AsciiJSON(http.StatusOK, "PONG")
}
