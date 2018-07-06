package service

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/model"
	"github.com/holywolfchan/yuncang/utils/logs"
)

//获取上下文中存在的userinfo信息
func GetContextUser(c *gin.Context) (model.UserFullInfo, error) {
	var resultuser model.UserFullInfo
	retuser, _ := c.Get("userinfo")
	ret := retuser.(string)
	if ret != "" {
		if err := json.Unmarshal([]byte(ret), &resultuser); err != nil {
			logs.Infof("unmarshal userinfo failed:%s", err)
			return resultuser, err
		}
		return resultuser, nil
	}
	return resultuser, errors.New("当前上下文中 userinfo 为空")
}
