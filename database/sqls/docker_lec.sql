CREATE TABLE IF NOT EXISTS Students(
    id CHAR(30) NOT NULL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(100)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
