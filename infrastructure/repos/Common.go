package repos

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var readEngine *xorm.Engine
var writeEngine *xorm.Engine

//InitDB 初始化DB引擎
func InitDB() {
	var err error
	readEngine, err = xorm.NewEngine("sqlite3", "kendopay.db")

	writeEngine, err = xorm.NewEngine("sqlite3", "kendopay.db")

	if err != nil {
		panic("不得了了！")
	}
}
