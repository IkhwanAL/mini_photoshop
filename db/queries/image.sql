-- name: UploadImage :one
INSERT INTO uploaded_image (
  filename,
  current_version_id
) VALUES ( ?, ?) RETURNING id;

-- name: AddImageToTrack :one
INSERT INTO image_tracker (
  image_id,
  parent_version_id,
  filename,
  operation
) VALUES (?, ?, ?, ?) RETURNING id;
