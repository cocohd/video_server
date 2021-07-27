create table sessions (
    session_id VARCHAR(64) PRIMARY KEY NOT NULL,
    TTL TINYTEXT,   //过期时间
    username VARCHAR(64)
);