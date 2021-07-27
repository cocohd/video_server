package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	// open操作时并未连接，到Prepare才会连接
	// 大坑！！！ 在使用 := 时 会创建一个新的sqlDb变量,新的sqlDb会把全局变量sqlDb覆盖掉
	// 蠢逼
	dbConn, err = sql.Open("mysql", "root:756979.Hd@tcp(localhost:3306)/video_server")
	if err != nil {
		panic(err.Error())
	}

	// 好家伙，连接关闭，你连个jb？
	//defer dbConn.Close()
}
