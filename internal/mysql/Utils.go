package mysql

import (
	"database/sql"
	"github.com/Musik-Bot/Musik-Bot/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetConn() (conn *sql.DB) {
	cfg, err := config.Parse()
	if err != nil {
		panic(err)
	}
	connstr := cfg.Database.Username + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ")/" + cfg.Database.Database
	conn, err = sql.Open("mysql", connstr)
	if err != nil {
		panic(err)
		return
	} else {
		return
	}
}
