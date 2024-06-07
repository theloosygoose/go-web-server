CREATE TABLE IF NOT EXISTS categories 
(id INTEGER PRIMARY KEY,
name TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS category_collection
(category_id INTEGER NOT NULL,
collection_id INTEGER NOT NULL,
FOREIGN KEY (category_id) REFERENCES categories(id),
FOREIGN KEY (collection_id) REFERENCES collections(id));
