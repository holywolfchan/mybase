package model

import (
	"fmt"

	"github.com/xormplus/xorm"
)

/* ------------------表常用操作---------------------*/
func (list *MaterialList) Get() ([]MaterialList, error) {
	var listarr []MaterialList
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialList) QueryByType(intslice []string) ([]MaterialList, error) {
	var listarr []MaterialList
	if err := DB.In("typeid", intslice).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialList) QueryByFilter(filter string) ([]MaterialList, error) {
	var listarr []MaterialList
	fstr := "%" + filter + "%"
	sql := "SELECT * FROM material_list AS l WHERE l.`name`  LIKE ? OR l.`code` LIKE ? "
	err := DB.Sql(sql, fstr, fstr).Find(&listarr)

	return listarr, err
}

type PartialStock struct {
	MaterialStockpartial `xorm:"extends"`
	Warehousename        string `json:"partial_warehousename"`
}

func (list *MaterialList) QueryPartialStock(id int64) ([]PartialStock, error) {
	var listarr []PartialStock
	sql := "SELECT p.materialid AS materialid,p.amount AS amount,w.`name` AS warehousename FROM material_stockpartial AS p LEFT JOIN material_warehouse AS w ON p.warehouseid=w.id WHERE p.materialid = ?"
	err := DB.Sql(sql, id).Find(&listarr)
	return listarr, err
}

