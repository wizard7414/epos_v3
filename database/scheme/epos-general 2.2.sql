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
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE rel_code (
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
INSERT INTO attr_type VALUES('rate', 'Rating', 'Рейтинг');

INSERT INTO attr_code VALUES('name', 'Name', 'Имя');
INSERT INTO attr_code VALUES('aliase', 'Aliase', 'Псевдоним');
INSERT INTO attr_code VALUES('devArtGalery', 'DeviantArt gallery', 'Работы на DeviantArt');
INSERT INTO attr_code VALUES('type', 'Type', 'Тип');
INSERT INTO attr_code VALUES('rarity', 'Rarity', 'Редкость');
INSERT INTO attr_code VALUES('element', 'Element', 'Элемент');

INSERT INTO source_type VALUES('illustrator', 'Illustrator', 'Художник');
INSERT INTO source_type VALUES('game', 'Video game', 'Видео игра');
INSERT INTO source_type VALUES('comics', 'Comics universe', 'Комикс');
INSERT INTO source_type VALUES('anime_manga', 'Anime / Manga', 'Анимэ и Манга');
INSERT INTO source_type VALUES('movie', 'Movie universe', 'Кино-вселенная');

INSERT INTO obj_type VALUES('character', 'Character', 'Персонаж', 'CH');
INSERT INTO obj_type VALUES('unit', 'Unit', 'Юнит', 'UN');
INSERT INTO obj_type VALUES('creature', 'Creature', 'Существо', 'CR');
INSERT INTO obj_type VALUES('real_technics', 'Real Technics', 'Реальная техника', 'RT');
INSERT INTO obj_type VALUES('fict_technics', 'Fictional Technics', 'Вымышленная техника', 'FT');

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
INSERT INTO source VALUES('SRC0000017', 'Age of Ishtaria', 'game');
INSERT INTO source VALUES('SRC0000018', 'Attack on Titan', 'anime_manga');
INSERT INTO source VALUES('SRC0000019', 'DanMachi', 'anime_manga');
INSERT INTO source VALUES('SRC0000020', 'Dungeons & Dragons', 'game');
INSERT INTO source VALUES('SRC0000021', 'Dota 2', 'game');
INSERT INTO source VALUES('SRC0000022', 'moonarc', 'illustrator');
INSERT INTO source VALUES('SRC0000023', 'phamoz', 'illustrator');
INSERT INTO source VALUES('SRC0000024', 'blackxsensei', 'illustrator');
INSERT INTO source VALUES('SRC0000025', 'Fate/Grand Order', 'game');
INSERT INTO source VALUES('SRC0000026', 'Battle Chasers', 'comics');
INSERT INTO source VALUES('SRC0000027', 'Coffin Comics', 'comics');
INSERT INTO source VALUES('SRC0000028', 'Crusade Comics', 'comics');
INSERT INTO source VALUES('SRC0000029', 'Star Wars', 'movie');

INSERT INTO object VALUES('CH0000001', 'Nimrod', 'character', 1622203757);
INSERT INTO object VALUES('CH0000002', 'Mikasa Ackerman', 'character', 1622206125);
INSERT INTO object VALUES('CH0000003', 'Hestia', 'character', 1622234181);
INSERT INTO object VALUES('CH0000004', 'Traxex', 'character', 1622235342);
INSERT INTO object VALUES('CH0000005', 'Oseh', 'character', 1622235103);
INSERT INTO object VALUES('CH0000006', 'Herodias', 'character', 1622289227);
INSERT INTO object VALUES('CH0000007', 'Hildr', 'character', 1622290178);
INSERT INTO object VALUES('CH0000008', 'Aenon', 'character', 1622290724);
INSERT INTO object VALUES('CH0000009', 'Lich', 'character', 1622291050);
INSERT INTO object VALUES('CH0000010', 'Agaliarept', 'character', 1622437275);
INSERT INTO object VALUES('CH0000011', 'Agares', 'character', 1622437761);
INSERT INTO object VALUES('CH0000012', 'Samyaza', 'character', 1622437947);
INSERT INTO object VALUES('CH0000013', 'Scáthach', 'character', 1622438307);
INSERT INTO object VALUES('CH0000014', 'Scheherazade', 'character', 1622438629);
INSERT INTO object VALUES('CH0000015', 'Red Monika', 'character', 1622438899);
INSERT INTO object VALUES('CH0000016', 'Lady Death', 'character', 1622447314);
INSERT INTO object VALUES('CH0000017', 'Shi', 'character', 1622447453);

INSERT INTO object VALUES('CR0000001', 'Orc Shaman', 'creature', 1622234930);

INSERT INTO object VALUES('FT0000001', 'AT-AT', 'fict_technics', 1622448445);

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
INSERT INTO attribute VALUES('AT0000000032', 'type', 'string', 'Pound', 'n/a');
INSERT INTO attribute VALUES('AT0000000033', 'type', 'string', 'Slice', 'n/a');
INSERT INTO attribute VALUES('AT0000000034', 'type', 'string', 'Flurry', 'n/a');
INSERT INTO attribute VALUES('AT0000000035', 'element', 'string', 'Null', 'n/a');
INSERT INTO attribute VALUES('AT0000000036', 'element', 'string', 'Water', 'n/a');
INSERT INTO attribute VALUES('AT0000000037', 'element', 'string', 'Earth', 'n/a');
INSERT INTO attribute VALUES('AT0000000038', 'element', 'string', 'Fire', 'n/a');
INSERT INTO attribute VALUES('AT0000000039', 'rarity', 'rate', '★★★★★★★★', '8');
INSERT INTO attribute VALUES('AT0000000040', 'rarity', 'rate', '★ ★ ★ ★ ★ ★ ★', 'P7');
INSERT INTO attribute VALUES('AT0000000041', 'rarity', 'rate', '★★★★★★★', '7');
INSERT INTO attribute VALUES('AT0000000042', 'rarity', 'rate', '★★★★★★', '6');
INSERT INTO attribute VALUES('AT0000000043', 'rarity', 'rate', '★★★★★', '5');
INSERT INTO attribute VALUES('AT0000000044', 'rarity', 'rate', '★★★★', '4');
INSERT INTO attribute VALUES('AT0000000045', 'rarity', 'rate', '★★★', '3');
INSERT INTO attribute VALUES('AT0000000046', 'rarity', 'rate', '★★', '2');
INSERT INTO attribute VALUES('AT0000000047', 'rarity', 'rate', '★', '1');
INSERT INTO attribute VALUES('AT0000000048', 'aliase', 'string', 'Shingeki no Kyojin', 'n/a');
INSERT INTO attribute VALUES('AT0000000049', 'aliase', 'string', '進撃の巨人', 'n/a');
INSERT INTO attribute VALUES('AT0000000050', 'aliase', 'string', 'Dungeon ni Deai wo Motomeru no wa Machigatteiru Darou ka', 'n/a');
INSERT INTO attribute VALUES('AT0000000051', 'aliase', 'string', 'Is It Wrong to Try to Pick Up Girls in a Dungeon?', 'n/a');
INSERT INTO attribute VALUES('AT0000000052', 'aliase', 'string', 'ダンジョンに出会いを求めるのは間違っているだろうか', 'n/a');
INSERT INTO attribute VALUES('AT0000000053', 'aliase', 'string', 'ダンまち', 'n/a');
INSERT INTO attribute VALUES('AT0000000054', 'aliase', 'string', 'D&D', 'n/a');
INSERT INTO attribute VALUES('AT0000000055', 'aliase', 'string', 'DnD', 'n/a');
INSERT INTO attribute VALUES('AT0000000056', 'aliase', 'string', 'Подземелья и драконы', 'n/a');
INSERT INTO attribute VALUES('AT0000000057', 'aliase', 'string', 'Drow Ranger', 'n/a');
INSERT INTO attribute VALUES('AT0000000058', 'aliase', 'string', 'Moonarc', 'n/a');
INSERT INTO attribute VALUES('AT0000000059', 'aliase', 'string', 'Audia Pahlevi', 'n/a');
INSERT INTO attribute VALUES('AT0000000060', 'devArtGalery', 'url', 'https://www.deviantart.com/moonarc/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000061', 'aliase', 'string', 'ZOMA', 'n/a');
INSERT INTO attribute VALUES('AT0000000062', 'devArtGalery', 'url', 'https://www.deviantart.com/phamoz/gallery/all', 'n/a');
INSERT INTO attribute VALUES('AT0000000063', 'aliase', 'string', 'BlackxSensei', 'n/a');
INSERT INTO attribute VALUES('AT0000000064', 'aliase', 'string', 'Black san', 'n/a');
INSERT INTO attribute VALUES('AT0000000065', 'devArtGalery', 'url', 'https://www.deviantart.com/blackxsensei/gallery/all', 'n/a');

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
INSERT INTO source_attr VALUES('SRC0000018', 'AT0000000048');
INSERT INTO source_attr VALUES('SRC0000018', 'AT0000000049');
INSERT INTO source_attr VALUES('SRC0000019', 'AT0000000050');
INSERT INTO source_attr VALUES('SRC0000019', 'AT0000000051');
INSERT INTO source_attr VALUES('SRC0000019', 'AT0000000052');
INSERT INTO source_attr VALUES('SRC0000019', 'AT0000000053');
INSERT INTO source_attr VALUES('SRC0000020', 'AT0000000054');
INSERT INTO source_attr VALUES('SRC0000020', 'AT0000000055');
INSERT INTO source_attr VALUES('SRC0000020', 'AT0000000056');
INSERT INTO source_attr VALUES('SRC0000022', 'AT0000000058');
INSERT INTO source_attr VALUES('SRC0000022', 'AT0000000059');
INSERT INTO source_attr VALUES('SRC0000022', 'AT0000000060');
INSERT INTO source_attr VALUES('SRC0000023', 'AT0000000061');
INSERT INTO source_attr VALUES('SRC0000023', 'AT0000000062');
INSERT INTO source_attr VALUES('SRC0000024', 'AT0000000063');
INSERT INTO source_attr VALUES('SRC0000024', 'AT0000000064');
INSERT INTO source_attr VALUES('SRC0000024', 'AT0000000065');

INSERT INTO obj_source VALUES('CH0000001', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000002', 'SRC0000018');
INSERT INTO obj_source VALUES('CH0000003', 'SRC0000019');
INSERT INTO obj_source VALUES('CH0000004', 'SRC0000021');
INSERT INTO obj_source VALUES('CH0000005', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000006', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000007', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000008', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000009', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000010', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000011', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000012', 'SRC0000017');
INSERT INTO obj_source VALUES('CH0000013', 'SRC0000025');
INSERT INTO obj_source VALUES('CH0000014', 'SRC0000025');
INSERT INTO obj_source VALUES('CH0000015', 'SRC0000026');
INSERT INTO obj_source VALUES('CH0000016', 'SRC0000027');
INSERT INTO obj_source VALUES('CH0000017', 'SRC0000028');

INSERT INTO obj_source VALUES('CR0000001', 'SRC0000020');

INSERT INTO obj_source VALUES('FT0000001', 'SRC0000029');

INSERT INTO obj_attr VALUES('CH0000001', 'AT0000000033');
INSERT INTO obj_attr VALUES('CH0000001', 'AT0000000035');
INSERT INTO obj_attr VALUES('CH0000001', 'AT0000000043');
INSERT INTO obj_attr VALUES('CH0000004', 'AT0000000057');
INSERT INTO obj_attr VALUES('CH0000005', 'AT0000000032');
INSERT INTO obj_attr VALUES('CH0000005', 'AT0000000035');
INSERT INTO obj_attr VALUES('CH0000005', 'AT0000000045');
INSERT INTO obj_attr VALUES('CH0000006', 'AT0000000032');
INSERT INTO obj_attr VALUES('CH0000006', 'AT0000000035');
INSERT INTO obj_attr VALUES('CH0000006', 'AT0000000044');
INSERT INTO obj_attr VALUES('CH0000007', 'AT0000000034');
INSERT INTO obj_attr VALUES('CH0000007', 'AT0000000035');
INSERT INTO obj_attr VALUES('CH0000007', 'AT0000000044');
INSERT INTO obj_attr VALUES('CH0000008', 'AT0000000032');
INSERT INTO obj_attr VALUES('CH0000008', 'AT0000000036');
INSERT INTO obj_attr VALUES('CH0000008', 'AT0000000046');
INSERT INTO obj_attr VALUES('CH0000009', 'AT0000000033');
INSERT INTO obj_attr VALUES('CH0000009', 'AT0000000035');
INSERT INTO obj_attr VALUES('CH0000009', 'AT0000000045');
INSERT INTO obj_attr VALUES('CH0000010', 'AT0000000032');
INSERT INTO obj_attr VALUES('CH0000010', 'AT0000000036');
INSERT INTO obj_attr VALUES('CH0000010', 'AT0000000045');
INSERT INTO obj_attr VALUES('CH0000011', 'AT0000000034');
INSERT INTO obj_attr VALUES('CH0000011', 'AT0000000037');
INSERT INTO obj_attr VALUES('CH0000011', 'AT0000000044');
INSERT INTO obj_attr VALUES('CH0000012', 'AT0000000034');
INSERT INTO obj_attr VALUES('CH0000012', 'AT0000000035');
INSERT INTO obj_attr VALUES('CH0000012', 'AT0000000043');
