package sql

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Run() {
	r := gin.Default()
	// 链接数据库
	InitializeDB("relearn_mysql", "root", "", "", "utf8")
	InitRouter(r)
	err := r.Run()
	if err != nil {
		return
	}
}
