
This directory contains migrations for database.
Its are sql scripts injected in go files on top of go-pg ORM.
see https://github.com/go-pg/migrations

use command:  
`go run *.go up` - to apply all migrations on database  
run  
`go run *.go --help` - to read usage details.

We use *dep* as dependency manager.

env variables:

- DP_DEBUG
- DP_ADDR
- DP_DB_ADDR
- DP_DB_USER
- DP_DB_PSW
- DP_DB_DATABASE
