CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
   id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
   last_name VARCHAR NOT NULL,
   first_name VARCHAR NOT NULL,
   username VARCHAR NOT NULL UNIQUE,
   password VARCHAR NOT NULL,
   email VARCHAR NOT NULL
);
