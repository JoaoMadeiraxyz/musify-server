CREATE TABLE IF NOT EXISTS artists (
  "id"        UUID           PRIMARY KEY   NOT NULL    DEFAULT gen_random_uuid(),
  "user_id"   UUID                         NOT NULL,
  "followers" BIGINT                       NOT NULL    DEFAULT 0,
  
  FOREIGN KEY (user_id) REFERENCES users(id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS artists;