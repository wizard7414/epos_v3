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
    description TEXT NOT NULL
);

CREATE TABLE obj_type (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    prefix TEXT NOT NULL UNIQUE
);

CREATE TABLE rel_type (
    code TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE rel_code (
    id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE
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
    id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    type TEXT NOT NULL,
    code TEXT NOT NULL
);

CREATE TABLE obj_relation (
    object_a TEXT NOT NULL,
    object_b TEXT NOT NULL,
    relation TEXT NOT NULL,
    PRIMARY KEY (object_a, object_b, relation)
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



INSERT INTO attr_type VALUES('string', 'String', 'Строка');
INSERT INTO attr_type VALUES('integer', 'Integer', 'Целое число');
INSERT INTO attr_type VALUES('url', 'URL', 'Гиперссылка');
INSERT INTO attr_type VALUES('date', 'Date', 'Дата');

INSERT INTO attr_code VALUES('name', 'Name', 'Имя');
INSERT INTO attr_code VALUES('aliase', 'Aliase', 'Псевдоним');
INSERT INTO attr_code VALUES('devArtGalery', 'DeviantArt gallery', 'Работы на DeviantArt');

INSERT INTO source_type VALUES('illustrator', 'Illustrator', 'Художник');
INSERT INTO source_type VALUES('game', 'Video game', 'Видео игра');
INSERT INTO source_type VALUES('comics', 'Comics universe', 'Комикс');

INSERT INTO obj_type VALUES('character', 'Character', 'Персонаж', 'CH');
INSERT INTO obj_type VALUES('unit', 'Unit', 'Юнит', 'UN');
INSERT INTO obj_type VALUES('creature', 'Creature', 'Существо', 'CR');

INSERT INTO source VALUES('SRC0000001', 'nicofari', 'illustrator');
INSERT INTO source VALUES('SRC0000002', 'donmeiklejohn', 'illustrator');
INSERT INTO source VALUES('SRC0000003', 'nyaka-n', 'illustrator');
INSERT INTO source VALUES('SRC0000004', 'reikun12', 'illustrator');
INSERT INTO source VALUES('SRC0000005', 'prattyvee', 'illustrator');
INSERT INTO source VALUES('SRC0000006', 'prywinko', 'illustrator');
INSERT INTO source VALUES('SRC0000007', 'themaestronoob', 'illustrator');
INSERT INTO source VALUES('SRC0000008', '4linex', 'illustrator');
INSERT INTO source VALUES('SRC0000009', 'darkeyez07', 'illustrator');
INSERT INTO source VALUES('SRC0000010', 'judaljiuda', 'illustrator');
INSERT INTO source VALUES('SRC0000011', 'raikoart', 'illustrator');
INSERT INTO source VALUES('SRC0000012', 'felox08', 'illustrator');
INSERT INTO source VALUES('SRC0000013', '7th--heaven', 'illustrator');
INSERT INTO source VALUES('SRC0000014', 'shibashake', 'illustrator');
INSERT INTO source VALUES('SRC0000015', 'diablo7707', 'illustrator');
INSERT INTO source VALUES('SRC0000016', 'luminyu', 'illustrator');

INSERT INTO attribute VALUES('AT0000000001', 'aliase', 'string', 'NicoFari', 'n/a');
INSERT INTO attribute VALUES('AT0000000002', 'devArtGalery', 'url', 'https://www.deviantart.com/nicofari/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000003', 'aliase', 'string', 'DonMeiklejohn', 'n/a');
INSERT INTO attribute VALUES('AT0000000004', 'aliase', 'string', 'DJM', 'n/a');
INSERT INTO attribute VALUES('AT0000000005', 'devArtGalery', 'url', 'https://www.deviantart.com/donmeiklejohn/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000006', 'aliase', 'string', 'Nyaka-N', 'n/a');
INSERT INTO attribute VALUES('AT0000000007', 'aliase', 'string', 'BaSha', 'n/a');
INSERT INTO attribute VALUES('AT0000000008', 'devArtGalery', 'url', 'https://www.deviantart.com/nyaka-n/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000009', 'devArtGalery', 'url', 'https://www.deviantart.com/reikun12/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000010', 'aliase', 'string', 'PrattyVee', 'n/a');
INSERT INTO attribute VALUES('AT0000000011', 'devArtGalery', 'url', 'https://www.deviantart.com/prattyvee/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000012', 'aliase', 'string', 'Prywinko', 'n/a');
INSERT INTO attribute VALUES('AT0000000013', 'devArtGalery', 'url', 'https://www.deviantart.com/prywinko/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000014', 'aliase', 'string', 'TheMaestroNoob', 'n/a');
INSERT INTO attribute VALUES('AT0000000015', 'devArtGalery', 'url', 'https://www.deviantart.com/themaestronoob/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000016', 'aliase', 'string', '4LineX', 'n/a');
INSERT INTO attribute VALUES('AT0000000017', 'devArtGalery', 'url', 'https://www.deviantart.com/4linex/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000018', 'devArtGalery', 'url', 'https://www.deviantart.com/darkeyez07/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000019', 'devArtGalery', 'url', 'https://www.deviantart.com/judaljiuda/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000020', 'devArtGalery', 'url', 'https://www.deviantart.com/raikoart/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000021', 'aliase', 'string', 'Felox08', 'n/a');
INSERT INTO attribute VALUES('AT0000000022', 'devArtGalery', 'url', 'https://www.deviantart.com/felox08/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000023', 'aliase', 'string', '7th--Heaven', 'n/a');
INSERT INTO attribute VALUES('AT0000000024', 'devArtGalery', 'url', 'https://www.deviantart.com/7th--heaven/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000025', 'aliase', 'string', 'Shiba Shake', 'n/a');
INSERT INTO attribute VALUES('AT0000000026', 'devArtGalery', 'url', 'https://www.deviantart.com/shibashake/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000027', 'aliase', 'string', 'Diablo7707', 'n/a');
INSERT INTO attribute VALUES('AT0000000028', 'aliase', 'string', 'Diablo', 'n/a');
INSERT INTO attribute VALUES('AT0000000029', 'devArtGalery', 'url', 'https://www.deviantart.com/diablo7707/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000030', 'aliase', 'string', 'LumiNyu', 'n/a');
INSERT INTO attribute VALUES('AT0000000031', 'devArtGalery', 'url', 'https://www.deviantart.com/luminyu/gallery/all', 'n/a');

INSERT INTO source_attr VALUES('SRC0000001', 'AT0000000001');
INSERT INTO source_attr VALUES('SRC0000001', 'AT0000000002');
INSERT INTO source_attr VALUES('SRC0000002', 'AT0000000003');
INSERT INTO source_attr VALUES('SRC0000002', 'AT0000000004');
INSERT INTO source_attr VALUES('SRC0000002', 'AT0000000005');
INSERT INTO source_attr VALUES('SRC0000003', 'AT0000000006');
INSERT INTO source_attr VALUES('SRC0000003', 'AT0000000007');
INSERT INTO source_attr VALUES('SRC0000003', 'AT0000000008');
INSERT INTO source_attr VALUES('SRC0000004', 'AT0000000009');
INSERT INTO source_attr VALUES('SRC0000005', 'AT0000000010');
INSERT INTO source_attr VALUES('SRC0000005', 'AT0000000011');
INSERT INTO source_attr VALUES('SRC0000006', 'AT0000000012');
INSERT INTO source_attr VALUES('SRC0000006', 'AT0000000013');
INSERT INTO source_attr VALUES('SRC0000007', 'AT0000000014');
INSERT INTO source_attr VALUES('SRC0000007', 'AT0000000015');
INSERT INTO source_attr VALUES('SRC0000008', 'AT0000000016');
INSERT INTO source_attr VALUES('SRC0000008', 'AT0000000017');
INSERT INTO source_attr VALUES('SRC0000009', 'AT0000000018');
INSERT INTO source_attr VALUES('SRC0000010', 'AT0000000019');
INSERT INTO source_attr VALUES('SRC0000011', 'AT0000000020');
INSERT INTO source_attr VALUES('SRC0000012', 'AT0000000021');
INSERT INTO source_attr VALUES('SRC0000012', 'AT0000000022');
INSERT INTO source_attr VALUES('SRC0000013', 'AT0000000023');
INSERT INTO source_attr VALUES('SRC0000013', 'AT0000000024');
INSERT INTO source_attr VALUES('SRC0000014', 'AT0000000025');
INSERT INTO source_attr VALUES('SRC0000014', 'AT0000000026');
INSERT INTO source_attr VALUES('SRC0000015', 'AT0000000027');
INSERT INTO source_attr VALUES('SRC0000015', 'AT0000000028');
INSERT INTO source_attr VALUES('SRC0000015', 'AT0000000029');
INSERT INTO source_attr VALUES('SRC0000016', 'AT0000000030');
INSERT INTO source_attr VALUES('SRC0000016', 'AT0000000031');
