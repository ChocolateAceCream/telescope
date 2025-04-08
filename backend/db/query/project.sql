-- name: GetProjectByName :one
select * from project
where project_name = $1;

-- name: NewProject :one
insert into project (project_name, comment, creator, status, address)
values(
  $1,
  $2,
  $3,
  $4,
  $5
)
returning *;

-- name: UpdateProject :one
UPDATE project
SET project_name = $1,
  comment = $2,
  status = $3,
  address = $4,
  updated_at = now()
WHERE id = $5
RETURNING *;


-- name: GetProjectPageInfo :one
SELECT (
    SELECT COUNT(*)
    FROM project
  ) AS total_items,
  CEIL(COUNT(*)::FLOAT / $1) AS total_pages
FROM project;

-- name: GetProjectList :many
EXECUTE format(
  '
  SELECT *
  FROM project
  ORDER BY %I %s
  LIMIT %s OFFSET %s',
  sqlc.arg('order_by'),
  sqlc.arg('sort_by'),
  sqlc.arg('limit'),
  sqlc.arg('offset')
);

-- name: GetTotalProjectCount :one
SELECT COUNT(*) FROM project;

-- name: GetProjectByID :one
SELECT * FROM project
WHERE id = $1;