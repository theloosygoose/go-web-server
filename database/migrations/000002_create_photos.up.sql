CREATE TABLE IF NOT EXISTS photos
(id INTEGER PRIMARY KEY, 
    name VARCHAR(100) NOT NULL, 
    location VARCHAR(100) NOT NULL, 
    date VARCHAR(100), 
    imagepath VARCHAR(100) NOT NULL, 
    description TEXT, 
    i_height VARCHAR(10), 
    i_width VARCHAR(10));
