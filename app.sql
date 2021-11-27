CREATE TABLE IF NOT EXISTS POST
(
    ID      int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TITLE   varchar(255),
    CONTENT text,
    AUTHOR  int NOT NULL
);

CREATE TABLE IF NOT EXISTS AUTHOR
(
    ID     int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    NAME   varchar(255)
);