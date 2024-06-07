CREATE TABLE IF NOT EXISTS collections 
(id INTEGER PRIMARY KEY,
name TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS image_collections 
(photo_id INTEGER NOT NULL,
collection_id INTEGER NOT NULL,
FOREIGN KEY (photo_id) REFERENCES photos(id),
FOREIGN KEY (collection_id) REFERENCES collections(id));
