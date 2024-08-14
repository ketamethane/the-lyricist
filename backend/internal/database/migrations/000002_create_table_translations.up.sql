CREATE TABLE Translations (
    id SERIAL PRIMARY KEY,
    song_id INTEGER REFERENCES Songs(id),
    language CHAR(2),
    content JSONB,
    composers JSONB,
    lyrics JSONB,
    producers JSONB
);

-- KIV "rj" as romaji