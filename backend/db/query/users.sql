-- name: InitUsers :exec
insert into users (username, email, info)
values (
    'superadmin',
    'superadmin@superadmin.com',
    '{"role": "super_admin","locale":"en"}'::jsonb
  ),
  (
    'admin',
    'admin@admin.com',
    '{"role": "admin","locale":"en"}'::jsonb
  ) on conflict(email) do update
SET
  username = COALESCE(users.username,EXCLUDED.username),
  info = users.info || excluded.info;

-- name: InitPasswordLogins :exec
insert into password_logins (password, email)
values (
    encode(digest($1, 'sha256'), 'hex'),
    'superadmin@superadmin.com'
  ),
  (
    encode(digest($1, 'sha256'), 'hex'),
    'admin@admin.com'
  ) on conflict(email) do nothing;

-- name: GetUserByUsername :one
select * FROM users where username = $1;

-- name: GetUserByEmail :one
select * FROM users where email = $1;

-- name: VerifyUserCredentials :one
select exists (
  select 1
  from password_logins
  where email = $1 and password = $2
) as valid;


-- name: CreateNewUser :exec
INSERT INTO users (username, email, info)
VALUES ($1, $2, $3)
ON CONFLICT (email) DO UPDATE
SET
  username = COALESCE(users.username,EXCLUDED.username),
  info = users.info || excluded.info;

-- name: CreateNewPasswordLogin :exec
INSERT INTO password_logins (password, email)
VALUES (
  $1,
  $2
);