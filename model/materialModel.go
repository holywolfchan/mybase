package model

import (
	"time"
)

type MaterialType struct {
	Id       int64  `xorm:"pk autoincr" json:"type_id,string,omitempty"`
	Name     string `xorm:"varchar(50) index" json:"type_name,omitempty"`
	Parentid int64  `xorm:"index" json:"type_parentid,string,omitempty"`
	Sortid   int    `json:"type_sortid,string,omitempty"`
}

type UnitType struct {
	Id   int64  `xorm:"pk autoincr" json:"unit_id,string"`
	Name string `xorm:"varchar(20)" json:"unit_name"`
}

type MaterialList struct {
	Id          int64     `xorm:"pk autoincr" json:"material_id,string"`
	Name        string    `xorm:"varchar(100)" json:"material_name"`
	Typeid      int64     `xorm:"index" json:"material_typeid"`
	Stockamount float64   `xorm:"default(0)" json:"material_stockamount"`
	Unitname    string    `json:"material_unitname"`
	Code        string    `json:"material_code"`
	Demo        string    `json:"material_demo,omitempty"`
	Level       int       `json:"material_level"`
	Upperlimit  float64   `json:"material_upperlimit,omitempty"`
	Lowerlimit  float64   `json:"material_lowerlimit,omitempty"`
	Creatime    time.Time `xorm:"created" json:"material_creatime"`
}

type MaterialLogin struct {
	Id          int64     `xorm:"pk autoincr" json:"login_id,string"`
	Materialid  int64     `xorm:"index" json:"login_materialid"`
	Amount      float64   `xorm:"default(0)" json:"login_amount"`
	Unitname    string    `json:"login_unitname"`
	From        string    `json:"login_from"`
	Fromid      int64     `xorm:"index" json:"login_fromid"`
	Picurl      string    `xorm:"text" json:"login_picurl"`
	Operator    string    `json:"login_operator"`
	Orderid     int64     `xorm:"index" json:"login_orderid"` //对应采购订单id
	Warehouseid int64     `xorm:"index" json:"login_warehouseid"`
	Demo        string    `json:"login_demo"`
	Logtime     time.Time `xorm:"created" json:"login_time"`
	Cancel      bool      `json:"login_cancel"`
	Canceltime  time.Time `json:"login_canceltime"`
}

type MaterialLogout struct {
	Id           int64     `xorm:"pk autoincr" json:"logout_id,string"`
	Docid        string    `xorm:"index" json:"logout_docid"`
	Materialid   int64     `xorm:"index" json:"logout_materialid"`
	Materialname string    `xorm:"<-" json:"logout_materialname"`
	Amount       float64   `xorm:"default(0)" json:"logout_amount"`
	Unitname     string    `json:"logout_unitname"`
	To           string    `json:"logout_to"`
	Toid         int64     `xorm:"index" json:"logout_toid"`
	Picurl       string    `xorm:"text" json:"logout_picurl"`
	Operator     string    `json:"logout_operator"`
	Warehouseid  int64     `xorm:"index" json:"logout_warehouseid"`
	Demo         string    `json:"logout_demo"`
	Planid       int64     `xorm:"index" json:"logout_planid"` //对应生产订单id
	Logtime      time.Time `xorm:"created" json:"logout_time"`
	Cancel       bool      `json:"logout_cancel"`
	Canceltime   time.Time `json:"logout_canceltime"`
}

type MaterialOrder struct {
	Id         int64     `xorm:"pk autoincr" json:"morder_id,string"`
	Materialid int64     `xorm:"index" json:"morder_materialid"`
	Amount     float64   `xorm:"default(0)" json:"morder_amount"`
	From       string    `json:"morder_from"`
	Fromid     int64     `xorm:"index" json:"morder_fromid"`
	Price      float64   `json:"morder_price"`
	Picurl     string    `json:"morder_picurl"`
	Operator   string    `json:"morder_operator"`
	Demo       string    `json:"morder_demo"`
	Logtime    time.Time `xorm:"created" json:"morder_buytime"`
	Arrivetime time.Time `json:"morder_arrivetime"`
	Finished   bool      `json:"morder_isfinished"`
}

type MaterialWarehouse struct {
	Id   int64  `xorm:"pk autoincr" json:"warehouse_id"`
	Name string `xorm:"varchar(50)" json:"warehouse_name"`
	Demo string `json:"warehouse_demo"`
}

type MaterialStockpartial struct {
	Warehouseid int64   `xorm:"pk" json:"partial_warehouseid,string"`
	Materialid  int64   `xorm:"pk" json:"partial_materialid"`
	Amount      float64 `json:"partial_amount"`
}

type MaterialLogoutdoc struct {
	Docid       string    `xorm:"pk" json:"logout_docid"`
	To          string    `json:"logout_to"`
	Toid        int64     `json:"logout_toid"`
	Operator    string    `json:"logout_operator"`
	Picurl      string    `json:"logout_picurl"`
	Planid      int64     `json:"logout_planid"`
	Warehouseid int64     `json:"logout_warehouseid"`
	Creatime    time.Time `xorm:"created" json:"logout_doctime"`
	Printed     bool      `json:"logout_docprinted"`
	Cancel      bool      `json:"logout_cancel"`
	Canceltime  time.Time `json:"logout_canceltime"`
	Demo        string    `json:"logout_docdemo"`
}
