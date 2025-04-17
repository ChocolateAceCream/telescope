-- name: GetEmailWhitelist :many
select email
from email_whitelist;


-- name: InitEmailWhitelist :exec
insert into email_whitelist (email)
values ('vincewang0101@gmail.com'),
  ('nuodi@hotmail.com'),
  ('ajvalino09@gmail.com'),
  ('superadmin@superadmin.com'),
  ('admin@admin.com'),
  ('liuguoxinn87@gmail.com')
ON CONFLICT (email) DO NOTHING;;

-- name: VerifyEmailWhitelistExist :one
SELECT EXISTS (
    SELECT 1
    FROM email_whitelist
    WHERE email = $1
  ) AS exists;

-- name: AddEmailWhitelist :exec
INSERT INTO email_whitelist (email)
VALUES ($1)
ON CONFLICT (email) DO NOTHING;

-- name: DeleteEmailWhitelist :exec
DELETE FROM email_whitelist
WHERE email = $1;