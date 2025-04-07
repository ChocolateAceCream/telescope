-- name: NewSketch :one
insert into sketch (project_name, project_id, uploader_id, full_image_url)
values (
  $1,
  $2,
  $3,
  $4
)
returning *;

-- name: GetSketchesByProjectID :many
select * from sketch
where sketch.project_id = $1;