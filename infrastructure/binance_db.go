package infrastructure

import (
	"database/sql"
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

func InsertQuery(dbIndex int, query string) int64 {
	if dbIndex < len(dbList) && dbIndex > -1 {
		res, err := dbList[dbIndex].Exec(query)
		if err != nil {
			log.Println(err)
			return -1
		}

		c, err := res.RowsAffected()
		if err != nil {
			log.Println(err)
			return -1
		}
		return c
	}
	return -1
}

func StartProfiles(dbIndex int) {
	InsertQuery(dbIndex, `SET profiling=1;`)
}
