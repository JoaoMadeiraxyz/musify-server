CREATE TABLE IF NOT EXISTS users (
  "id"        uuid           PRIMARY KEY   NOT NULL    DEFAULT gen_random_uuid(),
  "name"      VARCHAR(255)                 NOT NULL,
  "email"     VARCHAR(255)                 NOT NULL    UNIQUE,
  "password"  VARCHAR(255)                 NOT NULL
)

---- create above / drop below ----

DROP TABLE IF EXISTS users;