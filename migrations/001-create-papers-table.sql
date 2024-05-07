-- Up migration
CREATE TABLE IF NOT EXISTS paper
(
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    title            TEXT,
    authors          TEXT,
    publication_date TEXT,
    publisher        TEXT,
    publication_name TEXT,
    doi              TEXT,
    naid             TEXT,
    url              TEXT NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);