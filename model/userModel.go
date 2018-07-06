package model

import (
	"fmt"
	"time"
)

type User struct {
	Id            int64     `xorm:"pk autoincr" json:"user_id"`
	UserName      string    `xorm:"varchar(20) index" json:"username"`
	Passport      string    `xorm:"varchar(50) index" json:"passport,omitempty"`
	Password      string    `json:"password" json:"password,omitempty"`
	Avatar        string    `json:"avatar"`
	WeixinOpenId  string    `json:"openid"`
	WeixinUnionId string    `json:"unionid"`
	Createtime    time.Time `xorm:"created" json:"create_time"`
	RoleId        int       `xorm:"index" json:"role_id"`       //角色id
	DepartmentId  int       `xorm:"index" json:"department_id"` //部门id
	DomainId      int       `xorm:"index" json:"domain_id"`     //所属域ID
}

type Domain struct {
	Id         int64  `xorm:"pk autoincr" json:"domain_id"`
	DomainName string `xorm:"varchar(50)" json:"domainname"`
}
type Role struct {
	Id       int64  `xorm:"pk autoincr" json:"role_id"`
	RoleName string `xorm:"varchar(50)" json:"rolename"`
}

type UserFullInfo struct {
	User           `xorm:"extends"`
	RoleName       string `json:"rolename"`
	DomainName     string `json:"domainname,omitempty"`
	DepartmentName string `json:"DepartmentName,omitempty"`
}

////////////////model slices//////////
type userService struct {
}

//UserFullInfoSet is a slice of UserFullInfo
var UserService = new(userService)

////////////
func (userinfo *UserFullInfo) GetUserInfoByLogin(passport, password string) error {
	var userinfoslice []UserFullInfo
	sql := "SELECT u.id AS id, u.user_name AS user_name, u.passport AS passport, u.avatar AS avatar, u.weixin_open_id AS weixin_open_id, u.weixin_union_id AS weixin_union_id, u.createtime AS createtime, u.role_id AS role_id, u.department_id AS department_id, u.domain_id AS domain_id, r.role_name AS role_name, d.domain_name AS domain_name FROM `user` AS u LEFT JOIN role AS r ON u.role_id = r.id LEFT JOIN domain AS d ON u.domain_id = d.id where u.passport = ? and u.password = ?"
	if err := DB.Sql(sql, passport, password).Find(&userinfoslice); err != nil {
		return err
	}
	fmt.Printf("have:%v\n", userinfoslice)
	ret := userinfoslice[0]
	*userinfo = ret
	fmt.Printf("have userinfo:%v\n", userinfo)
	return nil

}

func (service *userService) GetUser() ([]User, error) {
	var listarr []User
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *userService) QueryUser(id int64) ([]User, error) {
	var listarr []User
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *userService) UpdateUser(info *User) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).Cols("role_id").Update(info)
	return
}
func (service *userService) InsertUser(info *User) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (service *userService) DeleteUser(id int64) (int64, error) {
	sql := "DELETE FROM user WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}

//////////////////
func (service *userService) GetRole() ([]Role, error) {
	var listarr []Role
	if err := DB.Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *userService) QueryRole(id int64) ([]Role, error) {
	var listarr []Role
	if err := DB.Id(id).Find(&listarr); err != nil {
		return nil, err
	}
	return listarr, nil
}
func (service *userService) UpdateRole(info *Role) (cnt int64, err error) {
	cnt, err = DB.Id(info.Id).AllCols().Update(info)
	return
}
func (service *userService) InsertRole(info *Role) (cnt int64, err error) {
	cnt, err = DB.Insert(info)
	return
}
func (service *userService) DeleteRole(id int64) (int64, error) {
	sql := "DELETE FROM role WHERE id = ?"
	ret, err := DB.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected()
}
