-- name: GoogleLoginUpdateUser :exec
INSERT INTO users (username, email, info)
VALUES (
  $1, -- username
  $2, -- email
  $3  -- info
)
ON CONFLICT (email) DO UPDATE
SET
  username = COALESCE(users.username,EXCLUDED.username),
  info = users.info || EXCLUDED.info;

-- name: GoogleLogin :exec
INSERT INTO google_logins (email, username, access_token, issued_at, expired_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (email) DO UPDATE
SET
  username = EXCLUDED.username,
  access_token = EXCLUDED.access_token,
  issued_at = EXCLUDED.issued_at,
  expired_at = EXCLUDED.expired_at;
