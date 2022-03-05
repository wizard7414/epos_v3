CREATE TABLE attribute_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE attribute_code (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE attribute (
    id INTEGER PRIMARY KEY NOT NULL,
    attribute_code INTEGER NOT NULL,
    attribute_type INTEGER NOT NULL,
    title TEXT NOT NULL
);

CREATE TABLE object_class (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE object_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE object_subtype (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE object (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    add_date INTEGER NOT NULL,
    object_class INTEGER NOT NULL,
    object_type INTEGER NOT NULL,
    object_subtype INTEGER NOT NULL
);

CREATE TABLE object_attribute (
    object INTEGER NOT NULL,
    attribute INTEGER NOT NULL,
    PRIMARY KEY (object, attribute)
);

CREATE TABLE relation_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE relation_code (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE relation (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    relation_type INTEGER NOT NULL,
    relation_code INTEGER NOT NULL
);

CREATE TABLE object_relation (
    object_a INTEGER NOT NULL,
    object_b INTEGER NOT NULL,
    relation INTEGER NOT NULL,
    PRIMARY KEY (object_a, object_b, relation)
);

CREATE TABLE graphics (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    extension INTEGER NOT NULL,
    url TEXT NOT NULL
);

CREATE TABLE graphics_attribute (
    graphics INTEGER NOT NULL,
    attribute INTEGER NOT NULL,
    PRIMARY KEY (graphics, attribute)
);

CREATE TABLE object_graphics (
    object INTEGER NOT NULL,
    graphics INTEGER NOT NULL,
    PRIMARY KEY (object, graphics)
);
