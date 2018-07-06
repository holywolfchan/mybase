package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/model"
)

func PassportCheck(c *gin.Context) {
	var user model.User
	passport := c.Query("passport")
	fmt.Printf("passport is:%s\n", passport)
	if passport != "" {
		user.Passport = passport
		if has, _ := DB.Exist(&user); has {
			c.JSON(200, model.ErrMessage{
				Errcode: Error_registerPassportExist,
				Errmsg:  "账号已存在",
			})
		} else {
			c.JSON(200, model.SuccessMessage{
				Errcode: Error_Success,
			})
		}
	} else {
		c.JSON(200, model.ErrMessage{
			Errcode: Error_Failed,
			Errmsg:  "账号不得为空",
		})
	}

}

func UserRegister(c *gin.Context) {
	var user model.User
	userinfo, _ := c.GetRawData()
	/* 	if err := c.BindJSON(&buser); err == nil {
		fmt.Printf("bindjson is :%v\n", buser)
	} */
	if err := json.Unmarshal(userinfo, &user); err == nil {
		fmt.Printf("jsonunmarshal is :%v\n", user)
	}
	if user.Passport == "" || user.Password == "" {
		c.JSON(200, model.ErrMessage{
			Errcode: Error_registerfailed,
			Errmsg:  "账号密码不得为空",
		})
		return
	}
	user.RoleId = 0
	user.DomainId = 1
	if _, err := DB.InsertOne(user); err == nil {
		c.JSON(200, model.SuccessMessage{
			Errcode: Error_Success,
		})
		return
	} else {
		c.JSON(200, model.ErrMessage{
			Errcode: Error_registerfailed,
			Errmsg:  fmt.Sprintf("注册失败：%s", err),
		})
	}

}
