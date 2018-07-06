package model

import (
	"github.com/holywolfchan/yuncang/db"
	"github.com/holywolfchan/yuncang/utils/logs"
)

var dropAndCreate bool = false
var DB = db.Engine
var TableStruct = []interface{}{
	new(User),
	new(Role),
	new(Domain),
	new(MaterialList),
	new(MaterialLogin),
	new(MaterialLogout),
	new(MaterialOrder),
	new(MaterialType),
	new(MaterialWarehouse),
	new(EntityStatus),
	new(UnitType),
	new(Factory),
	new(Supplier),
	new(MaterialStockpartial),
	new(MaterialLogoutdoc),
}

func init() {
	logs.Info("同步表结构开始...")
	for _, v := range TableStruct {
		Migrate(v)
	}
	logs.Info("同步表结构结束！\n")
}

func Migrate(v interface{}) {
	if has, err := DB.IsTableExist(v); err == nil && has && dropAndCreate {
		err := DB.DropTables(v)
		if err != nil {
			logs.Errorf("删除表出错:表%T,%v", v, err)
			return
		}
		logs.Infof("删除表成功:表%T", v)
	}
	err := DB.Sync2(v)
	if err != nil {
		logs.Errorf("同步表结构出错:表[%T],%v", v, err)
		return
	}
	logs.Infof("同步表结构成功:表%T", v)
}
