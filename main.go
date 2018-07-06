package main

import (
	"time"

	"github.com/holywolfchan/yuncang/utils"

	"github.com/holywolfchan/yuncang/utils/auth"

	"github.com/holywolfchan/yuncang/router"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/utils/logs"
)

func main() {
	r := gin.New()
	r.Use(logs.Ginlog(time.RFC3339Nano, false))
	r.Use(gin.Recovery())
	r.Use(utils.Cors())
	nonAuthRouters := r.Group("/")
	authRouters := r.Group("/v")
	authRouters.Use(auth.Authenticate())
	authRouters.Use(auth.Authorize())
	router.NonAuthRouters(nonAuthRouters)
	router.AuthRouters(authRouters)
	r.Run(":8083")

}
