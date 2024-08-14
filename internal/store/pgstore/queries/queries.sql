-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT
  *
FROM users
WHERE
  id = $1;

-- name: InsertUser :one
INSERT INTO users
  ( "id", "name", "email", "password" ) VALUES
  ( $1, $2, $3, $4 )
RETURNING "id";

-- name: GetArtists :many
SELECT * FROM artists;

-- name: GetArtist :one
SELECT
  *
FROM artists
WHERE
  id = $1;

-- name: GetUserArtistData :one
SELECT
  *
FROM artists
WHERE
  user_id = $1;

-- name: InsertArtist :one
INSERT INTO artists
  ( "id", "user_id" ) VALUES
  ( $1, $2 )
RETURNING "id";
