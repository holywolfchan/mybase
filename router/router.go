package router

import (
	"github.com/gin-gonic/gin"
	"github.com/holywolfchan/yuncang/controller"
)

func NonAuthRouters(r *gin.RouterGroup) {
	r.POST("/login", controller.Login)
	r.POST("/register", controller.UserRegister)
	r.GET("/passportcheck", controller.PassportCheck)

	r.GET("/user/get", controller.UserController.GetUser)
	r.GET("/user/query", controller.UserController.QueryUser)
	r.POST("/user/insert", controller.UserController.InsertUser)
	r.POST("/user/update", controller.UserController.UpdateUser)
	r.POST("/user/delete", controller.UserController.DeleteUser)

	r.GET("/role/get", controller.UserController.GetRole)
	r.GET("/role/query", controller.UserController.QueryRole)
	r.POST("/role/insert", controller.UserController.InsertRole)
	r.POST("/role/update", controller.UserController.UpdateRole)
	r.POST("/role/delete", controller.UserController.DeleteRole)

	r.POST("/material/login/input", controller.MaterialController.MaterialLoginCtr)
	r.POST("/material/logout/output", controller.MaterialController.MaterialLogoutCtr)
	r.POST("/material/login/cancel", controller.MaterialController.MaterialLoginCancelCtr)
	r.POST("/material/logout/cancel", controller.MaterialController.MaterialLogoutCancelCtr)
	r.POST("/material/logout/batch", controller.MaterialController.MaterialLogoutBatch)
	r.GET("/material/logout/querybydocid", controller.MaterialController.QueryMaterialLogoutByDocid)

	r.GET("/material/logoutdoc/get", controller.MaterialController.GetMaterialLogoutdoc)
	r.GET("/material/logoutdoc/query", controller.MaterialController.QueryMaterialLogoutdoc)
	r.POST("/material/logoutdoc/insert", controller.MaterialController.InsertMaterialLogoutdoc)
	r.POST("/material/logoutdoc/update", controller.MaterialController.UpdateMaterialLogoutdoc)
	r.POST("/material/logoutdoc/delete", controller.MaterialController.DeleteMaterialLogoutdoc)

	r.GET("/material/list/get", controller.MaterialController.GetMaterialList)
	r.GET("/material/list/query", controller.MaterialController.QueryList)
	r.GET("/material/list/filter", controller.MaterialController.QueryListByFilter)
	r.GET("/material/list/querypartialstock",controller.MaterialController.QueryPartialstockByid)
	r.POST("/material/list/querybytype", controller.MaterialController.QueryListByType)
	r.POST("/material/list/insert", controller.MaterialController.InsertList)
	r.POST("/material/list/update", controller.MaterialController.UpdateList)
	r.POST("/material/list/delete", controller.MaterialController.DeleteList)

	r.GET("/material/mtype/get", controller.MaterialController.GetMtype)
	r.GET("/material/mtype/query", controller.MaterialController.QueryMtype)
	r.POST("/material/mtype/insert", controller.MaterialController.InsertMtype)
	r.POST("/material/mtype/update", controller.MaterialController.UpdateMtype)
	r.POST("/material/mtype/delete", controller.MaterialController.DeleteMtype)

	r.GET("/material/morder/get", controller.MaterialController.GetMorder)
	r.GET("/material/morder/query", controller.MaterialController.QueryMorder)
	r.POST("/material/morder/insert", controller.MaterialController.InsertMorder)
	r.POST("/material/morder/update", controller.MaterialController.UpdateMorder)
	r.POST("/material/morder/delete", controller.MaterialController.DeleteMorder)

	r.GET("/material/unit/get", controller.MaterialController.GetMunit)
	r.GET("/material/unit/query", controller.MaterialController.QueryMunit)
	r.POST("/material/unit/insert", controller.MaterialController.InsertMunit)
	r.POST("/material/unit/update", controller.MaterialController.UpdateMunit)
	r.POST("/material/unit/delete", controller.MaterialController.DeleteMunit)

	r.GET("/material/factory/get", controller.EntityController.GetFactory)
	r.GET("/material/factory/query", controller.EntityController.QueryFactory)
	r.POST("/material/factory/insert", controller.EntityController.InsertFactory)
	r.POST("/material/factory/update", controller.EntityController.UpdateFactory)
	r.POST("/material/factory/delete", controller.EntityController.DeleteFactory)

	r.GET("/material/supplier/get", controller.EntityController.GetSupplier)
	r.GET("/material/supplier/query", controller.EntityController.QuerySupplier)
	r.POST("/material/supplier/insert", controller.EntityController.InsertSupplier)
	r.POST("/material/supplier/update", controller.EntityController.UpdateSupplier)
	r.POST("/material/supplier/delete", controller.EntityController.DeleteSupplier)

	r.GET("/material/warehouse/get", controller.MaterialController.GetMaterialWarehouse)
	r.GET("/material/warehouse/query", controller.MaterialController.QueryMaterialWarehouse)
	r.POST("/material/warehouse/insert", controller.MaterialController.InsertMaterialWarehouse)
	r.POST("/material/warehouse/update", controller.MaterialController.UpdateMaterialWarehouse)
	r.POST("/material/warehouse/delete", controller.MaterialController.DeleteMaterialWarehouse)

}
func AuthRouters(r *gin.RouterGroup) {
	r.GET("/getalluser", controller.GetAllUser)

}
