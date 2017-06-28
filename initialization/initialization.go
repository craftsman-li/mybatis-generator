package initialization

import (
	"database/sql"
	"flag"
	"fmt"
	"mybatis-generator/constant"
	"mybatis-generator/util"
)

// 运行时,初始化信息

// Host 主机地址
var Host = flag.String("h", "127.0.0.1:3306", "数据库地址(ip:port)")

// UserName 用户名
var UserName = flag.String("u", "root", "数据库用户名")

// Pass 密码
var Pass = flag.String("pass", "", "数据库密码")

// db 数据库名
var db = flag.String("db", "", "数据库名")

// Table 表名
var Table = flag.String("t", "", "数据表名")

// Packages 包名
var Packages = flag.String("package", "", "包名(用于生成resultMap的type属性)")

// Splitter 表名分隔符
var Splitter = flag.String("f", "_", "表名分隔符")

// DB 数据库连接
var DB *sql.DB

// Init 初始化方法
func Init() {
	flag.Parse()
	DB = openDbLink()
}

func openDbLink() *sql.DB {
	return openDBLink(UserName, Pass, Host, db)
}

// openDBLink 打开数据库连接
// @param userName 数据库用户名
// @param pass 数据库密码
// @param host 数据库地址
// @para database 数据库名
func openDBLink(userName *string, pass *string, host *string, database *string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", *userName, *pass, *host, *database))
	util.CheckError(err, constant.LinkDBFail)
	err = db.Ping()
	util.CheckError(err, constant.LinkDBFail)
	return db
}
