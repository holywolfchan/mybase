package controller

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/holywolfchan/yuncang/utils/logs"

	"github.com/json-iterator/go"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/model"
)

type materialController struct {
	list      *model.MaterialList
	login     *model.MaterialLogin
	logout    *model.MaterialLogout
	order     *model.MaterialOrder
	mtype     *model.MaterialType
	warehouse *model.MaterialWarehouse
	partial   *model.MaterialStockpartial
	munit     *model.UnitType
	logoutdoc *model.MaterialLogoutdoc
}

var MaterialController = new(materialController)

func (this materialController) GetMaterialList(c *gin.Context) {
	list, err := this.list.Get()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    list,
	})
}

func (this materialController) QueryList(c *gin.Context) {
	str := c.Query("id")
	if str == "" {
		fallback(c, errors.New("id 不得为空"))
		return
	}
	id, _ := strconv.ParseInt(str, 10, 64)
	rst, err := this.list.Query(id)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    rst,
	})

}

func (this materialController) QueryPartialstockByid(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		fallback(c, err)
		return
	}
	ret, err1 := this.list.QueryPartialStock(id)
	if err1 != nil {
		fallback(c, err1)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}

type reqData struct {
	Type []string `json:"type"`
}

func (this materialController) QueryListByType(c *gin.Context) {
	var reqd reqData
	arr, _ := c.GetRawData()

	if err := jsoniter.Unmarshal(arr, &reqd); err != nil {
		logs.Errorf("controller.materialcontroller.QueryListByType:%v", err)
	}
	ret, err := this.list.QueryByType(reqd.Type)
	if err != nil {
		fallback(c, err)
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}

func (this materialController) QueryListByFilter(c *gin.Context) {
	querystr := c.Query("qstring")
	if querystr != "" {
		ret, err := this.list.QueryByFilter(querystr)
		if err != nil {
			fallback(c, err)
			return
		}
		c.JSON(200, model.SuccessMessage{
			Errcode: Error_Success,
			Data:    ret,
		})
		return
	}
	fallback(c, errors.New("qstring is null"))
}

func (this materialController) InsertList(c *gin.Context) {
	var arr model.MaterialList
	if err := getJson(c, &arr); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.list.Insert(&arr)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    fmt.Sprintf("%d records inserted", cnt),
	})

}
func (this materialController) UpdateList(c *gin.Context) {
	var info model.MaterialList
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.list.Update(&info)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}
func (this materialController) DeleteList(c *gin.Context) {
	var list model.MaterialList

	if err := getJson(c, &list); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.list.Delete(list.Id)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}

/////////////////////////MaterialOrder
func (this materialController) GetMorder(c *gin.Context) {
	key := c.Query("key")
	op := c.Query("op")
	value := c.Query("value")
	var filterstr string
	if key == "" || value == "" {
		filterstr = "1==1"
		op = ""
	}
	switch op {
	case "eq":
		filterstr = "o." + key + " = " + value
	case "gt":
		filterstr = "o." + key + " > " + value
	case "lt":
		filterstr = "o." + key + " < " + value
	}
	list, err := this.order.Get(filterstr)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    list,
	})
}

func (this materialController) QueryMorder(c *gin.Context) {
	str := c.Query("id")
	if str == "" {
		fallback(c, errors.New("id 不得为空"))
		return
	}
	id, _ := strconv.ParseInt(str, 10, 64)
	rst, err := this.order.Query(id)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    rst,
	})

}

func (this materialController) InsertMorder(c *gin.Context) {
	var arr model.MaterialOrder
	if err := getJson(c, &arr); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.order.Insert(&arr)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    fmt.Sprintf("%d records inserted", cnt),
	})

}
func (this materialController) UpdateMorder(c *gin.Context) {
	var info model.MaterialOrder
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.order.Update(&info)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}
func (this materialController) DeleteMorder(c *gin.Context) {
	var list model.MaterialOrder

	if err := getJson(c, &list); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.order.Delete(list.Id)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}

////////////////////
func (this materialController) GetMtype(c *gin.Context) {
	list, err := this.mtype.Get()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    list,
	})
}

func (this materialController) QueryMtype(c *gin.Context) {
	str := c.Query("id")
	if str == "" {
		fallback(c, errors.New("id 不得为空"))
		return
	}
	id, _ := strconv.ParseInt(str, 10, 64)
	rst, err := this.mtype.Query(id)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    rst,
	})

}

func (this materialController) InsertMtype(c *gin.Context) {
	var arr model.MaterialType
	if err := getJson(c, &arr); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.mtype.Insert(&arr)
	if err != nil {
		fmt.Printf("ctrerror:%v\n", err)
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    fmt.Sprintf("%d records inserted", cnt),
	})

}
func (this materialController) UpdateMtype(c *gin.Context) {
	var info model.MaterialType
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.mtype.Update(&info)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}
func (this materialController) DeleteMtype(c *gin.Context) {
	var list model.MaterialType

	if err := getJson(c, &list); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.mtype.Delete(list.Id)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}

/////////////////
func (this materialController) GetMunit(c *gin.Context) {
	list, err := this.munit.Get()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    list,
	})
}

