CREATE TABLE IF NOT EXISTS collections 
(id INTEGER PRIMARY KEY,
title TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS imagecollections 
(imageid INTEGER NOT NULL,
tagid INTEGER NOT NULL,
FOREIGN KEY (imageid) REFERENCES collections(id),
FOREIGN KEY (tagid) REFERENCES photos(id),
);
