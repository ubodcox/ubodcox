package postgresqlx

import (
	sql "database/sql"
	"io/ioutil"
	"os"
	"sqlbar/server/src/config"
	"sqlbar/server/src/logs"

	sqlx "github.com/jmoiron/sqlx"
)

type (
	// PostgreDb type
	PostgreDb struct {
		*sql.DB
	}

	// PostgreDbx struct
	PostgreDbx struct {
		*sqlx.DB
	}
)

// Db var
var Db *PostgreDb

// Dbx var
var Dbx *PostgreDbx

// DbInit func
func DbInit() {
	logs.Log.PushFuncName("postgresql", "postgresql", "DbInit")
	defer logs.Log.PopFuncName()

	connStr := "user=" + config.DbUser + " " +
		"password=" + config.DbPassword + " " +
		"dbname=" + config.DbName + " " +
		"sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		logs.Log.Error("sql.Open", err)
	}

	connx, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logs.Log.Error("sql.Open", err)
	}

	Db = &PostgreDb{conn}
	Dbx = &PostgreDbx{connx}
}

// DbDeinit func
func DbDeinit() {
	Db.Close()
	Dbx.Close()
}

// DbUpdate func
func DbUpdate() {
	logs.Log.PushFuncName("postgresql", "postgresql", "DbUpdate")
	defer logs.Log.PopFuncName()

	b, err := ioutil.ReadFile("./files/sql/procedures.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'procedures.sql'", err)
		os.Exit(1)
	}

	sql := string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'procedures.sql'", err)
	}
	//------------------------------------------------
	b, err = ioutil.ReadFile("./files/sql/create.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'create.sql'", err)
		os.Exit(1)
	}

	sql = string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'create.sql'", err)
		return
	}
	//------------------------------------------------
	b, err = ioutil.ReadFile("./files/sql/exec.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'exec.sql'", err)
		os.Exit(1)
	}

	sql = string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'exec.sql'", err)
		return
	}
	//- chat -----------------------------------------
	/* b, err = ioutil.ReadFile("./files/sql/chat.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'chat.sql'", err)
		os.Exit(1)
	}

	sql = string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'chat.sql'", err)
		return
	}	 */

}
