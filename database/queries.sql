-- name: GetPhotoById :one
SELECT id, name, location, date, description,
imagepath, i_height, i_width
FROM photos 
WHERE id = ?;

-- name: CreatePhoto :one
INSERT INTO photos (name, location, date, description, imagepath, i_height, i_width)
VALUES(?, ?, ?, ?, ?, ?, ?) RETURNING id;

-- name: GetAllPhotos :many
SELECT id, name, date, imagepath, i_height, i_width FROM photos;

-- name: DeletePhoto :one
DELETE FROM photos
WHERE id = ?
RETURNING imagepath;

-- name: GetRandomPhoto :one
SELECT id, name, location, date, description, imagepath, i_height, i_width
FROM photos
ORDER BY random()
LIMIT 1;

-- name: UpdatePhotoDescription :exec
UPDATE photos
SET name=?, location=?,
date=?, description=?, imagepath=?, i_height=?, i_width=?
WHERE id=?;

-- name: GetCollectionPhotos :many
SELECT img.id, img.name, img.location, img.date, img.imagepath
    FROM photos AS img
INNER JOIN image_collections AS link ON
    link.photo_id = img.id
INNER JOIN collections AS collec ON
    link.collection_id = collec.id;

-- name: GetAllCollections :many
SELECT * FROM collections;

-- name: CreateCollection :exec
INSERT INTO collections (name) VALUES (?);

-- name: DeleteCollection :exec
DELETE FROM collections 
WHERE id=?;
DELETE FROM image_collections 
WHERE collection_id=?;

-- name: UpdateCollection :exec
INSERT INTO image_collections 
    (photo_id, collection_id) 
VALUES (?,?);
