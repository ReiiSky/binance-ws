package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
)

var dbList = []*sql.DB{}

func InitBinanceDB(url []string, dbname []string, user []string, password []string) {

	urlLen := len(url)
	dbnameLen := len(dbname)
	userLen := len(user)
	passwordLen := len(password)

	if !(urlLen == dbnameLen && urlLen == userLen && urlLen == passwordLen) {
		panic("Error URL, DBNAME, USER, PASSWORD length must be the same")
	}
	for i := 0; i < urlLen; i++ {
		db, err := OpenMySQL(url[i], dbname[i], user[i], password[i])
		if err != nil {
			log.Println(err, "\n")
		}
		dbList = append(dbList,
			db,
		)
	}
}

func InsertQuery(dbIndex int, query string) (int64, int64) {
	if dbIndex < len(dbList) && dbIndex > -1 {
		res, err := dbList[dbIndex].Exec(query)
		if err != nil {
			log.Println(err)
			return -1, -1
		}

		a, err := res.RowsAffected()
		if err != nil {
			log.Println(err)
			return -1, -1
		}
		i, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			return -1, -1
		}
		return i, a
	}
	return -1, -1
}

func RowQuery(dbIndex int, query string, result interface{}) {
	if dbIndex < len(dbList) && dbIndex > -1 {
		res := dbList[dbIndex].QueryRow(query)
		err := res.Err()
		if err != nil {
			log.Println(err)
		}
		err = res.Scan(result)
		if err != nil {
			log.Println(err)
		}
	}
}

func StartProfiles(dbIndex int) {
	InsertQuery(dbIndex, `SET profiling=1;`)
}

func GetProfiles(dbIndex int, id int64) float64 {
	var result float64
	RowQuery(dbIndex, fmt.Sprintf(`SELECT sum(duration) as duration FROM INFORMATION_SCHEMA.PROFILING WHERE QUERY_ID=%d;`, id), &result)
	return result
}
