CREATE TABLE users
(
    id     SERIAL,
    reg_no VARCHAR(20) UNIQUE NOT NULL,
    NAME   VARCHAR(50)        DEFAULT '',
    phone  VARCHAR(20) UNIQUE DEFAULT ''
);

CREATE INDEX users_reg
    ON users (reg_no);