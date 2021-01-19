CREATE TABLE "seed" (
    "title" TEXT,
    "title_jp" TEXT,
    "show_type" text,
    "source" TEXT, 
    "begin_date" text,
    "end_date" text,
    "genre" VARCHAR[],
    "season" text,
    "year" INT,
    "airing" boolean,
    "current_status" text,
    "num_episodes" integer,
    "episode_duration" text,
    "broadcast_time" text,
    "next_broadcast" text,
    "score" FLOAT,
    "scored_by" INT,
    "rank" INT,
    "popularity" INT,
    "favorites" INT,
    "image_url" TEXT, 
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO seed VALUES
    ('Yakusoku no Neverland', '約束のネバーランド', 'TV', 'Manga', '2019-01-10T00:00:00+00:00', '2019-03-29T00:00:00+00:00', 
    '{"Sci-Fi", "Mystery", "Horror", "Psychological", "Thriller"}',
    'Winter', 2019, FALSE, 'Finished Airing', 12, '22 min per ep', 'Fridays at 00:55 (JST)', NULL,
    8.66, 580064, 54, 65, 27250, 'https://cdn.myanimelist.net/images/anime/1125/96929.jpg'),
    
    ('Tate no Yuusha no Nariagari', '盾の勇者の成り上がり', 'TV', 'Light novel', '2019-01-09T00:00:00+00:00', '2019-06-26T00:00:00+00:00', 
    '{"Action", "Adventure", "Drama", "Fantasy"}',
    'Winter', 2019, FALSE, 'Finished Airing', 25, '24 min per ep', 'Wednesdays at 22:00 (JST)', NULL,
    8.01, 474202, 527, 89, 17086, 'https://cdn.myanimelist.net/images/anime/1490/101365.jpg'),
    
    ('Kaguya-sama wa Kokurasetai: Tensai-tachi no Renai Zunousen', 'かぐや様は告らせたい～天才たちの恋愛頭脳戦～', 'TV', 'Manga', '2019-01-12T00:00:00+00:00', '2019-03-30T00:00:00+00:00', 
    '{"Comedy", "Psychological", "Romance", "School", "Seinen"}',
    'Winter', 2019, FALSE, 'Finished Airing', 12, '25 min per ep', 'Saturdays at 23:30 (JST)', NULL,
    8.44, 486187, 129, 93, 19779, 'https://cdn.myanimelist.net/images/anime/1295/106551.jpg'),
    
    ('Mob Psycho 100 II', 'モブサイコ100 II', 'TV', 'Web manga', '2019-01-07T00:00:00+00:00', '2019-04-01T00:00:00+00:00', 
    '{"Action", "Slice of Life", "Comedy", "Supernatural"}',
    'Winter', 2019, FALSE, 'Finished Airing', 13, '24 min per ep', 'Mondays at 23:00 (JST)', NULL,
    8.85, 446851, 22, 104, 18908, 'https://cdn.myanimelist.net/images/anime/1918/96303.jpg'),
    
    ('Dororo', 'どろろ', 'TV', 'Manga', '2019-01-07T00:00:00+00:00', '2019-06-24T00:00:00+00:00', 
    '{"Action", "Adventure", "Demons", "Historical", "Samurai", "Shounen", "Supernatural"}',
    'Winter', 2019, FALSE, 'Finished Airing', 24, '24 min per ep', 'Mondays at 22:30 (JST)', NULL,
    8.19, 292867, 312, 143, 9815, 'https://cdn.myanimelist.net/images/anime/1879/100467.jpg'),
    
    ('Kakegurui××', '賭ケグルイ××', 'TV', 'Manga', '2019-01-09T00:00:00+00:00', '2019-03-27T00:00:00+00:00', 
    '{"Drama", "Game", "Mystery", "Psychological", "School", "Shounen"}',
    'Winter', 2019, FALSE, 'Finished Airing', 12, '24 min per ep', 'Wednesdays at 02:30 (JST)', NULL,
    7.36, 202008, 2073, 329, 1383, 'https://cdn.myanimelist.net/images/anime/1496/96519.jpg'),
    
    ('5-toubun no Hanayome', '五等分の花嫁', 'TV', 'Manga', '2019-01-11T00:00:00+00:00', '2019-03-29T00:00:00+00:00', 
    '{"Harem", "Comedy", "Romance", "School", "Shounen"}',
    'Winter', 2019, FALSE, 'Finished Airing', 12, '24 min per ep', 'Fridays at 01:28 (JST)', NULL,
    7.58, 215968, 1328, 337, 5249, 'https://cdn.myanimelist.net/images/anime/1819/97947.jpg'),
    
    ('Domestic na Kanojo', 'ドメスティックな彼女', 'TV', 'Manga', '2019-01-12T00:00:00+00:00', '2019-03-30T00:00:00+00:00', 
    '{"Drama", "Romance", "School", "Shounen"}',
    'Winter', 2019, FALSE, 'Finished Airing', 12, '25 min per ep', 'Saturdays at 01:55 (JST)', NULL,
    6.8, 197996, 4409, 348, 3027, 'https://cdn.myanimelist.net/images/anime/1021/95670.jpg'),
    
    ('Date A Live III', 'デート・ア・ライブⅢ', 'TV', 'Light novel', '2019-01-11T00:00:00+00:00', '2019-03-29T00:00:00+00:00', 
    '{"Sci-Fi", "Harem", "Comedy", "Romance", "Mecha", "School"}',
    'Winter', 2019, FALSE, 'Finished Airing', 12, '24 min per ep', 'Fridays at 21:30 (JST)', NULL,
    7.13, 89768, 3048, 677, 2495, 'https://cdn.myanimelist.net/images/anime/1055/100468.jpg'),
    
    ('Boogiepop wa Warawanai (2019)', 'ブギーポップは笑わない', 'TV', 'Light novel', '2019-01-04T00:00:00+00:00', '2019-03-29T00:00:00+00:00', 
    '{"Psychological", "Mystery", "Horror"}',
    'Winter', 2019, FALSE, 'Finished Airing', 18, '23 min per ep', 'Fridays at 21:00 (JST)', NULL,
    7.11, 42642, 3130, 885, 570, 'https://cdn.myanimelist.net/images/anime/1135/95454.jpg'),

    -- Fall 2020
    ('Jujutsu Kaisen (TV)', '呪術廻戦', 'TV', 'Manga', '2020-10-03T00:00:00+00:00', NULL, 
    '{"Action", "Demons", "Supernatural", "School", "Shounen"}',
    'Fall', 2020, TRUE, 'Currently Airing', 24, '23 min per ep', 'Saturdays at 01:25 (JST)', 'Tue, 19 Jan 2021 05:39:18 +0900',
    8.48, 118870, 106, 259, 8877, 'https://cdn.myanimelist.net/images/anime/1171/109222.jpg'),
    
    ('Haikyuu!!: To the Top 2nd Season', 'ハイキュー!! TO THE TOP', 'TV', 'Manga', '2020-10-03T00:00:00+00:00', '2020-12-19T00:00:00+00:00', 
    '{"Comedy", "Sports", "Drama", "School", "Shounen"}',
    'Fall', 2020, TRUE, 'Finished Airing', 12, '23 min per ep', 'Saturdays at 02:25 (JST)', 'Tue, 19 Jan 2021 05:39:16 +0900', 
    8.6, 96367, 69, 486, 3359, 'https://cdn.myanimelist.net/images/anime/1453/106768.jpg'),
    
    ('Dungeon ni Deai wo Motomeru no wa Machigatteiru Darou ka III', 'ダンジョンに出会いを求めるのは間違っているだろうかIII', 'TV', 'Light novel', '2020-10-03T00:00:00+00:00', '2020-12-19T00:00:00+00:00', 
    '{"Action", "Adventure", "Comedy", "Romance", "Fantasy"}',
    'Fall', 2020, FALSE, 'Finished Airing', 12, '24 min per ep', 'Saturdays at 00:30 (JST)', 'Tue, 19 Jan 2021 00:28:22 +0900', 
    7.52, 67093, 1516, 541, 1496, 'https://cdn.myanimelist.net/images/anime/1523/108380.jpg'),
    
    ('Tonikaku Kawaii', 'トニカクカワイイ', 'TV', 'Manga', '2020-10-03T00:00:00+00:00', '2020-12-19T00:00:00+00:00', 
    '{"Comedy", "Romance", "Shounen"}',
    'Fall', 2020, FALSE, 'Finished Airing', 12, '23 min per ep', 'Saturdays at 01:05 (JST)', 'Mon, 18 Jan 2021 23:58:04 +0900', 
    8.01, 84430, 528, 583, 5339, 'https://cdn.myanimelist.net/images/anime/1613/108722.jpg'),
    
    ('Mahouka Koukou no Rettousei: Raihousha-hen', '魔法科高校の劣等生 来訪者編', 'TV', 'Light novel', '2020-10-04T00:00:00+00:00', '2020-12-27T00:00:00+00:00', 
    '{"Action", "Sci-Fi", "Supernatural", "Magic", "Romance", "School"}',
    'Fall', 2020, FALSE, 'Finished Airing', 13, '23 min per ep', 'Sundays at 00:30 (JST)', 'Mon, 18 Jan 2021 18:42:39 +0900', 
    7.38, 34920, 1992, 833, 1112, 'https://cdn.myanimelist.net/images/anime/1788/106668.jpg'),
    
    ('Majo no Tabitabi', '魔女の旅々', 'TV', 'Light novel', '2020-10-02T00:00:00+00:00', '2020-12-18T00:00:00+00:00', 
    '{"Adventure", "Magic", "Fantasy"}',
    'Fall', 2020, FALSE, 'Finished Airing', 12, '24 min per ep', 'Sundays at 00:30 (JST)', 'Mon, 18 Jan 2021 18:36:45 +0900', 
    7.6, 47622, 1279, 879, 1365, 'https://cdn.myanimelist.net/images/anime/1802/108501.jpg'),
    
    ('Akudama Drive', 'アクダマドライブ', 'TV', 'Original', '2020-10-08T00:00:00+00:00', '2020-12-24T00:00:00+00:00', 
    '{"Action", "Sci-Fi"}',
    'Fall', 2020, FALSE, 'Finished Airing', 12, '23 min per ep', 'Thursdays at 21:30 (JST)', 'Mon, 18 Jan 2021 15:31:08 +0900', 
    7.82, 44324, 779, 915, 1249, 'https://cdn.myanimelist.net/images/anime/1468/109172.jpg'),
    
    ('Higurashi no Naku Koro ni Gou', 'ひぐらしのなく頃に業', 'TV', 'Visual novel', '2020-10-01T00:00:00+00:00', NULL, 
    '{"Dementia", "Horror", "Mystery", "Psychological", "Supernatural", "Thriller"}',
    'Fall', 2020, TRUE, 'Currently Airing', 24, '23 min per ep', 'Thursdays at 23:30 (JST)', 'Tue, 19 Jan 2021 05:39:16 +090', 
    7.14, 18757, 3029, 1016, 462, 'https://cdn.myanimelist.net/images/anime/1287/109031.jpg'),
    
    ('Kamisama ni Natta Hi', '神様になった日', 'TV', 'Original', '2020-10-11T00:00:00+00:00', '2020-12-27T00:00:00+00:00', 
    '{"Drama", "Fantasy"}',
    'Fall', 2020, FALSE, 'Finished Airing', 12, '24 min per ep', 'Sundays at 00:00 (JST)', 'Mon, 18 Jan 2021 15:11:58 +0900', 
    6.99, 35255, 3652, 1028, 1606, 'https://cdn.myanimelist.net/images/anime/1396/109465.jpg'),
    
    ('Munou na Nana', '無能なナナ', 'TV', 'Manga', '2020-10-04T00:00:00+00:00', '2020-12-27T00:00:00+00:00', 
    '{"Psychological", "Shounen", "Super Power", "Supernatural", "Thriller"}',
    'Fall', 2020, FALSE, 'Finished Airing', 13, '23 min per ep', 'Sundays at 21:30 (JST)', 'Mon, 18 Jan 2021 02:18:15 +0900', 
    7.41, 38836, 1865, 1099, 690, 'https://cdn.myanimelist.net/images/anime/1301/110433.jpg');