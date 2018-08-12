# Duplicator

Duplicator is a toy service, that allow to find user duplicates in connection log.

For example we have following records in table conn_log

 user_id |  ip_addr  |         ts
-----------------------------------------
       1 | 127.0.0.1 | 2018-01-01 17:51:59
       2 | 127.0.0.1 | 2018-01-01 17:52:59
       1 | 127.0.0.2 | 2018-01-01 17:53:59
       2 | 127.0.0.2 | 2018-01-01 17:54:59
       2 | 127.0.0.3 | 2018-01-01 17:55:59
       3 | 127.0.0.3 | 2018-01-01 17:55:59
       3 | 127.0.0.1 | 2018-01-01 17:56:59
       4 | 127.0.0.1 | 2018-01-01 17:57:59


We consider, that two user_id are duplicates, if they at least two times matched by ip_addr.
So following 4 records indicate that users with ids: 1 and 2 are duplicates:

 user_id |  ip_addr  |         ts
-----------------------------------------
       1 | 127.0.0.1 | 2018-01-01 17:51:59
       2 | 127.0.0.1 | 2018-01-01 17:52:59
       1 | 127.0.0.2 | 2018-01-01 17:53:59
       2 | 127.0.0.2 | 2018-01-01 17:54:59

We use *dep* as dependency manager.

env variables:

- DP_DEBUG - turn on debvug logs
- DP_ADDR - service address (for example localhost:8080)
- DP_DB_ADDR - db address (for example :5432)
- DP_DB_USER - db user
- DP_DB_PSW - db user password
- DP_DB_DATABASE - service database

How to use:
1. create database for application (let's name it 'duplicator')

2. apply migrations from */migrations* directory, for example
```
cd migrations/
env DP_DB_USER=postgres DP_DB_PSW=<password_for_db_user> DP_DB_DATABASE=duplicator DP_DB_ADDR=:5432 go run *.go up
```
(see readme in migrations/ dir.)

3. run service with command
`env DP_ADDR=:8080 DP_DB_USER=postgres DP_DB_PSW=<password_for_db_user> DP_DB_DATABASE=duplicator DP_DB_ADDR=:5432 go run main.go`
or with debug logs  
`env DP_DEBUG=1 DP_ADDR=:8080 DP_DB_USER=postgres DP_DB_PSW=<password_for_db_user> DP_DB_DATABASE=duplicator DP_DB_ADDR=:5432 go run main.go`

4. apply sql scripts with test data from */generators* directory (if predefined 1M records are not enough it can be changed to other number in sql scripts.)

5. run http request http://localhost:8080/duplicator/duplicate/{userId}/{userId} where ids are integers.
