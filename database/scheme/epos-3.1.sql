CREATE TABLE param_type
(
    code  TEXT PRIMARY KEY NOT NULL,
    title TEXT UNIQUE      NOT NULL,
    alias TEXT UNIQUE      NOT NULL
);

CREATE TABLE param_code
(
    code  TEXT PRIMARY KEY NOT NULL,
    title TEXT UNIQUE      NOT NULL,
    alias TEXT UNIQUE      NOT NULL
);

CREATE TABLE rel_type
(
    code  TEXT PRIMARY KEY NOT NULL,
    title TEXT UNIQUE      NOT NULL,
    alias TEXT UNIQUE      NOT NULL
);

CREATE TABLE content_type
(
    code  TEXT PRIMARY KEY NOT NULL,
    title TEXT UNIQUE      NOT NULL,
    alias TEXT UNIQUE      NOT NULL
);

CREATE TABLE entry
(
    id        TEXT PRIMARY KEY NOT NULL,
    added     INTEGER          NOT NULL,
    title     TEXT             NOT NULL,
    url       TEXT UNIQUE      NOT NULL,
    completed INTEGER          NOT NULL DEFAULT 0,
    ignored   INTEGER          NOT NULL DEFAULT 0,
    access    INTEGER          NOT NULL DEFAULT 5
);

CREATE TABLE object
(
    id        TEXT PRIMARY KEY NOT NULL,
    added     INTEGER          NOT NULL,
    title     TEXT             NOT NULL,
    completed INTEGER          NOT NULL DEFAULT 0,
    ignored   INTEGER          NOT NULL DEFAULT 0,
    access    INTEGER          NOT NULL DEFAULT 5
);

CREATE TABLE relation
(
    code        TEXT PRIMARY KEY NOT NULL,
    title       TEXT UNIQUE      NOT NULL,
    description TEXT             NOT NULL,
    type        TEXT             NOT NULL
);

CREATE TABLE content
(
    id     TEXT PRIMARY KEY NOT NULL,
    entry  TEXT             NOT NULL,
    title  TEXT             NOT NULL,
    type   TEXT             NOT NULL,
    data   TEXT             NOT NULL,
    access INTEGER          NOT NULL DEFAULT 5
);

CREATE TABLE entry_param
(
    id    TEXT PRIMARY KEY NOT NULL,
    code  TEXT             NOT NULL,
    type  TEXT             NOT NULL DEFAULT 'n\a',
    value TEXT             NOT NULL,
    extra TEXT             NOT NULL DEFAULT 'n\a'
);

CREATE TABLE object_param
(
    id    TEXT PRIMARY KEY NOT NULL,
    code  TEXT             NOT NULL,
    type  TEXT             NOT NULL,
    value TEXT             NOT NULL,
    extra TEXT             NOT NULL DEFAULT 'n\a'
);

CREATE TABLE r_entry_param
(
    entry TEXT NOT NULL,
    param TEXT NOT NULL,
    PRIMARY KEY (entry, param)
);

CREATE TABLE r_object_param
(
    object TEXT NOT NULL,
    param  TEXT NOT NULL,
    PRIMARY KEY (object, param)
);

CREATE TABLE r_object_relation
(
    object_a TEXT NOT NULL,
    relation TEXT NOT NULL,
    object_b TEXT NOT NULL,
    PRIMARY KEY (object_a, relation, object_b)
);

CREATE TABLE r_entry_object
(
    object TEXT NOT NULL,
    entry  TEXT NOT NULL,
    PRIMARY KEY (object, entry)
);