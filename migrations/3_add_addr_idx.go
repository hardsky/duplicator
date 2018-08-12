package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("create index on ip_addr column...")
		_, err := db.Exec(`CREATE INDEX addr_idx ON conn_log (ip_addr)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping i_addr index...")
		_, err := db.Exec(`DROP INDEX addr_idx`)
		return err
	})
}
