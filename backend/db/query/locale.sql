-- name: GetAllLocales :many
select language, raw_message,translated_message
from locale;