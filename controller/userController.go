package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/model"
)

type userController struct {
}

var UserController = new(userController)

func (this userController) GetUser(c *gin.Context) {
	ret, err := model.UserService.GetUser()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this userController) QueryUser(c *gin.Context) {
	id, err1 := getID(c)
	if err1 != nil {
		fallback(c, err1)
		return
	}
	ret, err2 := model.UserService.QueryUser(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this userController) InsertUser(c *gin.Context) {
	var info model.User
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.UserService.InsertUser(&info)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data: gin.H{
			"rowaffected": cnt,
		},
	})
}
func (this userController) UpdateUser(c *gin.Context) {
	var info model.User
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.UserService.UpdateUser(&info)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data: gin.H{
			"count": cnt,
		},
	})
}
func (this userController) DeleteUser(c *gin.Context) {
	var info model.User
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.UserService.DeleteUser(info.Id)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data: gin.H{
			"count": cnt,
		},
	})
}

////////////////
func (this userController) GetRole(c *gin.Context) {
	ret, err := model.UserService.GetRole()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this userController) QueryRole(c *gin.Context) {
	id, err1 := getID(c)
	if err1 != nil {
		fallback(c, err1)
		return
	}
	ret, err2 := model.UserService.QueryRole(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this userController) InsertRole(c *gin.Context) {
	var info model.Role
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.UserService.InsertRole(&info)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data: gin.H{
			"rowaffected": cnt,
		},
	})
}
func (this userController) UpdateRole(c *gin.Context) {
	var info model.Role
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.UserService.UpdateRole(&info)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data: gin.H{
			"count": cnt,
		},
	})
}
func (this userController) DeleteRole(c *gin.Context) {
	var info model.Role
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.UserService.DeleteRole(info.Id)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data: gin.H{
			"count": cnt,
		},
	})
}
