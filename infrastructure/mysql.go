package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// OpenMySQL connection to access mysql
// return db connector, and error if connecting is fail
func OpenMySQL(host, dbname, username, password string) (*sql.DB, error) {

	mysqlConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)
	connector, _ := sql.Open("mysql", mysqlConfig)
	err := connector.Ping()

	if err != nil {
		return nil, err
	}

	return connector, nil
}
