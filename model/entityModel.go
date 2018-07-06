package model

import (
	"time"
)

type EntityStatus struct {
	Id   int64  `xorm:"pk autoincr" json:"estatus_id,omitempty"`
	Name string `json:"estatus_name,omitempty"`
}

type Factory struct {
	Id       int64     `xorm:"pk autoincr" json:"factory_id,omitempty"`
	Name     string    `xorm:"varchar(100)" json:"factory_name,omitempty"`
	Level    int       `json:"factory_level,omitempty"`
	Demo     string    `json:"factory_demo,omitempty"`
	Status   string    `json:"factory_status,omitempty"`
	Creatime time.Time `json:"factory_creatime,omitempty"`
}
type Supplier struct {
	Id       int64     `xorm:"pk autoincr" json:"supplier_id,omitempty"`
	Name     string    `xorm:"varchar(100)" json:"supplier_name,omitempty"`
	Level    int       `json:"supplier_level,omitempty"`
	Demo     string    `json:"supplier_demo,omitempty"`
	Status   string    `json:"supplier_status,omitempty"`
	Creatime time.Time `json:"supplier_creatime,omitempty"`
}

type entityService struct {
}

var EntityService = new(entityService)

//////////
func (service *entityService) GetFactory() ([]Factory, error) {
	var listarr []Factory
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *entityService) QueryFactory(id int64) ([]Factory, error) {
	var listarr []Factory
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *entityService) UpdateFactory(info *Factory) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (service *entityService) InsertFactory(info *Factory) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (service *entityService) DeleteFactory(id int64) (int64, error) {
	sql := "DELETE FROM factory WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

///////////
func (service *entityService) GetSupplier() ([]Supplier, error) {
	var listarr []Supplier
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *entityService) QuerySupplier(id int64) ([]Supplier, error) {
	var listarr []Supplier
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *entityService) UpdateSupplier(info *Supplier) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (service *entityService) InsertSupplier(info *Supplier) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (service *entityService) DeleteSupplier(id int64) (int64, error) {
	sql := "DELETE FROM supplier WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

/////////
func (service *entityService) GetEntityStatus() ([]EntityStatus, error) {
	var listarr []EntityStatus
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *entityService) QueryEntityStatus(id int64) ([]EntityStatus, error) {
	var listarr []EntityStatus
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *entityService) UpdateEntityStatus(info *EntityStatus) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (service *entityService) InsertEntityStatus(info *EntityStatus) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (service *entityService) DeleteEntityStatus(id int64) (int64, error) {
	sql := "DELETE FROM entity_status WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}
