package db

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateDB(t *testing.T) {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}

	//创建表
	sql_table := `
CREATE TABLE userinfo (
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(64) NULL,
	department VARCHAR(64) NULL,
	created DATE NULL
);

CREATE TABLE userdetail (
	uid INT(10) NULL,
	intro TEXT NULL,
	profile TEXT NULL,
	PRIMARY KEY (uid)
);
`
	db.Exec(sql_table)

	db.Close()

}
