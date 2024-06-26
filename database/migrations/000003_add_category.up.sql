CREATE TABLE IF NOT EXISTS categories 
(id INTEGER PRIMARY KEY,
name TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS category_collection
(collection_id INTEGER NOT NULL,
category_id INTEGER NOT NULL,
FOREIGN KEY (collection_id) REFERENCES collections(id),
FOREIGN KEY (category_id) REFERENCES categories(id));