func (this materialController) QueryMunit(c *gin.Context) {
	str := c.Query("id")
	if str == "" {
		fallback(c, errors.New("id 不得为空"))
		return
	}
	id, _ := strconv.ParseInt(str, 10, 64)
	rst, err := this.munit.Query(id)
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    rst,
	})

}

func (this materialController) InsertMunit(c *gin.Context) {
	var arr model.UnitType
	if err := getJson(c, &arr); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.munit.Insert(&arr)
	if err != nil {
		fmt.Printf("ctrerror:%v\n", err)
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    fmt.Sprintf("%d records inserted", cnt),
	})

}
func (this materialController) UpdateMunit(c *gin.Context) {
	var info model.UnitType
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.munit.Update(&info)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}
func (this materialController) DeleteMunit(c *gin.Context) {
	var list model.UnitType

	if err := getJson(c, &list); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.munit.Delete(list.Id)
	if err != nil || cnt == 0 {
		fallback(c, fmt.Errorf("%d rows affected:%v", cnt, err))
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: fmt.Sprintf("%d rows affected", cnt),
	})
}

//////////////////
func (this materialController) GetMaterialLogoutdoc(c *gin.Context) {
	ret, err := this.logoutdoc.GetMaterialLogoutdoc()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this materialController) QueryMaterialLogoutdoc(c *gin.Context) {
	id := c.Query("id")

	ret, err2 := this.logoutdoc.QueryMaterialLogoutdoc(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this materialController) QueryMaterialLogoutByDocid(c *gin.Context) {
	id := c.Query("docid")

	ret, err2 := this.logout.QueryByDocid(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this materialController) InsertMaterialLogoutdoc(c *gin.Context) {
	var info model.MaterialLogoutdoc
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.logoutdoc.InsertMaterialLogoutdoc(&info)
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
func (this materialController) UpdateMaterialLogoutdoc(c *gin.Context) {
	var info model.MaterialLogoutdoc
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.logoutdoc.UpdateMaterialLogoutdoc(&info)
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
func (this materialController) DeleteMaterialLogoutdoc(c *gin.Context) {
	var info model.MaterialLogoutdoc
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.logoutdoc.DeleteMaterialLogoutdoc(info.Docid)
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

//物料入库
func (this materialController) MaterialLoginCtr(c *gin.Context) {
	var logininfo model.MaterialLogin
	if err := getJson(c, &logininfo); err != nil {
		fallback(c, err)
		return
	}
	if err := model.MaterialLoginTransaction(&logininfo); err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: "物料入库成功",
	})

}

//物料出库
func (this materialController) MaterialLogoutCtr(c *gin.Context) {
	var logoutinfo model.MaterialLogout
	if err := getJson(c, &logoutinfo); err != nil {
		fallback(c, err)
		return
	}

	if err := model.MaterialLogoutTransaction(&logoutinfo); err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: "物料出库成功",
	})
}

//冲销物料入库单
func (this materialController) MaterialLoginCancelCtr(c *gin.Context) {
	var logininfo model.MaterialLogin
	if err := getJson(c, &logininfo); err != nil {
		fallback(c, err)
		return
	}

	if err := model.MaterialLoginCancelTransaction(&logininfo); err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: "冲销物料入库单成功",
	})
}

//冲销物料出库单
func (this materialController) MaterialLogoutCancelCtr(c *gin.Context) {
	var logoutinfo model.MaterialLogout
	if err := getJson(c, &logoutinfo); err != nil {
		fallback(c, err)
		return
	}

	if err := model.MaterialLogoutCancelTransaction(&logoutinfo); err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Message: "冲销物料出库单成功",
	})
}

//出库单批量操作
func (this materialController) MaterialLogoutBatch(c *gin.Context) {
	var infoarr []model.MaterialLogout
	if err := getJson(c, &infoarr); err != nil {
		fallback(c, err)
		return
	}
	var docid = time.Now().Format("20060102150405")
	logs.Infof("出库单%s事务开始>>>", docid)
	logs.Infof("Array:%v", infoarr)
	for _, info := range infoarr {
		info.Docid = docid
		logs.Infof("Item:%v", info)
		if err := model.MaterialLogoutTransaction(&info); err != nil {
			fallback(c, err)
			return
		}
	}
	logs.Infof("出库单%s事务结束<<<", docid)

}

////////分仓操作
func (this materialController) GetMaterialWarehouse(c *gin.Context) {
	ret, err := this.warehouse.GetMaterialWarehouse()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this materialController) QueryMaterialWarehouse(c *gin.Context) {
	id, err1 := getID(c)
	if err1 != nil {
		fallback(c, err1)
		return
	}
	ret, err2 := this.warehouse.QueryMaterialWarehouse(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this materialController) InsertMaterialWarehouse(c *gin.Context) {
	var info model.MaterialWarehouse
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.warehouse.InsertMaterialWarehouse(&info)
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
func (this materialController) UpdateMaterialWarehouse(c *gin.Context) {
	var info model.MaterialWarehouse
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.warehouse.UpdateMaterialWarehouse(&info)
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
func (this materialController) DeleteMaterialWarehouse(c *gin.Context) {
	var info model.MaterialWarehouse
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := this.warehouse.DeleteMaterialWarehouse(info.Id)
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
