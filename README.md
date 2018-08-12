# Duplicator

Duplicator is a simple service, that allow to find user duplicates in connection log.

For example we have following records in table conn_log

 user_id |  ip_addr  |         ts
---------+-----------+---------------------
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
---------+-----------+---------------------
       1 | 127.0.0.1 | 2018-01-01 17:51:59
       2 | 127.0.0.1 | 2018-01-01 17:52:59
       1 | 127.0.0.2 | 2018-01-01 17:53:59
       2 | 127.0.0.2 | 2018-01-01 17:54:59

We use *dep* as dependency manager.
