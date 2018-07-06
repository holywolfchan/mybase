package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/holywolfchan/yuncang/utils"

	"github.com/holywolfchan/yuncang/utils/logs"
	"github.com/holywolfchan/yuncang/utils/redisgo"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/db"
	"github.com/holywolfchan/yuncang/model"
)

var (
	DB    = db.Engine
	Redis = redisgo.RedisEngine
)

type loginuser struct {
	Username       string `json:"username"`
	RoleId         int    `json:"role_id"`
	RoleName       string `json:"rolename"`
	Avatar         string `json:"avatar"`
	OpenId         string `json:"openid"`
	UnionId        string `json:"unionid"`
	DepartmentId   int    `json:"department_id"`
	DepartmentName string `json:"departmentname"`
	DomainId       int    `json:"domain_id"`
	DomainName     string `json:"domainname"`
}

func Login(c *gin.Context) {
	var user model.User
	var retuser loginuser
	var userinfo model.UserFullInfo
	if err := c.BindJSON(&user); err != nil {
		logs.Fatal("数据结构解析出错")
		c.JSON(500, gin.H{
			"errmsg": "cannot resolve the data,check it",
		})
		return
	}
	if strings.TrimSpace(user.Passport) == "" || strings.TrimSpace(user.Password) == "" {
		c.JSON(http.StatusBadRequest, model.ErrMessage{
			Errcode: Error_loginfailed,
			Errmsg:  "账号密码均不得为空",
		})
		return
	}
	if has, _ := DB.Get(&user); has {
		key, err := utils.GetUserUID(&user)
		if err != nil {
			c.JSON(500, model.ErrMessage{
				Errcode: Error_loginfailed,
				Errmsg:  fmt.Sprint(err),
			})
			return
		}
		if err := userinfo.GetUserInfoByLogin(user.Passport, user.Password); err != nil {
			logs.Fatal(err)
			c.JSON(500, model.ErrMessage{
				Errcode: Error_Failed,
				Errmsg:  fmt.Sprint(err),
			})
			return
		}
		userinfojson, _ := json.Marshal(userinfo)
		if err := Redis.Set(key, userinfojson, time.Hour*20).Err(); err != nil {
			logs.Debug("redis operate failed:%v\n", err)
		}
		logs.Accessf("User %s logged in:%v At %v", user.Passport, key, time.Now().Format("2006-01-02 15:04:05"))
		if err := json.Unmarshal(userinfojson, &retuser); err != nil {
			c.JSON(500, model.ErrMessage{
				Errcode: Error_Failed,
				Errmsg:  fmt.Sprint(err),
			})
			return
		}
		c.JSON(200, model.SuccessMessage{
			Errcode: Error_Success,
			Data:    retuser,
			Token:   key,
		})
		return
	}
	c.JSON(200, model.ErrMessage{
		Errcode: Error_loginfailed,
		Errmsg:  "账号密码错误",
	})
}

func GetAllUser(c *gin.Context) {
	var user = make([]model.UserFullInfo, 0)
	sql := "SELECT u.id AS id, u.user_name AS user_name, u.passport AS passport, u.avatar AS avatar, u.weixin_open_id AS weixin_open_id, u.weixin_union_id AS weixin_union_id, u.createtime AS createtime, u.role_id AS role_id, u.department_id AS department_id, u.domain_id AS domain_id, r.role_name AS role_name, d.domain_name AS domain_name FROM `user` AS u LEFT JOIN role AS r ON u.role_id = r.id LEFT JOIN domain AS d ON u.domain_id = d.id"
	if err := DB.Sql(sql).Find(&user); err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    user,
	})
}
