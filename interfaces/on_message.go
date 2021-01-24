package interfaces

import (
	"binance/listener/infrastructure"
	"fmt"
	"strings"
	"time"

	"github.com/sacOO7/gowebsocket"
	"github.com/tidwall/gjson"
)

func HandleMessage(message string, _ gowebsocket.Socket) {
	result := gjson.GetMany(message, "stream", "data.bids", "data.asks")
	streamName := strings.Join(
		strings.Split(
			result[0].String(), "@",
		)[:1],
		"",
	)

	tableName, dbConfig := infrastructure.GetBinanceIndex(streamName)
	if dbConfig == -1 {
		return
	}
	infrastructure.StartProfiles(dbConfig)
	tsWS := GetFormatDateNow()
	query := fmt.Sprintf(
		"INSERT INTO binance_%v (timestamp, buy, sell ,ts_ws, ts_data ) VALUES (NOW(), '%s', '%s', '%s', '%s');",
		tableName,
		result[1].String(),
		result[2].String(),
		tsWS,
		GetFormatDateNow(),
	)
	id, _ := infrastructure.InsertQuery(dbConfig, query)
	if id == -1 {
		return
	}
	duration := infrastructure.GetProfilesV2(dbConfig, query)

	query = fmt.Sprintf(
		"update binance_%v set time_execute = '%v' where id = %d",
		tableName,
		duration,
		id,
	)
	infrastructure.InsertQuery(dbConfig, query)
}

func GetFormatDateNow() string {
	now := time.Now().Local()
	return fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}
