package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/utils/redisgo"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"errmsg": err.Error(),
			})
			return
		}
		fmt.Printf("token is:%s\n", token)
		if result := redisgo.RedisEngine.Get(token); result.Err() == nil {
			if s := result.Val(); s != "" {
				fmt.Printf("userinfo authenticate:%s\n", s)
				c.Set("userinfo", s)
				c.Next()
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnavailableForLegalReasons, gin.H{
				"errmsg": "token Invalid or expired",
			})
		}

	}
}

func GetToken(c *gin.Context) (s string, err error) {
	if token := c.GetHeader("Authorization"); token == "" {
		token = c.DefaultQuery("token", "none")
		if token == "none" {
			return "", fmt.Errorf("%s", "token required")
		}
		return token, nil
	} else {

		return token, nil
	}
}
