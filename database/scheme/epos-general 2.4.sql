CREATE TABLE attr_type (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE attr_code (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE source_type (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    prefix TEXT NOT NULL UNIQUE
);

CREATE TABLE obj_type (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    prefix TEXT NOT NULL UNIQUE
);

CREATE TABLE rel_type (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE graph_extension (
    code TEXT PRIMARY KEY NOT NULL,
    value TEXT NOT NULL UNIQUE
);

CREATE TABLE attribute (
    id TEXT PRIMARY KEY NOT NULL,
    code TEXT NOT NULL,
    type TEXT NOT NULL,
    value TEXT NOT NULL,
    extra TEXT NOT NULL
);

CREATE TABLE source (
    id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    type TEXT NOT NULL
);

CREATE TABLE source_attr (
    source TEXT NOT NULL,
    attribute TEXT NOT NULL,
    PRIMARY KEY (source, attribute)
);

CREATE TABLE object (
    id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    type TEXT NOT NULL,
    add_date INTEGER NOT NULL
);

CREATE TABLE obj_source (
    object TEXT NOT NULL,
    source TEXT NOT NULL,
    PRIMARY KEY (object, source)
);

CREATE TABLE obj_attr (
    object TEXT NOT NULL,
    attribute TEXT NOT NULL,
    PRIMARY KEY (object, attribute)
);

CREATE TABLE relation (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    type TEXT NOT NULL
);

CREATE TABLE obj_relation (
    object_a TEXT NOT NULL,
    relation TEXT NOT NULL,
    object_b TEXT NOT NULL,
    PRIMARY KEY (object_a, relation, object_b)
);

CREATE TABLE graphics (
    id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    source TEXT NOT NULL,
    ext TEXT NOT NULL,
    url TEXT NOT NULL
);

CREATE TABLE graph_attr (
    graphics TEXT NOT NULL,
    attribute TEXT NOT NULL,
    PRIMARY KEY (graphics, attribute)
);

CREATE TABLE obj_graph (
    object TEXT NOT NULL,
    graphics TEXT NOT NULL,
    PRIMARY KEY (object, graphics)
);