func (list *MaterialList) Query(id int64) ([]MaterialList, error) {
	var listarr []MaterialList
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialList) Update(info *MaterialList) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (list *MaterialList) Insert(info *MaterialList) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (list *MaterialList) Delete(id int64) (int64, error) {
	sql := "DELETE FROM material_list WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

///////////////
func (list *MaterialLogin) Get() ([]MaterialLogin, error) {
	var listarr []MaterialLogin
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialLogin) Query(id int64) ([]MaterialLogin, error) {
	var listarr []MaterialLogin
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialLogin) Update(info *MaterialLogin) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (list *MaterialLogin) Insert(info *MaterialLogin) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (list *MaterialLogin) Delete(id int64) (int64, error) {
	sql := "DELETE FROM material_login WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

////////////
func (list *MaterialLogout) Get() ([]MaterialLogout, error) {
	var listarr []MaterialLogout
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialLogout) Query(id int64) ([]MaterialLogout, error) {
	var listarr []MaterialLogout
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}

type MaterialLogoutdocs struct {
	MaterialLogout `xorm:"extends"`
	Typename       string `json:"logout_typename"`
}

func (list *MaterialLogout) QueryByDocid(docid string) ([]MaterialLogoutdocs, error) {
	var listarr []MaterialLogoutdocs
	sql := "SELECT o.id AS id,o.docid AS docid,o.materialid AS materialid,o.amount AS amount,o.unitname AS unitname,o.`to` AS `to`,o.toid AS toid,o.picurl AS picurl,o.operator AS operator,o.warehouseid AS warehouseid,o.demo AS demo,o.planid AS planid,o.logtime AS logtime,o.cancel AS cancel,o.canceltime AS canceltime,l.`name` AS materialname,ty.`name` AS typename FROM material_logout AS o LEFT JOIN material_list AS l ON o.materialid=l.id LEFT JOIN material_type AS ty ON ty.id=l.typeid WHERE o.docid = ?"
	if err := DB.Sql(sql, docid).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil

}
func (list *MaterialLogout) Update(info *MaterialLogout) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (list *MaterialLogout) Insert(info *MaterialLogout) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (list *MaterialLogout) Delete(id int64) (int64, error) {
	sql := "DELETE FROM material_logout WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

/////////////
type MorderRet struct {
	MaterialOrder `xorm:"extends"`
	Loginsum      float64 `json:"morder_loginsum"`
	Materialname  string  `json:"morder_materialname"`
	Unitname      string  `json:"morder_unitname"`
}

func (list *MaterialOrder) Get(wherestr string) ([]MorderRet, error) {
	var listarr []MorderRet
	if wherestr == "" {
		wherestr = "1==1"
	}
	sql := "SELECT o.*, sum(l.amount) AS loginsum, m.`name` AS materialname,m.unitname AS unitname FROM material_order AS o LEFT JOIN material_login AS l ON o.id = l.orderid LEFT JOIN material_list AS m ON o.materialid = m.id WHERE " + wherestr + " GROUP BY o.id "
	if err := DB.Sql(sql).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialOrder) Query(id int64) ([]MaterialOrder, error) {
	var listarr []MaterialOrder
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialOrder) Update(info *MaterialOrder) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (list *MaterialOrder) Insert(info *MaterialOrder) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (list *MaterialOrder) Delete(id int64) (int64, error) {
	sql := "DELETE FROM material_order WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

//////////////
func (list *MaterialType) Get() ([]MaterialType, error) {
	var listarr []MaterialType
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialType) Query(id int64) ([]MaterialType, error) {
	var listarr []MaterialType
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *MaterialType) Update(info *MaterialType) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (list *MaterialType) Insert(info *MaterialType) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (list *MaterialType) Delete(id int64) (int64, error) {
	sql := "DELETE FROM material_type WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

///////////
func (list *Supplier) Get() ([]Supplier, error) {
	var listarr []Supplier
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *Supplier) Query(id int64) ([]Supplier, error) {
	var listarr []Supplier
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *Supplier) Update(info *Supplier) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (list *Supplier) Insert(info *Supplier) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (list *Supplier) Delete(id int64) (int64, error) {
	sql := "DELETE FROM supplier WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

////////////////////
func (service *MaterialLogoutdoc) GetMaterialLogoutdoc() ([]MaterialLogoutdoc, error) {
	var listarr []MaterialLogoutdoc
	if err := DB.Desc("creatime").Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *MaterialLogoutdoc) QueryMaterialLogoutdoc(id string) ([]MaterialLogoutdoc, error) {
	var listarr []MaterialLogoutdoc
	if err := DB.Where("docid = ?", id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *MaterialLogoutdoc) UpdateMaterialLogoutdoc(info *MaterialLogoutdoc) (cnt int64, err error) {
	cnt, err = DB.Id(info.Docid).AllCols().Update(info)
	return
}
func (service *MaterialLogoutdoc) InsertMaterialLogoutdoc(info *MaterialLogoutdoc) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (service *MaterialLogoutdoc) DeleteMaterialLogoutdoc(id string) (int64, error) {
	sql := "DELETE FROM material_logoutdoc WHERE docid = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

////////////
func (service *MaterialWarehouse) GetMaterialWarehouse() ([]MaterialWarehouse, error) {
	var listarr []MaterialWarehouse
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *MaterialWarehouse) QueryMaterialWarehouse(id int64) ([]MaterialWarehouse, error) {
	var listarr []MaterialWarehouse
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *MaterialWarehouse) UpdateMaterialWarehouse(info *MaterialWarehouse) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (service *MaterialWarehouse) InsertMaterialWarehouse(info *MaterialWarehouse) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (service *MaterialWarehouse) DeleteMaterialWarehouse(id int64) (int64, error) {
	sql := "DELETE FROM material_warehouse WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

/////////////
func (list *UnitType) Get() ([]UnitType, error) {
	var listarr []UnitType
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *UnitType) Query(id int64) ([]UnitType, error) {
	var listarr []UnitType
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (list *UnitType) Update(info *UnitType) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (list *UnitType) Insert(info *UnitType) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (list *UnitType) Delete(id int64) (int64, error) {
	sql := "DELETE FROM unit_type WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

////////

/*-------------------表事务处理及复杂业务逻辑---------------------*/
//物料入库事务
func MaterialLoginTransaction(logininfo *MaterialLogin) error {
	session := DB.NewSession()
	defer session.Close()

	err := session.Begin()
	//插入物料入库表
	_, err = session.Insert(logininfo)
	if err != nil {
		session.Rollback()
		return err
	}
	//更新物料总库存,正常应当有记录被更新，若更新记录数为0，则存在异常
	err = updateMaterial_liststock(session, logininfo.Materialid, logininfo.Amount, "+")
	if err != nil {
		session.Rollback()
		return err
	}
	//更新物料分仓库存
	err = updateMaterial_partialstock(session, logininfo.Materialid, logininfo.Warehouseid, logininfo.Amount, "+")
	if err != nil {
		session.Rollback()
		return err
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil

}

//物料出库事务
func MaterialLogoutTransaction(logoutinfo *MaterialLogout) error {
	session := DB.NewSession()
	defer session.Close()

	err := session.Begin()
	//插入物料出库表
	_, err = session.Insert(logoutinfo)
	if err != nil {
		session.Rollback()
		return err
	}
	//更新物料总库存
	err = updateMaterial_liststock(session, logoutinfo.Materialid, logoutinfo.Amount, "-")
	if err != nil {
		session.Rollback()
		return err
	}
	//更新物料分仓库存

	err = updateMaterial_partialstock(session, logoutinfo.Materialid, logoutinfo.Warehouseid, logoutinfo.Amount, "-")
	if err != nil {
		session.Rollback()
		return err
	}

	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

//冲销物料入库单
func MaterialLoginCancelTransaction(loginfo *MaterialLogin) error {
	session := DB.NewSession()
	defer session.Close()

	var cnt int64
	err := session.Begin()

	cnt, err = session.Id(loginfo.Id).Cols("cancel", "canceltime").Update(loginfo)
	if err != nil || cnt == 0 {
		session.Rollback()
		return fmt.Errorf("更新入库单状态出错:rowsAffected[%d] err[%v]", cnt, err)
	}

	err = updateMaterial_liststock(session, loginfo.Materialid, loginfo.Amount, "-")
	if err != nil {
		session.Rollback()
		return err
	}

	err = updateMaterial_partialstock(session, loginfo.Materialid, loginfo.Warehouseid, loginfo.Amount, "-")
	if err != nil {
		session.Rollback()
		return err
	}

	return nil
}

//冲销物料出库单
func MaterialLogoutCancelTransaction(logoutinfo *MaterialLogout) error {
	session := DB.NewSession()
	defer session.Close()

	var cnt int64
	err := session.Begin()

	cnt, err = session.Id(logoutinfo.Id).Cols("cancel", "canceltime").Update(logoutinfo)

	if err != nil || cnt == 0 {
		session.Rollback()
		return fmt.Errorf("更新出库单状态出错:rowsAffected[%d] err[%v]", cnt, err)
	}

	err = updateMaterial_liststock(session, logoutinfo.Materialid, logoutinfo.Amount, "+")
	if err != nil {
		session.Rollback()
		return err
	}

	err = updateMaterial_partialstock(session, logoutinfo.Materialid, logoutinfo.Warehouseid, logoutinfo.Amount, "+")
	if err != nil {
		session.Rollback()
		return err
	}

	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

//动作——更新总仓库存
func updateMaterial_liststock(session *xorm.Session, materialid int64, amount float64, mark string) error {
	sql := fmt.Sprintf("UPDATE material_list SET stockamount = stockamount %s ? WHERE id = ?", mark)
	ret, err := session.Exec(sql, amount, materialid)
	cnt, _ := ret.RowsAffected()
	if err != nil || cnt == 0 {
		return fmt.Errorf("更新物料总库存出错：rowsaffected[%d],err[%v]", cnt, err)
	}
	return nil
}

//动作——更新分仓库存
func updateMaterial_partialstock(session *xorm.Session, materialid int64, warehouseid int64, amount float64, mark string) error {
	sql := fmt.Sprintf("INSERT INTO material_stockpartial (warehouseid,materialid,amount) VALUES (?,?,?) ON DUPLICATE KEY UPDATE amount = amount %s VALUES(amount)", mark)
	_, err := session.Exec(sql, warehouseid, materialid, amount)
	if err != nil {
		return err
	}
	return nil
}
