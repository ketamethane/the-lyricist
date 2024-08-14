CREATE TABLE Lyrics (
    id SERIAL PRIMARY KEY,
    song_id INTEGER REFERENCES Songs(id),
    line_number INTEGER,
    content JSONB,
    translation_id INTEGER REFERENCES Translations(id)
);
