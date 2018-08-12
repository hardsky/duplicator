package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("create index on user_id column...")
		_, err := db.Exec(`CREATE INDEX user_idx ON conn_log (user_id)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping user_id index...")
		_, err := db.Exec(`DROP INDEX user_idx`)
		return err
	})
}
