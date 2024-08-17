CREATE TABLE Translations (
    id SERIAL PRIMARY KEY,
    song_id INTEGER REFERENCES Songs(id),
    language CHAR(2),
    content JSONB,
    composers JSONB,
    lyrics_full JSONB,
    producers JSONB,
    video_link TEXT
);

-- KIV "rj" as romaji