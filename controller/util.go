package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/holywolfchan/yuncang/utils/logs"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/model"
)

func fallback(c *gin.Context, err error) {
	c.JSON(200, model.ErrMessage{
		Errcode: Error_Failed,
		Errmsg:  fmt.Sprint(err),
	})
}

func getJson(c *gin.Context, v interface{}) error {
	data, err1 := c.GetRawData()
	logs.Infof("data:%s", data)
	if err1 != nil {
		return err1
	}
	if err := gjson.Unmarshal(data, v); err != nil {
		logs.Errorf("util.go/getjson:ERROR[%v]", err)
		return err
	}
	logs.Infof("v:%v", v)
	return nil
}

func getID(c *gin.Context) (int64, error) {
	str := c.Query("id")
	if str == "" {

		return 0, errors.New("id 不得为空")
	}
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return id, err
	}
	return id, nil
}
