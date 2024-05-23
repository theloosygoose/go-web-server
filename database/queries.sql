-- name: GetPhotoById :one
SELECT img.id, img.name, img.location, img.date, img.description,
img.imagepath, img.i_height, img.i_width, collec.name, collec.id
    FROM photos AS img 
INNER JOIN image_collections AS link ON
    link.photo_id = img.id
INNER JOIN collections AS collec ON
    link.collection_id = collec.id WHERE img.id=?;

-- name: CreatePhoto :one
INSERT INTO photos (name, location, date, description, imagepath, i_height, i_width)
VALUES(?, ?, ?, ?, ?, ?, ?) RETURNING id;

-- name: GetAllPhotos :many
SELECT img.id, img.name, img.date, img.imagepath, img.i_height, img.i_width, collec.name, collec.id
    FROM photos AS img 
INNER JOIN image_collections AS link ON
    link.photo_id = img.id
INNER JOIN collections AS collec ON
    link.collection_id = collec.id;

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
SELECT img.id, img.name, img.date, img.imagepath, img.i_height, img.i_width, collec.name, collec.id
    FROM photos AS img
INNER JOIN image_collections AS link ON
    link.photo_id = img.id
INNER JOIN collections AS collec ON
    link.collection_id = collec.id WHERE collec.id=?;

-- name: GetAllCollections :many
SELECT * FROM collections;

-- name: CreateCollection :one
INSERT INTO collections (name) VALUES (?) RETURNING *;

-- name: DeleteCollection :exec
DELETE FROM collections 
WHERE id=?;
DELETE FROM image_collections 
WHERE collection_id=?;

-- name: PhotoIntoCollection :exec
INSERT INTO image_collections 
    (photo_id, collection_id) 
VALUES (?,?);
