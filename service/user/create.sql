CREATE TABLE User
(id
    BINARY(16)
    UNIQUE
    NOT NULL
    DEFAULT (UUID_TO_BIN(UUID()))
,name
    VARCHAR(64)
    NOT NULL
    CHECK (LENGTH(name) BETWEEN 3 AND 64)
,email
    VARCHAR(255)
    UNIQUE
    NOT NULL
    CHECK (email REGEXP '^[^@[:space:]]+@[^@[:space:]\.]+\.[^@\.[:space:]]+$')
,password
    BINARY(60)
    NOT NULL
,registered_at
    DATETIME
    NOT NULL
    DEFAULT CURRENT_TIMESTAMP
,last_modified_at
    DATETIME
    NOT NULL
    DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP

,PRIMARY KEY (id)
);

CREATE TABLE Access
(id
    BINARY(16)
    UNIQUE
    NOT NULL
    DEFAULT (UUID_TO_BIN(UUID()))
,userID
    BINARY(16)
    NOT NULL
,created_at
    DATETIME
    NOT NULL
    DEFAULT CURRENT_TIMESTAMP

,PRIMARY KEY (id)
,FOREIGN KEY (userID) REFERENCES User(id)
);