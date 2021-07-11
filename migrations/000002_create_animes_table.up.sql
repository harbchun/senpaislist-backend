CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS animes(
   id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
   title VARCHAR NOT NULL,
   title_jp VARCHAR,
   source VARCHAR,
   summary VARCHAR,
   image_id VARCHAR,
   syoboi_tid INT,
   CONSTRAINT animes_id_syoboi_tid UNIQUE (id, syoboi_tid)
);
