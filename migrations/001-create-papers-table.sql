CREATE TABLE IF NOT EXISTS papers
(
    id               INT AUTO_INCREMENT PRIMARY KEY,
    title            TEXT,
    authors          TEXT,
    publication_date TEXT,
    publisher        TEXT,
    publication_name TEXT,
    doi              TEXT,
    naid             TEXT,
    url              TEXT NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);