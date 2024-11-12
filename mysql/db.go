package mysql

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	MainDB sqlx.SqlConn
)

const (
	TypeMain = "main" // 主库
)

// DBCfg DB配置
type DBCfg struct {
	Name     string
	Host     string
	Port     int
	User     string
	Passwd   string
	Database string
}

// Init 初始化
func Init(list []DBCfg) {
	for _, item := range list {
		address := fmt.Sprint(item.Host, ":", item.Port)
		dataSource := item.User + ":" + item.Passwd + "@tcp(" + address + ")/" + item.Database +
			"?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
		client := sqlx.NewMysql(dataSource)
		switch item.Name {
		case TypeMain:
			MainDB = client
		}
	}
}
