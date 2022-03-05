CREATE TABLE object_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE attribute_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE attribute_code (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE resource (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL UNIQUE
);

CREATE TABLE source (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL UNIQUE,
    resource INTEGER NOT NULL,
    change_date INTEGER NOT NULL
);

CREATE TABLE entry (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    source INTEGER NOT NULL,
    url TEXT NOT NULL,
    add_date INTEGER NOT NULL
);

CREATE TABLE object (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    entry INTEGER NOT NULL,
    url TEXT NOT NULL,
    add_date INTEGER NOT NULL,
    object_type INTEGER NOT NULL
);

CREATE TABLE attribute (
    id INTEGER PRIMARY KEY NOT NULL,
    code TEXT NOT NULL,
    value TEXT NOT NULL,
    attribute_type INTEGER NOT NULL
);

CREATE TABLE object_attribute (
    object INTEGER NOT NULL,
    attribute INTEGER NOT NULL,
    PRIMARY KEY (object, attribute)
);
