CREATE TABLE attr_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE attr_code (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE source_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE obj_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    prefix TEXT NOT NULL UNIQUE
);

CREATE TABLE rel_type (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE rel_code (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE graph_extension (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);




CREATE TABLE attribute (
    id INTEGER PRIMARY KEY NOT NULL,
    a_code INTEGER NOT NULL,
    a_type INTEGER NOT NULL,
    a_title TEXT NOT NULL,
    a_value TEXT NOT NULL,
    a_extra TEXT NOT NULL
);

CREATE TABLE source (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    s_type INTEGER NOT NULL
);

CREATE TABLE source_attr (
    source INTEGER NOT NULL,
    attribute INTEGER NOT NULL,
    PRIMARY KEY (source, attribute)
);

CREATE TABLE object (
    id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    o_type INTEGER NOT NULL,
    add_date INTEGER NOT NULL
);

CREATE TABLE obj_source (
    object TEXT NOT NULL,
    source INTEGER NOT NULL,
    PRIMARY KEY (object, source)
);

CREATE TABLE obj_attr (
    object TEXT NOT NULL,
    attribute INTEGER NOT NULL,
    PRIMARY KEY (object, attribute)
);

CREATE TABLE relation (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    r_type INTEGER NOT NULL,
    r_code INTEGER NOT NULL
);

CREATE TABLE obj_relation (
    object_a TEXT NOT NULL,
    object_b INTEGER NOT NULL,
    relation INTEGER NOT NULL,
    PRIMARY KEY (object_a, object_b, relation)
);

CREATE TABLE graphics (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    source INTEGER NOT NULL,
    extension INTEGER NOT NULL,
    url TEXT NOT NULL
);

CREATE TABLE graph_attr (
    graphics INTEGER NOT NULL,
    attribute INTEGER NOT NULL,
    PRIMARY KEY (graphics, attribute)
);

CREATE TABLE obj_graph (
    object TEXT NOT NULL,
    graphics INTEGER NOT NULL,
    PRIMARY KEY (object, graphics)
);






INSERT INTO source_type VALUES(71001, 'Illustrator', 'Художник');
INSERT INTO source_type VALUES(71002, 'Video game', 'Видео игра');

INSERT INTO obj_type VALUES(81001, 'Character', 'Персонаж', 'CH');
INSERT INTO obj_type VALUES(81002, 'Unit', 'Юнит', 'UN');
INSERT INTO obj_type VALUES(81003, 'Creature', 'Существо', 'CR');


