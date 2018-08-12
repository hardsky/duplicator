package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table conn_log...")
		_, err := db.Exec(`CREATE TABLE conn_log( user_id bigint, ip_addr varchar(15), ts timestamp )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table conn_log...")
		_, err := db.Exec(`DROP TABLE conn_log`)
		return err
	})
}
