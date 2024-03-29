CREATE TABLE object_type (
    code INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title TEXT NOT NULL
);

CREATE TABLE attribute_type (
    code INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title TEXT NOT NULL
);

CREATE TABLE source (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    change_date INTEGER NOT NULL
);

CREATE TABLE object (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title TEXT NOT NULL,
    source INTEGER NOT NULL,
    url TEXT NOT NULL,
    add_date INTEGER NOT NULL,
    type INTEGER NOT NULL
);

CREATE TABLE attribute (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    code TEXT NOT NULL,
    value TEXT NOT NULL,
    type INTEGER NOT NULL
);

CREATE TABLE object_attribute (
    object INTEGER NOT NULL,
    attribute INTEGER NOT NULL,
    PRIMARY KEY (object, attribute)
);
