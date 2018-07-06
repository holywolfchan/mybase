package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/model"
)

type entityController struct {
}

var EntityController = new(entityController)

func (this entityController) GetFactory(c *gin.Context) {
	ret, err := model.EntityService.GetFactory()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this entityController) QueryFactory(c *gin.Context) {
	id, err1 := getID(c)
	if err1 != nil {
		fallback(c, err1)
		return
	}
	ret, err2 := model.EntityService.QueryFactory(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}

func (this entityController) InsertFactory(c *gin.Context) {
	var info model.Factory
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.InsertFactory(&info)
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
func (this entityController) UpdateFactory(c *gin.Context) {
	var info model.Factory
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.UpdateFactory(&info)
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
func (this entityController) DeleteFactory(c *gin.Context) {
	var info model.Factory
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.DeleteFactory(info.Id)
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

//////////////////
func (this entityController) GetSupplier(c *gin.Context) {
	ret, err := model.EntityService.GetSupplier()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this entityController) QuerySupplier(c *gin.Context) {
	id, err1 := getID(c)
	if err1 != nil {
		fallback(c, err1)
		return
	}
	ret, err2 := model.EntityService.QuerySupplier(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this entityController) InsertSupplier(c *gin.Context) {
	var info model.Supplier
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.InsertSupplier(&info)
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
func (this entityController) UpdateSupplier(c *gin.Context) {
	var info model.Supplier
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.UpdateSupplier(&info)
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
func (this entityController) DeleteSupplier(c *gin.Context) {
	var info model.Supplier
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.DeleteSupplier(info.Id)
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

/////////////////
func (this entityController) GetEntityStatus(c *gin.Context) {
	ret, err := model.EntityService.GetEntityStatus()
	if err != nil {
		fallback(c, err)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this entityController) QueryEntityStatus(c *gin.Context) {
	id, err1 := getID(c)
	if err1 != nil {
		fallback(c, err1)
		return
	}
	ret, err2 := model.EntityService.QueryEntityStatus(id)
	if err2 != nil {
		fallback(c, err2)
		return
	}
	c.JSON(200, model.SuccessMessage{
		Errcode: Error_Success,
		Data:    ret,
	})
}
func (this entityController) InsertEntityStatus(c *gin.Context) {
	var info model.EntityStatus
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.InsertEntityStatus(&info)
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
func (this entityController) UpdateEntityStatus(c *gin.Context) {
	var info model.EntityStatus
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.UpdateEntityStatus(&info)
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
func (this entityController) DeleteEntityStatus(c *gin.Context) {
	var info model.EntityStatus
	if err := getJson(c, &info); err != nil {
		fallback(c, err)
		return
	}
	cnt, err := model.EntityService.DeleteEntityStatus(info.Id)
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
