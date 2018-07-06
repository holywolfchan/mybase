package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/holywolfchan/yuncang/utils/logs"
	"github.com/xormplus/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	var f *os.File
	
	Engine, err = xorm.NewEngine("mysql", "root:***@tcp(localhost:3306)/yuncang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logs.Panic(err)
	}
	logs.Infof("创建engine %v", Engine)
	Engine.ShowSQL(true)
	filename := fmt.Sprintf("./temp_log/sql/sql_%s.log", time.Now().Local().Format("2006_01_02"))
	f, err = os.OpenFile(filename, os.O_APPEND, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(filename)
		if err != nil {
			logs.Errorf("创建SQL日志出错:%v\n", err)
			return
		}
	}
	Engine.SetLogger(xorm.NewSimpleLogger(f))

}
