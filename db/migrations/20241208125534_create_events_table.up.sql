CREATE TABLE EVENTS (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    TITLE VARCHAR(255) NOT NULL,
    DESCRIPTION TEXT,
    ORGANIZER VARCHAR(255),
    LOCATION TEXT,
    LONGITUDE FLOAT,
    LATITUDE FLOAT,
    START_AT DATETIME,
    END_AT DATETIME,
    CATEGORY VARCHAR(100),
    PRICE DECIMAL(10, 2),
    CREATED_AT DATETIME DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT DATETIME DEFAULT CURRENT_TIMESTAMP
);