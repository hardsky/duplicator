package db

import (
	"os"

	"github.com/go-pg/pg"
)

var con *pg.DB

func init() {
	con = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DP_DB_ADDR"),
		User:     os.Getenv("DP_DB_USER"),
		Password: os.Getenv("DP_DB_PSW"),
		Database: os.Getenv("DP_DB_DATABASE"),
	})
}

func IsDuplicate(userID1, userID2 string) bool {
	var count int
	con.QueryOne(pg.Scan(&count), `
SELECT COUNT (1) FROM
(SELECT ip_addr FROM conn_log WHERE user_id = ?) c1
INNER JOIN
(SELECT ip_addr FROM conn_log WHERE user_id = ?) c2
ON c1.ip_addr = c2.ip_addr
`, userID1, userID2)

	return count >= 2
}
