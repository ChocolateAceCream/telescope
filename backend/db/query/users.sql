-- name: FindUserByUsername :exec
SELECT *
FROM a_users
WHERE username = $1;
-- name: VerifyUserCredentials :one
select exists (
    select 1
    from a_users
    where username = $1
      and password = $2
  ) as valid;

-- name: InitUsers :exec
insert into a_users (username, password, email, role)
values (
    'superadmin',
    encode(digest($1, 'sha256'), 'hex'),
    'superadmin@superadmin.com',
    'super_admin'
  ),
  (
    'admin',
    encode(digest($1, 'sha256'), 'hex'),
    'admin@admin.com',
    'admin'
  ) on conflict(username) do
update
SET password = excluded.password,
  email = excluded.email,
  role = excluded.role;

-- name: GetUserByUsername :one
select * FROM a_users where username = $1;