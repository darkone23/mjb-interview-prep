-- Up migration steps
CREATE TABLE users (
                       user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                       username VARCHAR(64) NOT NULL UNIQUE,
                       password VARCHAR(128) NOT NULL UNIQUE
);
