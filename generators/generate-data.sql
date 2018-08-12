INSERT INTO conn_log
SELECT (1000*random())::int AS user_id,
    '127.0.0.'::text || (100*random())::int::text AS ip_addr,
    current_timestamp
FROM generate_series(1,1000000);
